package linear

import (
	"bytes"
	"fmt"
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

// Add adds the equation with srcIndex to the equation with index dstIndex 'count' times.
func (s1 System) Add(srcIndex int, dstIndex int, count int) (System, error) {
	if srcIndex >= len(s1) || srcIndex < 0 {
		return System{}, fmt.Errorf("source index %d is not present in the system", srcIndex)
	}

	if dstIndex >= len(s1) || dstIndex < 0 {
		return System{}, fmt.Errorf("destination index %d is not present in the system", dstIndex)
	}

	op := []Line(s1)
	src := op[srcIndex]
	for i := 0; i < count; i++ {
		vec, err := op[dstIndex].NormalVector.Add(src.NormalVector)
		if err != nil {
			return System{}, err
		}
		op[dstIndex].NormalVector = vec
		op[dstIndex].ConstantTerm += src.ConstantTerm
	}
	return System(op), nil
}

// Sub subtracts the equation with srcIndex from the equation with index dstIndex 'count' times.
func (s1 System) Sub(srcIndex int, dstIndex int, count int) (System, error) {
	if srcIndex >= len(s1) || srcIndex < 0 {
		return System{}, fmt.Errorf("source index %d is not present in the system", srcIndex)
	}

	if dstIndex >= len(s1) || dstIndex < 0 {
		return System{}, fmt.Errorf("destination index %d is not present in the system", dstIndex)
	}

	op := []Line(s1)
	src := op[srcIndex]
	for i := 0; i < count; i++ {
		vec, err := op[dstIndex].NormalVector.Sub(src.NormalVector)
		if err != nil {
			return System{}, err
		}
		op[dstIndex].NormalVector = vec
		op[dstIndex].ConstantTerm -= src.ConstantTerm
	}
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
