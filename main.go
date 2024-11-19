package main

import (
	"fmt"
	"math"
	"strings"
)

/*
Calculate the square root of a number using the Newton Raphson method

Input:
- x: the number to find the square root of

Output: the square root of the number

Examples:

	    Sqrt(4) => 2.0
		Sqrt(10) => 3.1622776601683795
*/
func Sqrt(x float64) float64 {
	if math.IsNaN(x) {
		return math.NaN()
	}
	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

/*
Pic returns a slice of length dy that contains a slice of uint8 slices of dx values.
*/
func Pic(dx, dy int) [][]uint8 {
	picValue := make([][]uint8, dy)

	// Iterate over the rows of the picture
	for i := range picValue {
		// Create a slice for each row with length equal to dx
		// and fill it with the value of dx
		picValue[i] = make([]uint8, dx)
		for j := range picValue[i] {
			picValue[i][j] = uint8(dx)
		}
	}

	return picValue

}

/*
return a map of the counts of each “word” in the string s
*/
func WordCount(s string) map[string]int {
	stringMap := make(map[string]int)
	for _, v := range strings.Split(s, " ") {
		stringMap[v]++
	}
	return stringMap
}

// fibonacci returns a function that generates successive Fibonacci numbers (https://en.wikipedia.org/wiki/Fibonacci_sequence).
// Each call to the returned function yields the next number in the Fibonacci sequence.
func Fibonacci() func() int {
	firstNum := 0
	secondNum := 1
	return func() int {
		firstNum, secondNum = secondNum, firstNum+secondNum
		return firstNum
	}
}

// implement fmt.Stringer to print the address as a dotted quad
type IPAddr [4]byte

// This function is a Stringer that prints the address as a dotted quad
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {

	num := 16.0
	sqrt := Sqrt(num)
	fmt.Printf("The square root of %.2f is %.2f\n", num, sqrt)
	for i := 0; i < 1000; i++ {
		fmt.Printf("The iteration %v times \n", i)
	}
	test := Pic(8, 10)
	for _, v := range test {
		fmt.Println(v)
	}
}
