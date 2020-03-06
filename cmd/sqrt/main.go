package main

import (
	"fmt"
	"math"
)

// Sqrt finds the nearest square root
func Sqrt(x float64) (z float64, counter int) {
	z = float64(1)
	previous := float64(0)

	for math.Abs(z-previous) > 0.00000000000001 {
		previous = z
		z -= (z*z - x) / (2 * z)
		counter++
	}

	return
}

func main() {
	fmt.Println(Sqrt(83))
}
