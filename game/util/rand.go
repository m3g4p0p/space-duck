package util

import "math/rand"

func RandIntMN(m, n int) int {
	return rand.Intn(n-m) + m
}
