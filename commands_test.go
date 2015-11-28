package main

import (
	"github.com/digitalocean/godo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssign_doAssign(t *testing.T) {
	resp := `{"action":{"status":"in-progress", "region":{"slug":"nyc3"}}}`
	expected := &godo.Action{Status: "in-progress", Region: &godo.Region{Slug: "nyc3"}}

	apiServer(t, "/v2/floating_ips/192.168.0.1/actions", resp, func() {
		action := doAssign("192.168.0.1", 1)
		assert.Equal(t, action, expected, "they should be equal")
	})
}

func TestCreate_doCreate_droplet(t *testing.T) {
	resp := `{"floating_ip":{"region":{"slug":"nyc3"},"droplet":{"id":1},"ip":"192.168.0.1"}}`
	expected := &godo.FloatingIP{Region: &godo.Region{Slug: "nyc3"}, Droplet: &godo.Droplet{ID: 1}, IP: "192.168.0.1"}

	createRequest := &godo.FloatingIPCreateRequest{
		DropletID: 1,
	}

	apiServer(t, "/v2/floating_ips", resp, func() {
		floatingIP := doCreate(createRequest)
		assert.Equal(t, floatingIP, expected, "they should be equal")
	})
}

func TestCreate_doCreate_region(t *testing.T) {
	resp := `{"floating_ip":{"region":{"slug":"nyc3"},"droplet":{"id":1},"ip":"192.168.0.1"}}`
	expected := &godo.FloatingIP{Region: &godo.Region{Slug: "nyc3"}, Droplet: &godo.Droplet{ID: 1}, IP: "192.168.0.1"}

	createRequest := &godo.FloatingIPCreateRequest{
		Region: "nyc3",
	}

	apiServer(t, "/v2/floating_ips", resp, func() {
		floatingIP := doCreate(createRequest)
		assert.Equal(t, floatingIP, expected, "they should be equal")
	})
}

func TestList_doList(t *testing.T) {
	resp := `{"floating_ips": [{"region":{"slug":"nyc3"},"droplet":{"id":1},"ip":"192.168.0.1"},{"region":{"slug":"nyc3"},"droplet":{"id":2},"ip":"192.168.0.2"}]}`
	expected := []godo.FloatingIP{
		{Region: &godo.Region{Slug: "nyc3"}, Droplet: &godo.Droplet{ID: 1}, IP: "192.168.0.1"},
		{Region: &godo.Region{Slug: "nyc3"}, Droplet: &godo.Droplet{ID: 2}, IP: "192.168.0.2"},
	}
	apiServer(t, "/v2/floating_ips", resp, func() {
		floatingIP := doList()
		assert.Equal(t, floatingIP, expected, "they should be equal")
	})
}

func TestShow_doShow(t *testing.T) {
	resp := `{"floating_ip":{"region":{"slug":"nyc3"},"droplet":{"id":1},"ip":"192.168.0.1"}}`
	expected := &godo.FloatingIP{Region: &godo.Region{Slug: "nyc3"}, Droplet: &godo.Droplet{ID: 1}, IP: "192.168.0.1"}

	apiServer(t, "/v2/floating_ips/192.168.0.1", resp, func() {
		floatingIP := doShow("192.168.0.1")
		assert.Equal(t, floatingIP, expected, "they should be equal")
	})
}

func TestAssign_doUnassign(t *testing.T) {
	resp := `{"action":{"status":"in-progress", "region":{"slug":"nyc3"}}}`
	expected := &godo.Action{Status: "in-progress", Region: &godo.Region{Slug: "nyc3"}}

	apiServer(t, "/v2/floating_ips/192.168.0.1/actions", resp, func() {
		action := doUnassign("192.168.0.1")
		assert.Equal(t, action, expected, "they should be equal")
	})
}
