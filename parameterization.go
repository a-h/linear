package linear

import (
	"bytes"
	"fmt"
)

// Parameterization provides a way of enumerating the many possible solutions to a system of equations
// which has infinite solutions.
type Parameterization struct {
	Basepoint        Vector
	DirectionVectors []Vector
}

func (p1 Parameterization) String() string {
	buf := bytes.NewBufferString("{ ")

	buf.WriteString(fmt.Sprintf("Basepoint: %v, ", p1.Basepoint))

	buf.WriteString("Direction Vectors: ")
	buf.WriteString(" [ ")
	for i, e := range p1.DirectionVectors {
		buf.WriteString(e.String())
		if i < len(p1.DirectionVectors)-1 {
			buf.WriteString(", ")
		}
	}

	buf.WriteString(" ] }")
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
