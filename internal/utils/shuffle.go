package utils

import "math/rand"

func ShuffleBytes(chars []byte) string {
	for i := len(chars) - 1; i > 0; i--{
		random_index := rand.Intn(i + 1)
		chars[i], chars[random_index] = chars[random_index], chars[i]
	}

	return string(chars)
}
