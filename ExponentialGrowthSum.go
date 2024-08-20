//Exponential Growth Sum

package main

import(
	"fmt"
	"math"
)

func ExponentialGrowthSum(n int) float64 {
 
	sum := 0.0
	for i:=0; i < n; i++ {
		term := math.Exp(float64(i)) // where e is Euler's number (~2.718).
		sum+=term
	}
	return sum // Sum=e^0+e^1+e^2+e^3
}

func main () {

	n := 4
	result := ExponentialGrowthSum(n)
	fmt.Printf("The sum of the first %d terms of exponential growth is: %.2f\n", n, result)
}
