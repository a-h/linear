package linear

import (
	"bytes"
	"fmt"
	"math"

	tolerancepkg "github.com/a-h/linear/tolerance"
)

// Vector represents an array of values.
type Vector []float64

// NewVector creates a vector with the dimensions specified by the argument.
func NewVector(values ...float64) Vector {
	return Vector(values)
}

// DefaultTolerance is the tolerance to use for comparisons.
const DefaultTolerance = 1e-10

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

// EqWithinTolerance tests that a vector is equal, within a given tolerance.
func (v1 Vector) EqWithinTolerance(v2 Vector, tolerance float64) bool {
	if len(v1) != len(v2) {
		return false
	}
	for i := 0; i < len(v2); i++ {
		if !tolerancepkg.IsWithin(v1[i], v2[i], tolerance) {
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

// Normalize normalizes the magnitude of a vector to 1 and returns a new vector.
func (v1 Vector) Normalize() Vector {
	mag := v1.Magnitude()
	if mag == 0 {
		return Vector(make([]float64, len(v1))) // Return a vector of zeroes if the magnitude is zero.
	}
	return v1.Scale(float64(1.0) / mag)
}

// IsZeroVector returns true if all of the values in the vector are within tolerance of zero.
func (v1 Vector) IsZeroVector() bool {
	for _, v := range v1 {
		if !tolerancepkg.IsWithin(v, 0, DefaultTolerance) {
			return false
		}
	}
	return true
}

// DotProduct calculates the dot product of the current vector and the input vector, or an error if the dimensions
// of the vectors do not match.
func (v1 Vector) DotProduct(v2 Vector) (float64, error) {
	var rv float64
	if len(v1) != len(v2) {
		return rv, fmt.Errorf("cannot calculate the dot product of the vectors because they have different dimensions (%d and %d)", len(v1), len(v2))
	}
	for i := 0; i < len(v1); i++ {
		rv += v1[i] * v2[i]
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
	return Radian(math.Acos(dp / (v1.Magnitude() * v2.Magnitude()))), nil
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

	parallelAndSameDirection := u1.EqWithinTolerance(u2, DefaultTolerance)
	parallelAndOppositeDirection := func() bool { return u1.EqWithinTolerance(u2.Scale(-1), DefaultTolerance) }

	return parallelAndSameDirection || parallelAndOppositeDirection(), nil
}

// IsOrthogonalTo calculates whether the current vector is orthogonol to the input vector by calculating
// the dot product. If the dot product is zero, then the vectors are orthogonol.
func (v1 Vector) IsOrthogonalTo(v2 Vector) (bool, error) {
	f, err := v1.DotProduct(v2)
	if err != nil {
		return false, fmt.Errorf("error calculating whether the vectors are orthogonol: %v", err)
	}
	return tolerancepkg.IsWithin(f, 0, DefaultTolerance), nil
}

// Project projects the v2 vector onto the basis vector (v1) by calculating the unit vector of v1 and scaling it.
func (v1 Vector) Project(v2 Vector) (Vector, error) {
	unitVectorOfBasis := v1.Normalize()
	dotProduct, err := v2.DotProduct(unitVectorOfBasis)
	if err != nil {
		return Vector{}, fmt.Errorf("error projecting %v onto %v with error: %v", v2, v1, err)
	}
	return unitVectorOfBasis.Scale(dotProduct), nil
}
