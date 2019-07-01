package main

import (
	"encoding/json"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPatch(t *testing.T) {

	body := strings.NewReader(`{
		"firstName": "Anna",
		"lastName": "Walker",
		"company": "Anna's Awesome Artichokes",
		"title": null
	}`)
	r := httptest.NewRequest("PATCH", "/patch", body)
	w := httptest.NewRecorder()

	NewApp("patch").ServeHTTP(w, r)

	if w.Code != 200 {
		t.Fatalf("ded")
	}

	want := Contact{
		ID:        seed.ID,
		FirstName: "Anna",
		LastName:  "Walker",
		Company:   "Anna's Awesome Artichokes",
		Title:     "",
		Roles:     []string{"Marketing", "Food Safety", "Executive"},
	}

	var got Contact
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("response fail\n%s", diff)
	}
}

func TestManual(t *testing.T) {

	body := strings.NewReader(`{
		"firstName": "Anna",
		"lastName": "Walker",
		"company": "Anna's Awesome Artichokes",
		"title": null
	}`)
	r := httptest.NewRequest("PATCH", "/manual", body)
	w := httptest.NewRecorder()

	NewApp("manual").ServeHTTP(w, r)

	if w.Code != 200 {
		t.Fatalf("ded")
	}

	want := Contact{
		ID:        seed.ID,
		FirstName: "Anna",
		LastName:  "Walker",
		Company:   "Anna's Awesome Artichokes",
		Title:     "",
		Roles:     []string{"Marketing", "Food Safety", "Executive"},
	}

	var got Contact
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("response fail\n%s", diff)
	}
}

func BenchmarkAPIManual(b *testing.B) {

	a := NewApp("manual")

	for i := 0; i < b.N; i++ {
		is := strconv.Itoa(i)
		body := strings.NewReader(`{
			"firstName": "Anna-` + is + `",
			"lastName": "Walker-` + is + `",
			"company": "Anna's Awesome Artichokes-` + is + `"
		}`)
		r := httptest.NewRequest("PATCH", "/manual", body)
		w := httptest.NewRecorder()

		a.ServeHTTP(w, r)

		if w.Code != 200 {
			b.Fatalf("ded")
		}
	}
}

func BenchmarkAPIPatch(b *testing.B) {

	a := NewApp("patch")

	for i := 0; i < b.N; i++ {
		id := strconv.Itoa(i)
		body := strings.NewReader(`{
			"firstName": "Anna-` + id + `",
			"lastName": "Walker-` + id + `",
			"company": "Anna's Awesome Artichokes-` + id + `"
		}`)
		r := httptest.NewRequest("PATCH", "/patch", body)
		w := httptest.NewRecorder()

		a.ServeHTTP(w, r)

		if w.Code != 200 {
			b.Fatalf("ded")
		}
	}
}
