package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func apiServer(t testing.TB, path string, resp string, test func()) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			t.Errorf("Wrong URL: %v", r.URL.String())
			return
		}
		w.WriteHeader(200)
		fmt.Fprintln(w, resp)
	}))

	u, err := url.Parse(server.URL)
	if err != nil {
		panic(err)
	}
	GodoBase = u

	defer server.Close()
	test()
}

func metadataServer(t testing.TB, path string, resp string, test func()) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			t.Errorf("Wrong URL: %v", r.URL.String())
			return
		}
		w.WriteHeader(200)
		fmt.Fprintln(w, resp)
	}))

	u, err := url.Parse(server.URL)
	if err != nil {
		panic(err)
	}
	MetadataBase = u

	defer server.Close()
	test()
}
