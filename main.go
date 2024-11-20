package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"image"
	"image/color"
	"io"
	"math"
	"os"
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

// This is an error type for the Sqrt function to handle the case where the input value is negative
type ErrNegativeSqrt float64

// The method to implement the error on the ErrNegativeSqrt type
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot compute square root of negative number: %v", float64(e))
}

// An update of the Sqrt function to handle negative numbers
func SqrtV2(v float64) (float64, error) {
	if v < 0 {
		return 0, ErrNegativeSqrt(v)
	}
	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - v) / (2 * z)
	}
	return z, nil
}

// a simple struct
type MyReader struct{}

// Reader type that emits an infinite stream of the ASCII character 'A'
func (r MyReader) Read(p []byte) (int, error) {
	p[0] = 'A'
	return 1, nil
}

// Rot13 type that implements a reader
type rot13Reader struct {
	r io.Reader
}

// Rot13Reader implements io.Reader and reads from an io.Reader
// Modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.
func (r *rot13Reader) Read(p []byte) (int, error) {
	b, err := r.r.Read(p)
	if err != nil {
		return 0, err
	}
	for i := 0; i < b; i++ {
		if p[i] >= 'A' && p[i] <= 'Z' {
			p[i] = (p[i]-'A'+13)%26 + 'A'
		} else if p[i] >= 'a' && p[i] <= 'z' {
			p[i] = (p[i]-'a'+13)%26 + 'a'
		}
	}
	return b, nil
}

// Image type that implements the necessary methods for an Image builtin interface
type Image struct{}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 199, 199)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x int, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
        return
    }
    Walk(t.Left, ch)
    ch <- t.Value
    Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
    ch2 := make(chan int)
    
    go func() {
        Walk(t1, ch1)
        close(ch1)
    }()
    
    go func() {
        Walk(t2, ch2)
        close(ch2)
    }()
    
    // Compare the received values from the chanels
    for {
        v1, ok1 := <-ch1
        v2, ok2 := <-ch2
        
        if ok1 != ok2 || v1 != v2 {
            return false
        }
        if !ok1 {
            break
        }
    }
    return true
}

func main() {

	num := 16.0
	sqrt := Sqrt(num)
	fmt.Printf("The square root of %.2f is %.2f\n", num, sqrt)
	test := Pic(8, 10)
	for _, v := range test {
		fmt.Println(v)
	}

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	ch := make(chan int)
    go Walk(tree.New(1), ch)
    
    // Test of Walk
    for i := 0; i < 10; i++ {
        fmt.Println(<-ch)
    }
    
    // Test of Same
    fmt.Println(Same(tree.New(1), tree.New(1))) // doit afficher true
    fmt.Println(Same(tree.New(1), tree.New(2))) 
}
