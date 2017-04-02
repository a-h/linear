package linear

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/a-h/linear/tolerance"
)

// Line consists of a normal vector which specifies the direction, and a constant term.
// The basepoint is calculated as required.
type Line struct {
	NormalVector Vector
	ConstantTerm float64
}

// NewLine creates a new Line based on the variable terms, e.g.:
// Ax + By = C
// For example, for a 2D vector (2, 3) and constant 5 would result in 2x + 3y = 5
// The slope intercept (y = mx + b) form of that would be:
//  2x + 3y = 5
//       3y = 5 - 2x
//        y = 1/3(5-2x)
// To get back to standard form:
// To find the y intercept (a point on the line), set x to zero.
//  2x + 3y = 5
//       3y = 5 - 0
//        y = 5/3
//  y intercept = (0, 5/3)
// To find the x intercept, set y to zero.
//  2x + 3y = 5
//  2x - 0 = 5
//  2x = 5
//  x = 5/2
//   x intercept = (5/2, 0)
func NewLine(normalVector Vector, constantTerm float64) Line {
	return Line{
		ConstantTerm: constantTerm,
		NormalVector: normalVector,
	}
}

// NonZeroValuePoint finds a point on the Line where one of the dimension values is not zero.
// If a non-zero coefficient is not found, ok is set to false.
func (l1 Line) NonZeroValuePoint() (nonzero Vector, ok bool) {
	basepointVector := make([]float64, len(l1.NormalVector))
	index, value, ok := firstNonZeroElement(l1.NormalVector)
	if !ok {
		return Vector{}, false
	}
	basepointVector[index] = l1.ConstantTerm / value
	return Vector(basepointVector), true
}

// FirstNonZeroCoefficient finds the first non-zero coefficient of the normal vector of the plane.
// If a non-zero coefficient is not found, ok is set to false.
func (l1 Line) FirstNonZeroCoefficient() (index int, value float64, ok bool) {
	return firstNonZeroElement(l1.NormalVector)
}

func firstNonZeroElement(v Vector) (index int, value float64, ok bool) {
	for i, value := range v {
		if !tolerance.IsWithin(value, 0, DefaultTolerance) {
			return i, value, true
		}
	}

	return 0, 0, false
}

