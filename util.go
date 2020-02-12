package main

import "math/rand"

// RandomString will return a random string with characters and numbers of the given length.
func RandomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	randStr := make([]rune, n)
	for i := range randStr {
		randStr[i] = letters[rand.Intn(len(letters))]
	}

	return string(randStr)
}
