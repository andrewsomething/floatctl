package main

import (
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
