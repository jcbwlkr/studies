package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Logger is a logging function that can be injected where needed. In
// production it should write to stderr but in tests it can be a t.Logf.
type Logger func(string, ...interface{})

// App holds the shared state and various handlers of our application.
type App struct {
	logf Logger
}

// NewApp constructs an App from its dependencies and wires all of its handlers
// to different http routes.
func NewApp(logf Logger) http.Handler {

	a := App{logf: logf}

	m := http.NewServeMux()

	m.HandleFunc("/", a.Home)

	return m
}

// Home serves the home page of our site
func (a App) Home(w http.ResponseWriter, r *http.Request) {
	a.logf("Home was called with query values %s", r.URL.RawQuery)

	if r.URL.Query().Get("foo") != "bar" {
		a.logf("Blocking user")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	a.logf("Letting user in")

	fmt.Fprint(w, "Welcome to my site!")
}

func main() {

	// Make a logger that writes to stderr. We wrap the Output method with a
	// calldepth value of 2 so the line number reported is where we actually
	// invoke this func.
	l := log.New(os.Stderr, "[app] ", log.LstdFlags|log.Lshortfile)
	logger := func(format string, args ...interface{}) {
		l.Output(2, fmt.Sprintf(format, args...))
	}

	a := NewApp(logger)

	log.Fatal(http.ListenAndServe(":3000", a))
}
