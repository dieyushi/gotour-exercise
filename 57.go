/*
Exercise: Errors

Copy your Sqrt function from the earlier exercises and modify it to return an error value.

Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

Create a new type

type ErrNegativeSqrt float64
and make it an error by giving it a

func (e ErrNegativeSqrt) Error() string
method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".

Note: a call to fmt.Print(e) inside the Error method will send the program into an infinite loop. 
You can avoid this by converting e first: fmt.Print(float64(e)). Why?

Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.
*/

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"math"
// )

// func Sqrt(f float64) (float64, error) {
// 	if f < 0 {
// 		e := errors.New("cannot Sqrt negative number")
// 		return f, e
// 	}
// 	return math.Sqrt(f), nil
// }

// func main() {
// 	fmt.Println(Sqrt(2))
// 	fmt.Println(Sqrt(-2))
// }

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number ", float64(e))
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return f, ErrNegativeSqrt(f)
	}
	return math.Sqrt(f), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
