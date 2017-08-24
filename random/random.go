package random

import (
	"crypto/rand"
	"encoding/base64"
)

func RandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func RandomString(s int) (string, error) {
	b, err := RandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
