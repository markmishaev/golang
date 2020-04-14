package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

//GenDisplaceFn rturns a function which computes displacement as a function of time
func GenDisplaceFn(a, vo, so float64) func(float64) float64 {

	fn := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + vo*t + so
	}

	return fn
}

func main() {

	in := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter acceleration: ")
	in.Scan()
	acceleration := in.Text()
	numAcceleration, _ := strconv.ParseFloat(acceleration, 64)

	fmt.Print("Enter initial velocity: ")
	in.Scan()
	velocity := in.Text()
	numVelocity, _ := strconv.ParseFloat(velocity, 64)

	fmt.Print("Enter initial displacement : ")
	in.Scan()
	displacement := in.Text()
	numDisplacement, _ := strconv.ParseFloat(displacement, 64)

	fmt.Print("Enter time : ")
	in.Scan()
	time := in.Text()
	numTime, _ := strconv.ParseFloat(time, 64)

	fn := GenDisplaceFn(numAcceleration, numVelocity, numDisplacement)

	fmt.Println(fn(numTime))

}
