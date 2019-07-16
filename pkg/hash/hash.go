package hash

import (
	"crypto/sha256"
	"github.com/cespare/xxhash"
	"github.com/pubgo/errors"

	"golang.org/x/crypto/ripemd160"
)

func Sha256(bytes []byte) []byte {
	h := sha256.New()
	defer h.Reset()

	h.Write(bytes)
	return h.Sum(nil)
}

func Ripemd160(bytes []byte) []byte {
	h := ripemd160.New()
	defer h.Reset()

	h.Write(bytes)
	return h.Sum(nil)
}

func XXHash(bytes []byte) []byte {
	h := xxhash.New()
	defer h.Reset()

	_, err := h.Write(bytes)
	errors.Panic(err)

	return h.Sum(nil)
}
