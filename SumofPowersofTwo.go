// Sum of Powers of 2

package main

import (

	"fmt"
	"math"
)

func sumOfPowersOfTwo(n int) float64 {
	sum := 0.0
	for i:=0; i < n; i++ {
		term := math.Pow(2, float64(i))
		sum+=term
	}
	return sum
}

func main () {

	n := 5
	result := sumOfPowersOfTwo(n)
	fmt.Printf("The sum of the first %d powers of 2 is: %.2f\n", n, result)
}