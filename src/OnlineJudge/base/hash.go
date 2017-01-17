package base

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

const SaltSize = 32

func GenHash(secret []byte) ([]byte, error) {
	buf := make([]byte, SaltSize, SaltSize+sha256.Size)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		return nil, err
	}
	h := sha256.New()
	h.Write(buf)
	h.Write(secret)
	return h.Sum(buf), nil
}

func MatchHash(data, secret []byte) bool {
	if len(data) != SaltSize+sha256.Size {
		return false
	}
	h := sha256.New()
	h.Write(data[:SaltSize])
	h.Write(secret)
	return bytes.Equal(h.Sum(nil), data[SaltSize:])
}

func MD5Hash(data []byte) []byte {
	h := md5.New()
	h.Write(data)
	return h.Sum(nil)
}
