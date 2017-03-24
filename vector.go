package linear

import (
	"bytes"
	"fmt"
	"math"

	tolerancepkg "github.com/a-h/linear/tolerance"
	"github.com/a-h/round"
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
	return v1.EqWithinTolerance(v2, DefaultTolerance)
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

// IsOrthogonalTo calculates whether the current vector is orthogonal to the input vector by calculating
// the dot product. If the dot product is zero, then the vectors are orthogonal.
func (v1 Vector) IsOrthogonalTo(v2 Vector) (bool, error) {
	f, err := v1.DotProduct(v2)
	if err != nil {
		return false, fmt.Errorf("error calculating whether the vectors are orthogonal: %v", err)
	}
	return tolerancepkg.IsWithin(f, 0, DefaultTolerance), nil
}

// Projection calculates the projection of the v2 vector onto the basis vector (v1) by calculating the unit vector of v1 and scaling it.
func (v1 Vector) Projection(v2 Vector) (Vector, error) {
	unitVectorOfBasis := v1.Normalize()
	dotProduct, err := v2.DotProduct(unitVectorOfBasis)
	if err != nil {
		return Vector{}, fmt.Errorf("error projecting %v onto %v with error: %v", v2, v1, err)
	}
	return unitVectorOfBasis.Scale(dotProduct), nil
}

// ProjectionOrthogonalComponent calculates the projection of v2 onto the basis vector (v1) and uses that to calculate a component
// which is orthogonal to the basis vector and perpendicular to v2.
func (v1 Vector) ProjectionOrthogonalComponent(v2 Vector) (Vector, error) {
	projection, err := v1.Projection(v2)
	if err != nil {
		return Vector{}, err
	}
	return v2.Sub(projection)
}

// Round rounds the vector to the specified number of places.
func (v1 Vector) Round(decimals int) Vector {
	op := make([]float64, len(v1))
	for i := 0; i < len(v1); i++ {
		op[i] = round.ToEven(v1[i], decimals)
	}
	return Vector(op)
}

// CrossProduct calculates the cross product of the current vector and the input vector. The cross product
// produces a vector which is:
//  - orthogonal to both v1 and v2.
//  - has a magnitude of the magnitude of v1 * the magnitude of v2 * the sine of the angle between v1 and v2
// v2 must be a vector with 3 dimensions.
func (v1 Vector) CrossProduct(v2 Vector) (Vector, error) {
	if len(v1) != 3 {
		return Vector{}, fmt.Errorf("the basis vector has %d dimensions but must have 3 because cross products do not generalize to multiple dimensions", len(v1))
	}

	if len(v2) != 3 {
		return Vector{}, fmt.Errorf("the input vector has %d dimensions but must have 3 because cross products do not generalize to multiple dimensions", len(v2))
	}

	var a1, a2, a3 = v1[0], v1[1], v1[2]
	var b1, b2, b3 = v2[0], v2[1], v2[2]

	c1 := (a2 * b3) - (a3 * b2)
	c2 := (a3 * b1) - (a1 * b3)
	if c2 == -0 {
		c2 = 0
	}
	c3 := (a1 * b2) - (a2 * b1)

	return NewVector(c1, c2, c3), nil
}

// AreaOfParallelogram calculates the area of a parallelogram spanned by the basis vector and input vector for 2D and 3D inputs.
func (v1 Vector) AreaOfParallelogram(v2 Vector) (float64, error) {
	// Add a z dimension initialised to zero for 2D inputs.
	if len(v1) == 2 {
		v1 = NewVector(v1[0], v1[0], 0)
	}
	if len(v2) == 2 {
		v2 = NewVector(v2[0], v2[1], 0)
	}

	cp, err := v1.CrossProduct(v2)
	if err != nil {
		return 0.0, err
	}
	return cp.Magnitude(), nil
}

// AreaOfTriangle calculates the area of a parallelogram spanned by the basis vector and input vector for 2D and 3D inputs.
func (v1 Vector) AreaOfTriangle(v2 Vector) (float64, error) {
	// Add a z dimension initialised to zero for 2D inputs.
	if len(v1) == 2 {
		v1 = NewVector(v1[0], v1[0], 0)
	}
	if len(v2) == 2 {
		v2 = NewVector(v2[0], v2[1], 0)
	}

	p, err := v1.AreaOfParallelogram(v2)
	if err != nil {
		return p, err
	}
	return p * 0.5, nil
}
