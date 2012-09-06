/*
Advanced Exercise: Complex cube roots

Let's explore Go's built-in support for complex numbers via the complex64 and complex128 types. 
For cube roots, Newton's method amounts to repeating:
z = z-(z^3-x)/(3*z^2)

Find the cube root of 2, just to make sure the algorithm works. There is a Pow function in the math/cmplx package.
*/

package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := complex128(1)
	for i := 0; i < 100; i++ {
		z = z - (cmplx.Pow(z, 3)-x)/(3*cmplx.Pow(z, 2))
	}
	return z
}

func main() {
	fmt.Println(Cbrt(2))
}
