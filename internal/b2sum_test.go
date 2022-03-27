package internal

import (
	"bytes"
	"strings"
	"testing"
)

const kayos = `a2F5b3MKeGoC90IBWQPGxv2FJVLScpEvR0DhWEdhiobiF/cfVBnSXhAxr+5YUxOJZESTTrBLkDpoWxRIt1XVb3Aa/pvizg==`

func TestC264(t *testing.T) {
	soyak := C264(strings.NewReader("kayos\n"))
	if soyak != kayos {
		t.Fatalf("wanted: %s, got %s", kayos, soyak)
	}
	coyak := C264(bytes.NewReader([]byte("kayos\n")))
	if coyak != kayos {
		t.Fatalf("wanted: %s, got %s", kayos, coyak)
	}
	notsoyak := C264(strings.NewReader("billy\n"))
	if notsoyak == kayos {
		t.Fatalf("%s should not have been %s", notsoyak, kayos)
	}
}
