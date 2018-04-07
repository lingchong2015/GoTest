package mathoperator

import (
	"math"
	"fmt"
)

type ErrNegstiveSqrt struct {
	x float64
}

func (e *ErrNegstiveSqrt) Error() string {
	// Note: a call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop.
	return fmt.Sprintf("Can't Sqrt negative number: %v", e.x)
}

func Add(x int, y int) int {
	return x + y
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, &ErrNegstiveSqrt{
			x,
		}
	}
	z := 1.0
	for ; math.Abs(z * z - x) > 0.000001; {
		z -= (z * z - x) / (2 * z)
	}

	return z, nil
}
