package utils

import "math/rand"

func RandomRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
