package collatzconjecture

import (
	"errors"
	"fmt"
)

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n is negative")
	}

	fmt.Printf("n: %v\n", n)
	steps := 0
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}

		steps += 1
		// fmt.Println(steps, n)
	}

	return steps, nil

}
