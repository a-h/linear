package linear

import (
	"bytes"
	"fmt"
)

// Vector represents an array of values.
type Vector []float64

// NewVector creates a vector with the dimensions specified by the argument.
func NewVector(values ...float64) Vector {
	return Vector(values)
}

func (v Vector) String() string {
	if len(v) == 1 {
		return fmt.Sprintf("[%v]", v[0])
	}
	buf := bytes.NewBufferString("[")
	for i, p := range v {
		buf.WriteString(fmt.Sprintf("%v", p))

		if i < len(v)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")
	return buf.String()
}
