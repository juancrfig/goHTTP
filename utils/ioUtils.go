package utils

import (
	"io"
	"strings"
)

func PrintLinesFromFile(r io.Reader, w io.Writer) error {

	if w == nil {
		w = os.Stdout
	}

	return err
}
