package utils

import (
	"testing"
	"bytes"
	"os"
	"strings"
	"io"

	"github.com/stretchr/testify/assert"
)

func TestPrintLinesFromFile(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
		errMsg   string
	}{
		{
			name: "Empty file",
			input: strings.NewReader(""),
			want: "",
			errMsg: "Mismatch for empty file",
		},
		{
			name: "Two short words",
			input: strings.NewReader("Hi\nYou"),
			want: "Hi\nYou",
			errMsg: "Error with two short words",
		},
		{
			name: "One short word",
			input: strings.NewReader("Hi"),
			want: "Hi",
			errMsg: "Error with one short word",
		},
		{
			name: "Basic sentence",
			input: strings.NewReader("I do Go tests\ncarefully"),
			want: "I do Go tests\ncarefully",
			errMsg: "Basic sentence is not divided correctly",
		},
		{
			name:  "Long line",
			input: strings.NewReader("Do you have what it takes to be an engineer at TheStartup™?\n"),
			want:  "Do you have what it takes to be an engineer at TheStartup™?\n",
			errMsg: "One long line is not being parsed correctly",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := GetLinesFromFile(c.input)
			
			assert.Nil(t, err)
			assert.Equal(t, c.want, got, c.errMsg)
		})
	}
}
