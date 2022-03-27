package internal

import (
	"encoding/base64"
	"io"
	"os"

	"golang.org/x/crypto/blake2b"
)

// C264 read all data from i and computes a blake2b checksum and returns said checksum encoded with base64.
func C264(i io.Reader) string {
	hashr, _ := blake2b.New(64, nil)
	cytes, err := io.ReadAll(i)
	hashr.Write(cytes)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	return base64.StdEncoding.EncodeToString(hashr.Sum(nil))
}
