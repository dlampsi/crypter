package crypter

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"io"
	mrand "math/rand"
	"time"
)

// GenerateRandString Generates random string selected length.
// If length is negative, returnd empty string.
func GenerateRandString(length int) string {
	if length < 0 {
		return ""
	}
	mrand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789=!-?*/\\")
	buf := make([]rune, length)
	for i := range buf {
		buf[i] = chars[mrand.Intn(len(chars))]
	}
	return string(buf)
}

// Encrypt data.
func Encrypt(data []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(crand.Reader, nonce); err != nil {
		return nil, err
	}
	return gcm.Seal(nonce, nonce, data, nil), nil
}

// Decrypt data.
func Decrypt(data []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, err
	}
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
