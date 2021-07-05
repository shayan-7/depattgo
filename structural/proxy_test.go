package structural

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProxy(t *testing.T) {
	n := NewNginx()

	cases := []struct {
		description string
		url         string
		method      string
		status      int
		message     string
	}{
		{
			"Getting the status",
			"/status",
			"GET",
			200,
			"OK",
		},
		{
			"Getting the status by wrong method",
			"/status",
			"POST",
			404,
			"Not Found",
		},
		{
			"Getting the status out of the rate limit",
			"/status",
			"GET",
			403,
			"Not Allowed",
		},
		{
			"Posting the user",
			"/users",
			"POST",
			201,
			"User Created",
		},
		{
			"Posting the user by wrong method",
			"/users",
			"GET",
			404,
			"Not Found",
		},
		{
			"Getting the status out of the rate limit",
			"/users",
			"POST",
			403,
			"Not Allowed",
		},
		{
			"Invalid url",
			"/invalid/url",
			"PATCH",
			404,
			"Not Found",
		},
	}

	for _, c := range cases {
		s, m := n.HandleRequest(c.url, c.method)
		assert.Equal(t, c.status, s)
		assert.Equal(t, c.message, m)
	}
}
