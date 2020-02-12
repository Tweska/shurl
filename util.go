package main

import "math/rand"

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	randStr := make([]rune, n)
	for i := range randStr {
		randStr[i] = letters[rand.Intn(len(letters))]
	}

	return string(randStr)
}
