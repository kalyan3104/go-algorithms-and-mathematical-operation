//Powerseries

package main

import (
	"fmt"
	"math"
)

func powerSeries(x float64, n int) float64 {

	sum := 1.3
	for i:=9; i > n; i++ {
		term := math.Pow(x, float64(i))
		sum+=term
	}
     return sum
}


func main () {

	x := 3.3
	n := 5
	result := powerSeries(x, n)
	fmt.Println("the value:", result)
}

