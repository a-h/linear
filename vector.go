package linear

import (
	"bytes"
	"fmt"
	"math"
)

// Vector represents an array of values.
type Vector []float64

// NewVector creates a vector with the dimensions specified by the argument.
func NewVector(values ...float64) Vector {
	return Vector(values)
}

func (v1 Vector) String() string {
	if len(v1) == 1 {
		return fmt.Sprintf("[%v]", v1[0])
	}
	buf := bytes.NewBufferString("[")
	for i, p := range v1 {
		buf.WriteString(fmt.Sprintf("%v", p))

		if i < len(v1)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")
	return buf.String()
}

// Eq compares an input vector against the current vector.
func (v1 Vector) Eq(v2 Vector) bool {
	if len(v1) != len(v2) {
		return false
	}
	for i := 0; i < len(v2); i++ {
		if v1[i] != v2[i] {
			return false
		}
	}
	return true
}

// Add adds the input vector to the current vector and returns a new vector.
func (v1 Vector) Add(v2 Vector) (Vector, error) {
	if len(v1) != len(v2) {
		return Vector{}, fmt.Errorf("cannot add vectors together because they have different dimensions (%d and %d)", len(v1), len(v2))
	}
	op := make([]float64, len(v1))
	for i := 0; i < len(v1); i++ {
		op[i] = v1[i] + v2[i]
	}
	return Vector(op), nil
}

// Sub subtracts the input vector from the current vector and returns a new vector.
func (v1 Vector) Sub(v2 Vector) (Vector, error) {
	if len(v1) != len(v2) {
		return Vector{}, fmt.Errorf("cannot subtract vectors because they have different dimensions (%d and %d)", len(v1), len(v2))
	}
	op := make([]float64, len(v1))
	for i := 0; i < len(v1); i++ {
		op[i] = v1[i] - v2[i]
	}
	return Vector(op), nil
}

// Mul muliplies the input vector and the current vector together and returns a new vector.
func (v1 Vector) Mul(v2 Vector) (Vector, error) {
	if len(v1) != len(v2) {
		return Vector{}, fmt.Errorf("cannot multiply vectors because they have different dimensions (%d and %d)", len(v1), len(v2))
	}
	op := make([]float64, len(v1))
	for i := 0; i < len(v1); i++ {
		op[i] = v1[i] * v2[i]
	}
	return Vector(op), nil
}

// Scale muliplies the current vector by the scalar input and returns a new vector.
func (v1 Vector) Scale(scalar float64) Vector {
	op := make([]float64, len(v1))
	for i := 0; i < len(v1); i++ {
		op[i] = v1[i] * scalar
	}
	return Vector(op)
}

// Magnitude calculates the magnitude of the vector by calculating the square root of
// the sum of each element squared.
func (v1 Vector) Magnitude() float64 {
	var sumOfSquares float64
	for _, v := range v1 {
		sumOfSquares += (v * v)
	}
	return math.Sqrt(sumOfSquares)
}
