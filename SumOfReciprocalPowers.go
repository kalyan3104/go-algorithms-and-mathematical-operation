// Sum of Reciprocal Powers

package main

import (
	"fmt"
	"math"
)

func sumOfReciprocalPowers(x float64, n int) float64 {
	sum := 0.0
	for i := 0; i < n; i++ {
		term := 1 / math.Pow(x, float64(i)) // hole sum of math.Pow will divied by 1/x^n
		sum += term
	}
	return sum
}

func main() {
	x := 2.0
	n := 4
	result := sumOfReciprocalPowers(x, n)
	fmt.Printf("The sum of the first %d reciprocal powers of %.1f is: %.4f\n", n, x, result)
}
