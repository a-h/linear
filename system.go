package linear

import (
	"bytes"
	"errors"
	"fmt"
	"sort"

	"github.com/a-h/linear/tolerance"
)

// System defines a system of equations, where each equation
// is a linear equation.
type System []Line

// NewSystem creates a new system of lines.
func NewSystem(lines ...Line) System {
	return System(lines)
}

// String writes out each equation in the system, delineated by commas and
// surrounded by braces, e.g. { 1x₁ + 2x₂ + 3x₃ = 4, 5x₁ + 6x₂ + 7x₃ = 8 }
func (s1 System) String() string {
	buf := bytes.NewBufferString("{ ")

	for i, e := range s1 {
		buf.WriteString(e.String())
		if i < len(s1)-1 {
			buf.WriteString(", ")
		}
	}

	buf.WriteString(" }")
	return buf.String()
}

// Eq determies whether two systems are equal.
func (s1 System) Eq(s2 System) (bool, error) {
	if len(s1) != len(s2) {
		return false, nil
	}

	for i, e1 := range s1 {
		e2 := s2[i]
		equals, err := e1.Eq(e2)

		if err != nil || !equals {
			return false, err
		}
	}

	return true, nil
}

// Swap swaps elements in the system.
func (s1 System) Swap(a int, b int) (System, error) {
	if a >= len(s1) || a < 0 {
		return System{}, fmt.Errorf("index %d is not present in the system", a)
	}

	if b >= len(s1) || b < 0 {
		return System{}, fmt.Errorf("index %d is not present in the system", b)
	}

	op := []Line(s1)
	op[a], op[b] = op[b], op[a]
	return System(op), nil
}

// Multiply multiplies an equation by a coefficient.
func (s1 System) Multiply(index int, coefficient float64) (System, error) {
	if index >= len(s1) || index < 0 {
		return System{}, fmt.Errorf("index %d is not present in the system", index)
	}

	op := []Line(s1)
	op[index].NormalVector = op[index].NormalVector.Scale(coefficient)
	op[index].ConstantTerm *= coefficient
	return System(op), nil
}

// Add adds the equation with srcIndex multiplied by the coefficient to the equation with index dstIndex.
func (s1 System) Add(srcIndex int, dstIndex int, coefficient int) (System, error) {
	if srcIndex >= len(s1) || srcIndex < 0 {
		return System{}, fmt.Errorf("source index %d is not present in the system", srcIndex)
	}

	if dstIndex >= len(s1) || dstIndex < 0 {
		return System{}, fmt.Errorf("destination index %d is not present in the system", dstIndex)
	}

	op := []Line(s1)
	src := op[srcIndex]
	multipliedVector := src.NormalVector.Scale(float64(coefficient))

	dst := op[dstIndex]
	destinationVector, err := dst.NormalVector.Add(multipliedVector)
	if err != nil {
		return op, err
	}
	op[dstIndex].NormalVector = destinationVector
	op[dstIndex].ConstantTerm += src.ConstantTerm * float64(coefficient)
	return System(op), nil
}

// FindFirstNonZeroCoefficients finds the indices of the first non-zero coefficient of each equation in the
// system.
func (s1 System) FindFirstNonZeroCoefficients() (indices []int, err error) {
	indices = make([]int, len(s1))
	for i, e := range s1 {
		idx, _, ok := e.FirstNonZeroCoefficient()
		if !ok {
			return indices, fmt.Errorf("failed to find a non-zero coefficient for equation at index %d - %v", i, e)
		}
		indices[i] = idx
	}
	return indices, nil
}

// TriangularForm organises the system by leading term.
func (s1 System) TriangularForm() (System, error) {
	sort.Sort(NonZeroLeadingTerms(s1))
	return s1, nil
}

// IsTriangularForm determines whether the system is in triangular form, where the top row starts with a non-zero
// term, the next one down starts with a zero etc., the one after that starts with two zero terms etc.
func (s1 System) IsTriangularForm() (bool, error) {
	for i, e := range s1 {
		if len(s1) != len(e.NormalVector) {
			return false, errors.New("all equations in a system need to have the same number of terms")
		}
		if tolerance.IsWithin(e.NormalVector[i], 0, DefaultTolerance) {
			return false, nil
		}
	}
	return true, nil
}

// NonZeroLeadingTerms sorts the system into triangular form, where the top row starts with the earliest non-zero
// term, the next one down starts with a zero etc.
type NonZeroLeadingTerms System

func (tf NonZeroLeadingTerms) Len() int      { return len(tf) }
func (tf NonZeroLeadingTerms) Swap(i, j int) { tf[i], tf[j] = tf[j], tf[i] }
func (tf NonZeroLeadingTerms) Less(i, j int) bool {
	fnz1, _, ok := tf[i].FirstNonZeroCoefficient()
	if !ok {
		fnz1 = len(tf[i].NormalVector)
	}
	fnz2, _, ok := tf[j].FirstNonZeroCoefficient()
	if !ok {
		fnz2 = len(tf[j].NormalVector)
	}
	return fnz1 < fnz2
}
