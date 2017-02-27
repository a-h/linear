package linear

import (
	"bytes"
	"fmt"
	"math"
	"strconv"

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

func (l1 Line) String() string {
	buf := bytes.Buffer{}
	for i, p := range l1.NormalVector {
		if i == 0 {
			// The first element should have an integrated +/- sign.
			if !tolerance.IsWithin(p, 0.0, DefaultTolerance) {
				// If the value is not zero, write it out.
				buf.WriteString(fmt.Sprintf("%v", p))
			}
			// Write out the x_1 specifier.
			buf.WriteString(fmt.Sprintf("x%s", getSubscript(i+1)))
			continue
		}

		// For anything after index zero, the sign becomes the operator.
		buf.WriteString(operator(p))

		// Write out the absolute value if it's non-zero.
		if !tolerance.IsWithin(p, 0.0, DefaultTolerance) {
			buf.WriteString(fmt.Sprintf("%v", math.Abs(p)))
		}
		// Write out the x_1 specifier.
		buf.WriteString(fmt.Sprintf("x%s", getSubscript(i+1)))
	}
	// Write out the constant term.
	buf.WriteString(fmt.Sprintf(" = %v", l1.ConstantTerm))
	return buf.String()
}

func operator(v float64) string {
	if v < 0 {
		return " - "
	}
	return " + "
}

func getSubscript(i int) string {
	buf := bytes.Buffer{}
	str := strconv.Itoa(i)
	for _, v := range str {
		switch v {
		case '0':
			buf.WriteRune('₀')
		case '1':
			buf.WriteRune('₁')
		case '2':
			buf.WriteRune('₂')
		case '3':
			buf.WriteRune('₃')
		case '4':
			buf.WriteRune('₄')
		case '5':
			buf.WriteRune('₅')
		case '6':
			buf.WriteRune('₆')
		case '7':
			buf.WriteRune('₇')
		case '8':
			buf.WriteRune('₈')
		case '9':
			buf.WriteRune('₉')
		}
	}

	return buf.String()
}

// IsParallelTo determines whether two lines are parallel to each other.
func (l1 Line) IsParallelTo(l2 Line) bool {
	return l1.NormalVector.Normalize().Eq(l2.NormalVector.Normalize())
}
