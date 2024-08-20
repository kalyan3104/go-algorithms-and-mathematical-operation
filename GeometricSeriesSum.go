//Geometric Series Sum

package main

import (
	"fmt"
	"math"
)


func GeometricSeriesSum(a, r float64, n int) float64 {

	sum := 0.0
	for i:=0; i < n; i++ {
		term := a * math.Pow(r, float64(i))
		sum+=term
	}
	return sum
}

func main () {
    a := 2.0
	r := 2.0
    n := 5
	
	result := GeometricSeriesSum(a, r, n)
	fmt.Printf("The sum of the first %d terms of the geometric series is: %.2f\n", n, result)
}