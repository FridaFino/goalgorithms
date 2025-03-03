// This is a pure Go implementation of the bogosort algorithm,
// also known as permutation sort, stupid sort, slowsort, shotgun sort, or monkey sort.
// Bogosort generates random permutations until it guesses the correct one.

// More info on: https://en.wikipedia.org/wiki/Bogosort

package sort

import (
	"math/rand"

	"github.com/FridaFino/goalgorithms/constraints"
)

func isSorted[T constraints.Number](arr []T) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}

	return true
}

func shuffle[T constraints.Number](arr []T) {
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func Bogo[T constraints.Number](arr []T) []T {
	for !isSorted(arr) {
		shuffle(arr)
	}

	return arr
}
