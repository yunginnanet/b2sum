package internal

import (
	"bytes"
	"strings"
	"testing"
)

const kayos = `Kr+6ONDx+cq/WvhHpQE/4LVuJYi9QHz1TztHNTWwa9KJWqHxfTNLKF3YxrcLptA3wO0KHm83Lq7gpBWgCQzPag==`

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
