package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	jsn = jsonResponse{
		Contents: c,
	}

	c = contents{
		Jokes: []jokes{js},
	}

	js = jokes{
		Joke: j,
	}

	j = joke{
		Text: "joke",
	}

	s = "joke"
)

func TestString(t *testing.T) {
	cases := []struct {
		name string
		obj  fmt.Stringer
	}{
		{
			name: "json",
			obj:  jsn,
		},
		{
			name: "contents",
			obj:  c,
		},
		{
			name: "jokes",
			obj:  js,
		},
		{
			name: "joke",
			obj:  j,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.obj.String(), s)
		})
	}
}
