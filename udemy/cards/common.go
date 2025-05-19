package main

import "math/rand"

func ShuffleArray[T any](arr []T) []T {
	for _, item := range arr {
		randI := rand.Intn(len(arr))
		temp := arr[randI]

		arr[randI] = item
		item = temp
	}

	return arr
}
