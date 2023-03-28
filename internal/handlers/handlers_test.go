package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{
		name:               "home",
		url:                "/",
		method:             "GET",
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "about",
		url:                "/about",
		method:             "GET",
		expectedStatusCode: http.StatusOK,
	}, {
		name:               "major",
		url:                "/majors-suite",
		method:             "GET",
		expectedStatusCode: http.StatusOK,
	}, {
		name:               "general",
		url:                "/generals-quarters",
		method:             "GET",
		expectedStatusCode: http.StatusOK,
	}, {
		name:               "contact",
		url:                "/contact",
		method:             "GET",
		expectedStatusCode: http.StatusOK,
	}, {
		name:               "make-reservation",
		url:                "/make-reservation",
		method:             "GET",
		expectedStatusCode: http.StatusOK,
	}, {
		name:   "search-availability",
		url:    "/search-availability",
		method: "POST",
		params: []postData{
			{key: "start", value: "2023-03-21"},
			{key: "end", value: "2023-03-23"},
		},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:   "search-availability-json",
		url:    "/search-availability-json",
		method: "POST",
		params: []postData{
			{key: "start", value: "2023-03-21"},
			{key: "end", value: "2023-03-23"},
		},
		expectedStatusCode: http.StatusOK,
	}, {
		name:   "make-reservation",
		url:    "/make-reservation",
		method: "POST",
		params: []postData{
			{key: "first_name", value: "John"},
			{key: "last_name", value: "Smith"},
			{key: "email", value: "me@golang.com"},
			{key: "phone", value: "555-555-55555"},
		},
		expectedStatusCode: http.StatusOK,
	},
}

func TestNewHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}

		if e.method == "POST" {
			values := url.Values{}
			for _, x := range e.params {
				values.Set(x.key, x.value)
			}

			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}

	}
}
