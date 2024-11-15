package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func logger(t *testing.T) func(format string, args ...interface{}) {
	t.Helper()
	return t.Logf
}

func TestHomePasses(t *testing.T) {
	a := NewApp(logger(t))

	req := httptest.NewRequest(http.MethodGet, "/?foo=bar", nil)
	res := httptest.NewRecorder()

	a.ServeHTTP(res, req)

	got := http.StatusOK
	want := res.Code
	if got != want {
		t.Error("Didn't get expected status code")
		t.Log("Got ", got)
		t.Log("Want", want)
	}
}

func TestHomeFails(t *testing.T) {
	a := NewApp(logger(t))

	req := httptest.NewRequest(http.MethodGet, "/?foo=foo", nil)
	res := httptest.NewRecorder()

	a.ServeHTTP(res, req)

	got := http.StatusOK
	want := res.Code
	if got != want {
		t.Error("Didn't get expected response")
		t.Log("Got ", got)
		t.Log("Want", want)
	}
}
