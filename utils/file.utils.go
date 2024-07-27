package utils

import (
	"math/rand"
	"mime/multipart"
)

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandomString(n int) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func FileValidation(fileHeader *multipart.FileHeader, fileType []string) bool {
	contentType := fileHeader.Header.Get("Content-Type")
	result := false

	for _, typeF := range fileType {
		if contentType == typeF {
			result = true
			break
		}
	}

	return result
}