func (l1 Line) String() string {
	buf := bytes.Buffer{}
	for i, p := range l1.NormalVector {
		if i == 0 {
			// The first element should have an integrated +/- sign.
			buf.WriteString(fmt.Sprintf("%v", p))
			// Write out the x_1 specifier.
			buf.WriteString(fmt.Sprintf("x%s", getSubscript(i+1)))
			continue
		}

		// For anything after index zero, the sign becomes the operator.
		buf.WriteString(operator(p))
		buf.WriteString(fmt.Sprintf("%v", math.Abs(p)))
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
func (l1 Line) IsParallelTo(l2 Line) (bool, error) {
	return l1.NormalVector.IsParallelTo(l2.NormalVector)
}

// Eq determines if two lines are equal.
func (l1 Line) Eq(l2 Line) (bool, error) {
	// If either vector is zero, and the other isn't they're not equal.
	l1IsZero := l1.NormalVector.IsZeroVector()
	l2IsZero := l2.NormalVector.IsZeroVector()

	if l1IsZero || l2IsZero {
		if l1IsZero && !l2IsZero || !l1IsZero && l2IsZero {
			return false, nil
		}
		// Check the constant terms are the same if both are zero vectors.
		return tolerance.IsWithin(l1.ConstantTerm, l2.ConstantTerm, DefaultTolerance), nil
	}

	// If they're not parallel, there's no way they're going to be equal.
	isParallel, err := l1.IsParallelTo(l2)
	if !isParallel || err != nil {
		return false, err
	}

	// Subtract a point on l2 from l1, which creates a vector between the two points.
	// The vector that joins the lines should be orthogonal to the normal vector, or it's not equal.
	// No need to capture the sub error here, because the error would be because the number of terms in the vector
	// is different, which is already captured by the parallel check.
	// No need to capture the ok coming back from l1 / l2's NonZeroValuePoint() because there's already a check
	// to see if they're zero vectors above.
	l1NonZeroPoint, _ := l1.NonZeroValuePoint()
	l2NonZeroPoint, _ := l2.NonZeroValuePoint()
	connectingVector, _ := l1NonZeroPoint.Sub(l2NonZeroPoint)

	// No need to check orthogonality of both vectors, because they're parallel to each other.
	return connectingVector.IsOrthogonalTo(l1.NormalVector)
}

// Y gets the Y value for a given X.
func (l1 Line) Y(x float64) (float64, error) {
	if len(l1.NormalVector) != 2 {
		return 0, errors.New("The Y function only supports lines with 2 dimensions.")
	}
	// ax + by = c
	// by = c - ax
	// y = (c - ax) / b
	return (l1.ConstantTerm - (l1.NormalVector[0] * x)) / l1.NormalVector[1], nil
}

// X gets the X value for a given Y value.
func (l1 Line) X(y float64) (float64, error) {
	if len(l1.NormalVector) != 2 {
		return 0, errors.New("The X function only supports lines with 2 dimensions.")
	}
	// ax + by = c
	// ax = c - by
	// x = (c - by) / a
	return (l1.ConstantTerm - (l1.NormalVector[1] * y)) / l1.NormalVector[0], nil
}

// IntersectionWith calculates the intersection with another 2D line.
// intersects is set to true if the lines intersect.
// equal is set to true if the lines are equal and therefore intersect infinitely many times.
func (l1 Line) IntersectionWith(l2 Line) (intersection Vector, intersects bool, equal bool, err error) {
	if len(l1.NormalVector) != 2 || len(l2.NormalVector) != 2 {
		return Vector{}, false, false, fmt.Errorf("The IntersectionWith function requires that both lines must have 2 dimensions. The base line has %d dimensions, l2 has %d dimensions.", len(l1.NormalVector), len(l2.NormalVector))
	}

	// Handle zero vector edge case.
	if l1.NormalVector.IsZeroVector() || l2.NormalVector.IsZeroVector() {
		return Vector{}, false, false, nil
	}

	// If the lines are equal, there are infinitely many intersections.
	// No need to catch the error, because we've already checked that the vectors have equal lengths.
	if eq, _ := l1.Eq(l2); eq {
		l1NonZeroValuePoint, _ := l1.NonZeroValuePoint()
		return l1NonZeroValuePoint, true, true, nil
	}

	// If the lines are parallel but not equal, there will never be an intersection unless the lines are equal.
	// No need to catch the error, the Eq test above has already done the same.
	isParallel, _ := l1.IsParallelTo(l2)
	if isParallel {
		return Vector{}, false, false, nil
	}

	// Explanation at http://math.stackexchange.com/questions/48395/how-to-find-the-point-of-intersection-of-two-lines
	// And https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection#Given_the_equations_of_the_lines
	// At the point where the two lines intersect (if they do), both y coordinates will be the same, hence the following equality:
	// y=ax+c and y=bx+d
	// i.e. ax+c=bx+d
	a, b, c := l1.NormalVector[0], l1.NormalVector[1], l1.ConstantTerm
	d, e, f := l2.NormalVector[0], l2.NormalVector[1], l2.ConstantTerm

	// Given that x and y have the same value in each equation.
	// ax + by = c
	// dx + ey = f
	// Find the definition of y
	// by = c - ax
	// y = (c - ax)/b (use later to calculate the y value once we have the value of x)
	// Insert the reworked equation in to replace y in the 2nd equation and get the value of x
	// dx + ey = f
	// dx + e((c - ax)/b) = f
	// dx + ec/b - eax/b = f
	// dx - eax/b = f - ec/b
	// bdx - eax = bf - ec
	// x(bd - ea) = bf - ec
	// x = bf - ec / bd - ea
	x := ((b * f) - (e * c)) / ((b * d) - (e * a))
	y := (c - (a * x)) / b
	return NewVector(x, y), true, false, nil
}

// CancelTerm cancels a term in the target line by determining the coefficient which links them
// and applying the first term to the second term to cancel them out.
func (l1 Line) CancelTerm(target Line, termIndex int) (Line, error) {
	if termIndex >= len(l1.NormalVector) || termIndex < 0 {
		return Line{}, fmt.Errorf("term index %d is not present in l1", termIndex)
	}

	if termIndex >= len(target.NormalVector) || termIndex < 0 {
		return Line{}, fmt.Errorf("term index %d is not present in the target line", termIndex)
	}

	srcCoefficient := l1.NormalVector[termIndex]
	dstCoefficient := target.NormalVector[termIndex]

	if tolerance.IsWithin(srcCoefficient, 0, DefaultTolerance) {
		return target, fmt.Errorf("the source line %v has a zero coefficient for term index %d, so can't be used to clear that term from %v", l1, termIndex, target)
	}

	factor := dstCoefficient / -srcCoefficient

	// Multiply by the difference between them.
	multipliedVector := l1.NormalVector.Scale(factor)
	multipliedConstant := l1.ConstantTerm * factor

	outputVector, err := target.NormalVector.Add(multipliedVector)
	if err != nil {
		return Line{}, err
	}
	outputConstant := target.ConstantTerm + multipliedConstant
	return NewLine(outputVector, outputConstant), nil
}

// Scale scales the line by a scalar multiplier.
func (l1 Line) Scale(scalar float64) Line {
	return NewLine(l1.NormalVector.Scale(scalar), l1.ConstantTerm*scalar)
}

// PivotIndex returns the index where the first non-zero term is one, then everything after it is zero.
func (l1 Line) PivotIndex() (index int, hasPivot bool) {
	alreadyHadNonZeroTerm := false
	for termIndex, v := range l1.NormalVector {
		if !tolerance.IsWithin(v, 0, DefaultTolerance) {
			if alreadyHadNonZeroTerm {
				// Can't have a pivot if there are two non-zero coefficients.
				return -1, false
			}
			if !tolerance.IsWithin(v, 1, DefaultTolerance) {
				// It doesn't have a pivot, if there's a term that isn't one or zero.
				return -1, false
			}
			alreadyHadNonZeroTerm = true
			index = termIndex
		}
	}
	if !alreadyHadNonZeroTerm {
		return -1, false
	}
	return index, true
}
