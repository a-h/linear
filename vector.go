package linear

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/a-h/linear/bigfloat"
)

// Vector represents an array of values.
type Vector []*big.Float

// Define a few constants to save allocations.
var floatNegativeOne = newFloat().SetInt64(-1)
var floatZero = newFloat().SetInt64(0)
var floatOne = newFloat().SetInt64(1)

const prec = 53

func newFloat() *big.Float {
	return new(big.Float).SetPrec(prec)
}

// NewVector creates a vector with the dimensions specified by the argument.
func NewVector(values ...float64) Vector {
	v := make([]*big.Float, len(values))
	for i, val := range values {
		v[i] = big.NewFloat(val)
		v[i].SetPrec(prec)
	}
	return Vector(v)
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
		if v1[i].Cmp(v2[i]) != 0 {
			return false
		}
	}
	return true
}

func (v1 Vector) EqTest(n string, v2 Vector) bool {
	if len(v1) != len(v2) {
		fmt.Printf("%s: length not equal\n", n)
		return false
	}
	for i := 0; i < len(v2); i++ {
		if v1[i].Cmp(v2[i]) != 0 {
			fmt.Printf("%s: value index %d not equal %v and %v\n", n, i, v1[i], v2[i])
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
	op := make([]*big.Float, len(v1))
	for i := 0; i < len(v1); i++ {
		op[i] = newFloat().Add(v1[i], v2[i])
	}
	return Vector(op), nil
}

// Sub subtracts the input vector from the current vector and returns a new vector.
func (v1 Vector) Sub(v2 Vector) (Vector, error) {
	if len(v1) != len(v2) {
		return Vector{}, fmt.Errorf("cannot subtract vectors because they have different dimensions (%d and %d)", len(v1), len(v2))
	}
	op := make([]*big.Float, len(v1))
	for i := 0; i < len(v1); i++ {
		op[i] = newFloat().Sub(v1[i], v2[i])
	}
	return Vector(op), nil
}

// Mul muliplies the input vector and the current vector together and returns a new vector.
func (v1 Vector) Mul(v2 Vector) (Vector, error) {
	if len(v1) != len(v2) {
		return Vector{}, fmt.Errorf("cannot multiply vectors because they have different dimensions (%d and %d)", len(v1), len(v2))
	}
	op := make([]*big.Float, len(v1))
	for i := 0; i < len(v1); i++ {
		op[i] = newFloat().Mul(v1[i], v2[i])
	}
	return Vector(op), nil
}

// Scale muliplies the current vector by the scalar input and returns a new vector.
func (v1 Vector) Scale(scalar *big.Float) Vector {
	op := make([]*big.Float, len(v1))
	for i := 0; i < len(v1); i++ {
		op[i] = newFloat().Mul(v1[i], scalar)
	}
	return Vector(op)
}

// Magnitude calculates the magnitude of the vector by calculating the square root of
// the sum of each element squared.
func (v1 Vector) Magnitude() *big.Float {
	sumOfSquares := newFloat()
	for _, v := range v1 {
		squared := newFloat().Mul(v, v)
		sumOfSquares.Add(sumOfSquares, squared)
	}

	return bigfloat.Sqrt(sumOfSquares)
}

// Normalize normalizes the magnitude of a vector to 1 and returns a new vector.
func (v1 Vector) Normalize() Vector {
	mag := v1.Magnitude()
	if mag.Cmp(floatZero) == 0 {
		zeroes := make([]float64, len(v1))
		return NewVector(zeroes...) // Return a vector of zeroes if the magnitude is zero.
	}
	// one / mag
	scalar := newFloat().Quo(floatOne, mag)
	return v1.Scale(scalar)
}

// IsZeroVector returns true if all of the values in the vector are zero.
func (v1 Vector) IsZeroVector() bool {
	for _, v := range v1 {
		if v.Cmp(floatZero) != 0 {
			return false
		}
	}
	return true
}

// DotProduct calculates the dot product of the current vector and the input vector, or an error if the dimensions
// of the vectors do not match.
func (v1 Vector) DotProduct(v2 Vector) (*big.Float, error) {
	rv := newFloat()
	if len(v1) != len(v2) {
		return rv, fmt.Errorf("cannot calculate the dot product of the vectors because they have different dimensions (%d and %d)", len(v1), len(v2))
	}
	for i := 0; i < len(v1); i++ {
		rv.Add(rv, newFloat().Mul(v1[i], v2[i]))
	}
	return rv, nil
}

// AngleBetween returns the angle (in radians) between the current vector and v2, or an error if the dimensions of
// the vectors do not match.
func (v1 Vector) AngleBetween(v2 Vector) (Radian, error) {
	dp, err := v1.DotProduct(v2)
	if err != nil {
		return 0, err
	}

	magTimesMag := newFloat().Mul(v1.Magnitude(), v2.Magnitude())
	angle := newFloat().Quo(dp, magTimesMag)

	r, _ := bigfloat.Acos(angle).Float64()
	return Radian(r), nil
}

// IsParallelTo calculates whether the current vector is parallel to the input vector by normalizing both
// vectors, and comparing them. In the case that the
func (v1 Vector) IsParallelTo(v2 Vector) (bool, error) {
	if len(v1) != len(v2) {
		return false, fmt.Errorf("cannot calculate whether the vectors are parallel because they have different dimensions (%d and %d)", len(v1), len(v2))
	}
	if v1.IsZeroVector() || v2.IsZeroVector() {
		return true, nil
	}
	u1 := v1.Normalize()
	u2 := v2.Normalize()

	parallelAndSameDirection := u1.Eq(u2)
	parallelAndOppositeDirection := func() bool { return u1.Eq(u2.Scale(floatNegativeOne)) }

	return parallelAndSameDirection || parallelAndOppositeDirection(), nil
}

// IsOrthogonalTo calculates whether the current vector is orthogonol to the input vector by calculating
// the dot product. If the dot product is zero, then the vectors are orthogonol.
func (v1 Vector) IsOrthogonalTo(v2 Vector) (bool, error) {
	f, err := v1.DotProduct(v2)
	if err != nil {
		return false, fmt.Errorf("error calculating whether the vectors are orthogonol: %v", err)
	}
	return f.Cmp(floatZero) == 0, nil
}
