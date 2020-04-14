package main

import (
	"fmt"
	"math"
)

//GenDisplaceFn rturns a function which computes displacement as a function of time
func GenDisplaceFn(a, vo, so float64) func(float64) float64 {

	fn := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + vo*t + so
	}

	return fn

}

func main() {

	fn := GenDisplaceFn(10, 2, 1)

	fmt.Println(fn(3))

	fmt.Println(fn(5))
}
