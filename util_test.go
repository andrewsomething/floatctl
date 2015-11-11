package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhoAmI(t *testing.T) {
	resp := "12345"

	cmd := &cobra.Command{Use: "Testing..."}
	metadataServer(t, "/metadata/v1/id", resp, func() {
		id := WhoAmI(cmd)
		droplet, err := strconv.Atoi(resp)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, droplet, id, "they should be equal")
	})
}

func TestAssignedFIP(t *testing.T) {
	resp := `{"floating_ip": {"ipv4": {"ip_address": "192.168.0.100","active": true}}}`

	cmd := &cobra.Command{Use: "Testing..."}
	metadataServer(t, "/metadata/v1.json", resp, func() {
		assigned := AssignedFIP(cmd)

		assert.Equal(t, assigned, "192.168.0.100", "they should be equal")
	})
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
