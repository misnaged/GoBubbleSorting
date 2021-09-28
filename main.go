package main

import (
	"fmt"
)

var ab = []int{0, 3, 8, 7, 2, 1, 5, 4, 9, 6}

func Bubb() {

	for i, _ := range ab {

		i2 := i + 1

		if i2 == len(ab) {
			break
		}
		if ab[i] > ab[i2] {

			ab[i], ab[i2] = ab[i2], ab[i]

		}

	}

}
func main() {
	fmt.Println(ab)
	for i, _ := range ab {
		if i == len(ab) {
			break
		}
		ab3 := append(ab[:i], ab[i:]...)

		if ab3[i] == ab[i] {
			Bubb()
		}

	}
	fmt.Println(ab)
}
