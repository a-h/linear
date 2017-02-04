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

// Eq compares an input vector against the current vector.
func (v Vector) Eq(cmp Vector) bool {
	return VectorsAreEqual(v, cmp)
}

// VectorsAreEqual returns whether two Vectors are equal.
func VectorsAreEqual(v1 Vector, v2 Vector) bool {
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
