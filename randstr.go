package randstr

import (
	"crypto/rand"
)

const defaultCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func getBytes(n uint) ([]byte, error) {
	bs := make([]byte, n)
	if _, err := rand.Read(bs); err != nil {
		return nil, err
	}
	return bs, nil
}

// String generates a random string with n length with default charset.
func String(n uint) (string, error) {
	return StringWithCharset(n, defaultCharset)
}

// StringWithCharset generates a random string with n length with charset given.
func StringWithCharset(n uint, charset string) (string, error) {
	lenCharset := len(charset)
	bs, err := getBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bs {
		bs[i] = charset[b%byte(lenCharset)]
	}
	return string(bs), nil
}
