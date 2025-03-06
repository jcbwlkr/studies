package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
)

// User is a registered user of our system.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// FetchUser makes an API call to a given URL and marshals the result into a
// User. If the server responds with a non-200 status then an error is
// returned.
func FetchUser(url, token string) (User, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Authorization", "token "+token)

	c := http.Client{Timeout: 5 * time.Second}

	res, err := c.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return User{}, errors.New("api responded with " + res.Status)
	}

	var u User
	if err := json.NewDecoder(res.Body).Decode(&u); err != nil {
		return User{}, err
	}

	return u, nil
}

// mockUserServer mimics the functionality of an upstream API server. Requests
// should be to the path /api/users/:id and server will respond with the user's
// data in JSON form. Two users are defined with ids 1 and 2. Requests for
// other users or to other paths will receive a 404. Requests should have an
// access token set in the Authorization header or receive a 401. The header
// should have the prefix "token "
//
// The caller should handle closing this server
func mockUserServer() *httptest.Server {
	tokenCheck := regexp.MustCompile("^token .*$")
	pathCheck := regexp.MustCompile("^/api/users/[0-9]+$")

	fn := func(w http.ResponseWriter, r *http.Request) {
		if !tokenCheck.MatchString(r.Header.Get("Authorization")) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !pathCheck.MatchString(r.URL.Path) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		users := map[string]string{
			"1": `{"id": 1, "name": "Jenn Walker", "email": "jenn@example.com"}`,
			"2": `{"id": 2, "name": "Jacob Walker", "email": "jacob@example.com"}`,
		}

		id := r.URL.Path[len("/api/users/"):]
		u, ok := users[id]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.Write([]byte(u))
	}

	return httptest.NewServer(http.HandlerFunc(fn))
}

func TestFetchUserTable(t *testing.T) {
	srv := mockUserServer()
	defer srv.Close()

	tests := []struct {
		Name      string
		URL       string
		Token     string
		Want      User
		ShouldErr bool
	}{
		{"User 1", srv.URL + "/api/users/1", "secret", User{ID: 1, Name: "Jenn Walker", Email: "jenn@example.com"}, false},
		{"User 2", srv.URL + "/api/users/2", "secret", User{ID: 2, Name: "Jacob Walker", Email: "jacob@example.com"}, false},
		{"User 3", srv.URL + "/api/users/3", "secret", User{}, true},
		{"Bad path", srv.URL + "/api/projects", "secret", User{}, true},
		{"Bad url", "", "secret", User{}, true},
		{"No Token", srv.URL + "/api/users/1", "", User{}, true},
	}

	for i, test := range tests {
		t.Logf("Test %d: %s", i, test.Name)

		got, err := FetchUser(test.URL, test.Token)
		if test.ShouldErr && err == nil {
			t.Errorf("%d: FetchUser should error but did not", i)
		}
		if !test.ShouldErr && err != nil {
			t.Errorf("%d: FetchUser should not error but got %v", i, err)
		}
		if got != test.Want {
			t.Errorf("%d: fetched user did not match", i)
			t.Errorf("%d: Got : %+v", i, got)
			t.Errorf("%d: Want: %+v", i, test.Want)
		}
	}
}

func TestFetchUserSubs(t *testing.T) {

	// test accepts the parameters needed to construct a single test case for
	// FetchUser. It returns a func that can be fed into t.Run
	test := func(url, token string, want User, shouldErr bool) func(*testing.T) {
		return func(t *testing.T) {
			got, err := FetchUser(url, token)
			if shouldErr && err == nil {
				t.Errorf("FetchUser should error but did not")
			}
			if !shouldErr && err != nil {
				t.Errorf("FetchUser should not error but got %v", err)
			}
			if got != want {
				t.Errorf("fetched user did not match")
				t.Errorf("Got : %+v", got)
				t.Errorf("Want: %+v", want)
			}
		}
	}

	srv := mockUserServer()
	defer srv.Close()

	t.Run("User 1", test(srv.URL+"/api/users/1", "secret", User{ID: 1, Name: "Jenn Walker", Email: "jenn@example.com"}, false))
	t.Run("User 2", test(srv.URL+"/api/users/2", "secret", User{ID: 2, Name: "Jacob Walker", Email: "jacob@example.com"}, false))
	t.Run("User 3", test(srv.URL+"/api/users/3", "secret", User{}, true))
	t.Run("Bad path", test(srv.URL+"/api/projects", "secret", User{}, true))
	t.Run("Bad url", test("", "secret", User{}, true))
	t.Run("No Token", test(srv.URL+"/api/users/1", "", User{}, true))
}

func TestFetchUserTableSubs(t *testing.T) {
	srv := mockUserServer()
	defer srv.Close()

	tests := []struct {
		Name      string
		URL       string
		Token     string
		Want      User
		ShouldErr bool
	}{
		{
			Name:      "User 1",
			URL:       srv.URL + "/api/users/1",
			Token:     "secret",
			Want:      User{ID: 1, Name: "Jenn Walker", Email: "jenn@example.com"},
			ShouldErr: false,
		},
		{
			Name:      "User 2",
			URL:       srv.URL + "/api/users/2",
			Token:     "secret",
			Want:      User{ID: 2, Name: "Jacob Walker", Email: "jacob@example.com"},
			ShouldErr: false,
		},
		{
			Name:      "User 3",
			URL:       srv.URL + "/api/users/3",
			Token:     "secret",
			Want:      User{},
			ShouldErr: true,
		},
		{
			Name:      "Bad path",
			URL:       srv.URL + "/api/projects",
			Token:     "secret",
			Want:      User{},
			ShouldErr: true,
		},
		{
			Name:      "Bad url",
			URL:       "",
			Token:     "secret",
			Want:      User{},
			ShouldErr: true,
		},
		{
			Name:      "No Token",
			URL:       srv.URL + "/api/users/1",
			Token:     "",
			Want:      User{},
			ShouldErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got, err := FetchUser(test.URL, test.Token)
			if test.ShouldErr && err == nil {
				t.Errorf("FetchUser should error but did not")
			}
			if !test.ShouldErr && err != nil {
				t.Errorf("FetchUser should not error but got %v", err)
			}
			if got != test.Want {
				t.Errorf("fetched user did not match")
				t.Errorf("Got : %+v", got)
				t.Errorf("Want: %+v", test.Want)
			}
		})
	}
}
