package linear

import (
	"bytes"
	"fmt"
	"math"

	"github.com/a-h/linear/tolerance"
)

// Line consists of a normal vector which specifies the direction, and a constant term.
// The basepoint is calculated by the NewLine function.
type Line struct {
	Basepoint    Vector
	NormalVector Vector
	ConstantTerm float64
}

// NewLine creates a new Line based on the variable terms, e.g.:
// Ax + By = C
// For example, for a 2D vector (2, 3) and constant 5 would result in 2x + 3y = 5
func NewLine(normalVector Vector, constantTerm float64) Line {
	basepointVector := make([]float64, len(normalVector))
	nonZeroIndex := firstNonZeroElement(normalVector)
	basepointVector[nonZeroIndex] = constantTerm / normalVector[nonZeroIndex]

	return Line{
		ConstantTerm: constantTerm,
		Basepoint:    basepointVector,
		NormalVector: normalVector,
	}
}

func firstNonZeroElement(v Vector) (index int) {
	for i := 0; i < len(v); i++ {
		if !tolerance.IsWithin(v[i], 0, DefaultTolerance) {
			return i
		}
	}

	return 0
}

func (l Line) String() string {
	buf := bytes.Buffer{}
	for i, p := range l.NormalVector {
		if i == 0 {
			// The first element should have an integrated +/- sign.
			if !tolerance.IsWithin(p, 0.0, DefaultTolerance) {
				// If the value is not zero, write it out.
				buf.WriteString(fmt.Sprintf("%v", p))
			}
			// Write out the x_1 specifier.
			buf.WriteString(fmt.Sprintf("x_%d", i+1))
			continue
		}

		// For anything after index zero, the sign becomes the operator.
		buf.WriteString(operator(p))

		// Write out the absolute value if it's non-zero.
		if !tolerance.IsWithin(p, 0.0, DefaultTolerance) {
			buf.WriteString(fmt.Sprintf("%v", math.Abs(p)))
		}
		// Write out the x_1 specifier.
		buf.WriteString(fmt.Sprintf("x_%d", i+1))
	}
	// Write out the constant term.
	buf.WriteString(fmt.Sprintf(" = %v", l.ConstantTerm))
	return buf.String()
}

func operator(v float64) string {
	if v < 0 {
		return " - "
	}
	return " + "
}
