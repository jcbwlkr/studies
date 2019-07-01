package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("http://api.meetup.com/2/events?status=upcoming&order=time")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Add("limited_events", "False")
	q.Add("group_urlname", "devict")
	q.Add("desc", "false")
	q.Add("offset", "0")
	q.Add("photo-host", "public")
	q.Add("format", "json")
	q.Add("page", "20")
	q.Add("fields", "")
	q.Add("sig_id", "73273692")
	q.Add("sig", "9cdd3af6b5a26eb664fe5abab6e5cf7bfaaf090e")
	u.RawQuery = q.Encode()

	//fmt.Println(u.String())

	q = url.Values{}
	q.Add("filter", `{"bar": "baz"}`)

	fmt.Println(q.Encode())

	u2 := url.URL{
		Scheme:   "http",
		Host:     "localhost:8070",
		Path:     "api/history",
		RawQuery: q.Encode(),
	}

	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println(u2.String())
}
