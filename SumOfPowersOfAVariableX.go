//Sum of Powers of a Variable x

package main
 
import (
	"fmt"
	"math"
)


func SumOfPowersOfAVariableX(x float64, n int) float64 {
	sum := 0.0
	for i:=0; i < n; i++ {
    term := math.Pow(x, float64(i))
	sum+=term
	}
   return sum
}

func main() {

	x := 3.0
	n := 4
	result := SumOfPowersOfAVariableX(x, n)
	fmt.Println("the value:", result)
}