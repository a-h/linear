package linear

import (
	"bytes"
	"errors"
	"fmt"

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
	op[index] = op[index].Scale(coefficient)
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
	// Copy the input to a new value.
	op := s1

	if !s1.AllEquationsHaveSameNumberOfTerms() {
		return op, errors.New("all equations in a system need to have the same number of terms")
	}

	// Iterate through and elimate each term in order.
	var termIndex int
	for i := 0; i < len(op)-1; i++ {
		currentCoefficient := op[i].NormalVector[termIndex]
		if tolerance.IsWithin(currentCoefficient, 0, DefaultTolerance) {
			// Swap the current equation with the first one below it that has a non-zero coefficient for the term.
			for j := i + 1; j < len(op); j++ {
				nextEquation := op[j].NormalVector
				nextCoefficient := nextEquation[termIndex]

				if !tolerance.IsWithin(nextCoefficient, 0, DefaultTolerance) {
					op, _ = op.Swap(i, j)
					break
				}
			}
		}

		// Apply the cancellation to all subsequent equations.
		currentEquation := op[i]
		currentCoefficient = currentEquation.NormalVector[termIndex]
		if !tolerance.IsWithin(currentCoefficient, 0, DefaultTolerance) {
			for j := i + 1; j < len(op); j++ {
				nextEquation := op[j]
				// No need to capture the error, the only possible error is mismatched or out-of-band terms
				// This is tested for in AllEquationsHaveSameNumberOfTerms above.
				op[j], _ = currentEquation.CancelTerm(nextEquation, termIndex)
			}
		}
		termIndex++
	}

	return op, nil
}

// IsTriangularForm determines whether the system is in triangular form, where the top row starts with a non-zero
// term, the next one down starts with a zero etc., the one after that starts with two zero terms etc.
func (s1 System) IsTriangularForm() (bool, error) {
	if !s1.AllEquationsHaveSameNumberOfTerms() {
		return false, errors.New("all equations in a system need to have the same number of terms")
	}
	for i, e := range s1 {
		// Check that everything leading up to the current term is zero.
		for j := 0; j < i; j++ {
			if !tolerance.IsWithin(e.NormalVector[j], 0, DefaultTolerance) {
				return false, nil
			}
		}
	}
	return true, nil
}

// AllEquationsHaveSameNumberOfTerms returns true when all equations in the system have the same number of terms.
func (s1 System) AllEquationsHaveSameNumberOfTerms() bool {
	var length int
	for i, e := range s1 {
		if i == 0 {
			length = len(e.NormalVector)
			continue
		}
		if length != len(e.NormalVector) {
			return false
		}
	}
	return true
}

// ComputeRREF computes the Reduced Row Echelon Form of the system. ok returns
// whether all of the terms in the equation have got a value.
func (s1 System) ComputeRREF() (s System, ok bool, err error) {
	s, err = s1.TriangularForm()
	if err != nil {
		return s, false, err
	}

	var termIndexIsNonZero []bool
	if len(s1) > 0 {
		termIndexIsNonZero = make([]bool, len(s1[0].NormalVector))
	}

	// Iterate from bottom to top.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i].NormalVector.IsZeroVector() {
			// Nothing to solve, skip this line.
			continue
		}

		// Make the leading term have a coefficient of one.
		nonZeroTermIndex, v, _ := s[i].FirstNonZeroCoefficient()
		termIndexIsNonZero[nonZeroTermIndex] = true
		coefficient := float64(1.0) / v
		s[i] = s[i].Scale(coefficient)

		// Cancel this term in the equations above this one.
		for j := i - 1; j >= 0; j-- {
			// No need to catch the error, we've already checked that the s[i] term is nonzero and that the
			// equations have the same number of terms.
			s[j], _ = s[i].CancelTerm(s[j], nonZeroTermIndex)
		}
	}
	ok = allTrue(termIndexIsNonZero)
	return s, ok, nil
}

func allTrue(bools []bool) bool {
	for _, v := range bools {
		if !v {
			return false
		}
	}
	return true
}

// IsRREF determines whether a system is in Reduced Row Echelon form.
func (s1 System) IsRREF() (bool, error) {
	isTriangular, err := s1.IsTriangularForm()
	if !isTriangular || err != nil {
		return isTriangular, err
	}

	var rowsWithNonZeroTerms []bool
	if len(s1) > 0 {
		rowsWithNonZeroTerms = make([]bool, len(s1[0].NormalVector))
	}

	for _, e := range s1 {
		// Check that the value of the first non-zero term is one, then everything after is zero.
		hadNonZeroTerm := false
		for termIndex, v := range e.NormalVector {
			if !tolerance.IsWithin(v, 0, DefaultTolerance) {
				if hadNonZeroTerm {
					return false, nil
				}
				if !tolerance.IsWithin(v, 1, DefaultTolerance) {
					// It's not RREF if there's a term that isn't one or zero.
					return false, nil
				}
				hadNonZeroTerm = true
				rowsWithNonZeroTerms[termIndex] = true
			}
		}
	}
	return allTrue(rowsWithNonZeroTerms), nil
}

// Solve solves the equation using Gaussian Elimination and returns whether the solution has
// a single solution, no solutions, infinite solutions or can't be calculated due to an error.
func (s1 System) Solve() (solution Vector, noSolution bool, infiniteSolutions bool, err error) {
	s, allVariablesSet, err := s1.ComputeRREF()
	if err != nil {
		return solution, true, false, err
	}

	// Check whether we're in a 0=1 situation.
	for _, equation := range s {
		// First the first non-zero coefficient.
		_, _, ok := equation.FirstNonZeroCoefficient()
		// If we don't have a non-zero coefficient, and the constant term is not zero.
		// Then we have a situation where 0 is not equal to zero, i.e. the equation is
		// inconsistent.
		if !ok && !tolerance.IsWithin(equation.ConstantTerm, 0, DefaultTolerance) {
			// Return that we have no solution.
			return solution, true, false, nil
		}
	}

	if !allVariablesSet {
		// We have a free variable, so we can generate infinite
		// solutions by modifying the free variable(s).
		return solution, false, true, nil
	}

	// We must have a single intersection.
	var solutionVector []float64
	if len(s) > 0 {
		solutionVector = make([]float64, len(s[0].NormalVector))
	}
	for i, equation := range s {
		solutionVector[i] = equation.ConstantTerm
	}
	return Vector(solutionVector), false, false, nil
}
