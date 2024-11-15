package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"syscall"

	"go.opencensus.io/examples/exporter"
	"go.opencensus.io/plugin/ochttp/propagation/tracecontext"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
)

var hf tracecontext.HTTPFormat

var reqs bytes.Buffer

func main() {

	e := &exporter.PrintExporter{}
	view.RegisterExporter(e)
	trace.RegisterExporter(e)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	http.HandleFunc("/", handler)

	go func() {
		srv := http.Server{Addr: ":8080"}
		log.Println("listening on", srv.Addr)
		srv.ListenAndServe()
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	fmt.Println("Waiting for shutdown")
	<-ch

	fmt.Println("Shutting down")

	fmt.Println("Upstream requests made")
	fmt.Println(reqs.String())
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx, span := trace.StartSpan(r.Context(), r.URL.Path)
	defer span.End()

	// Make 2 requests to upstream servers.
	_, _ = fetch(ctx, "gophers")
	_, _ = fetch(ctx, "concurrency")

	fmt.Fprintf(w, "Hello, world\n")
}

func fetch(ctx context.Context, page string) ([]byte, error) {
	ctx, span := trace.StartSpan(ctx, "fetch/"+page)
	defer span.End()

	// Make a request to send to an upstream server. Annotate it with the span context.
	r, _ := http.NewRequest("GET", "google.com/"+page, nil)
	hf.SpanContextToRequest(span.SpanContext(), r)

	// Pretend to make the request. Just dump it to a string and stick it in the buffer.
	req, _ := httputil.DumpRequest(r, false)
	reqs.WriteString("Request for page: \"" + page + "\"\n")
	reqs.Write(req)
	reqs.WriteString("\n")

	return nil, nil
}
