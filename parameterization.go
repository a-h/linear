package linear

import (
	"bytes"
	"fmt"

	"math"

	"github.com/a-h/linear/tolerance"
)

// Parameterization provides a way of enumerating the many possible solutions to a system of equations
// which has infinite solutions.
type Parameterization struct {
	Basepoint        Vector
	DirectionVectors []Vector
}

func (p1 Parameterization) String() string {
	buf := bytes.NewBufferString("{ ")

	for variableIndex, basepointValue := range p1.Basepoint {
		buf.WriteString(fmt.Sprintf("x%v = ", getSubscript(variableIndex+1)))

		var nonzero bool
		if !tolerance.IsWithin(basepointValue, 0, DefaultTolerance) {
			buf.WriteString(fmt.Sprintf("%v", basepointValue))
			nonzero = true
		}

		for directionIndex, directionVector := range p1.DirectionVectors {
			value := directionVector[variableIndex]
			if tolerance.IsWithin(value, 0, DefaultTolerance) {
				continue
			}
			var sign string
			if nonzero {
				if value < 0 {
					sign = " - "
				} else {
					sign = " + "
				}
			}
			var freeVariableCoefficient string
			if !tolerance.IsWithin(math.Abs(value), 1, DefaultTolerance) {
				freeVariableCoefficient = fmt.Sprintf("%v", math.Abs(value))
			}
			buf.WriteString(fmt.Sprintf("%v%v%v", sign, freeVariableCoefficient, getFreeVariableName(directionIndex)))
		}
		if variableIndex < len(p1.Basepoint)-1 {
			buf.WriteString(", ")
		}
	}

	buf.WriteString(" }")
	return buf.String()
}

func getFreeVariableName(index int) string {
	var name string
	if index%2 == 0 {
		name = "t"
		index = index / 2
	} else {
		name = "s"
		index = (index - 1) / 2
	}
	if index > 0 {
		name += getSubscript(index)
	}
	return name
}
