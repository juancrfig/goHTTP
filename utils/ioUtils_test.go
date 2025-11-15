package utils

import (
	"testing"
	"strings"
	"bytes"
	"io"

	"github.com/stretchr/testify/assert"
)

func TestPrintLinesFromFile(t *testing.T) {
	cases := []struct {
		name   string
		input  io.Reader
		want   string
	}{
		{
			name: "Empty file",
			input: strings.NewReader(""),
			want: "",
		},
		{
			name: "One word",
			input: strings.NewReader("Hello\n"),
			want: "read: Hello\n",
		},
		{
			name: "One long sentence",
			input: strings.NewReader("Do you have what it takes to be an engineer at TheStartup™?\n"),
			want: "read: Do you have what it takes to be an engineer at TheStartup™?\n",
		},
		{
			name: "Two lines",
			input: strings.NewReader("Do you have what it takes to be an engineer at TheStartup™?\nAre you willing to work 80 hours a week in hopes that your 0.001% equity is worth something?\n"),
			want: "read: Do you have what it takes to be an engineer at TheStartup™?\nread: Are you willing to work 80 hours a week in hopes that your 0.001% equity is worth something?\n",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var buf bytes.Buffer

			err := PrintLinesFromFile(c.input, &buf)
			if err != nil {
				t.Fatalf("Error: %v\n", err)
			}

			got := buf.String()

			assert.Equal(t, c.want, got)
		})
	}
}
