package linear

import (
	"math"
	"testing"

	"strings"

	"github.com/a-h/linear/tolerance"
)

func TestCreatingVectorsUsingVariadicInput(t *testing.T) {
	v := NewVector(1, 2)
	if v[0] != 1 {
		t.Errorf("For index zero, expected 1, but got %d", v[0])
	}
	if v[1] != 2 {
		t.Errorf("For index one, expected 2, but got %d", v[1])
	}
}

func TestCreatingVectorsFromAnArray(t *testing.T) {
	array := []float64{1, 2}
	v := NewVector(array...)
	if v[0] != 1 {
		t.Errorf("For index zero, expected 1, but got %d", v[0])
	}
	if v[1] != 2 {
		t.Errorf("For index one, expected 2, but got %d", v[1])
	}
}

func TestCreatingVectorsByConversion(t *testing.T) {
	array := []float64{1, 2}
	v := Vector(array)
	if v[0] != 1 {
		t.Errorf("For index zero, expected 1, but got %d", v[0])
	}
	if v[1] != 2 {
		t.Errorf("For index one, expected 2, but got %d", v[1])
	}
}

func TestVectorStringRepresentations(t *testing.T) {
	tests := []struct {
		name     string
		in       []float64
		expected string
	}{
		{
			name:     "Zero dimensions",
			in:       []float64{},
			expected: "[]",
		},
		{
			name:     "Single dimension",
			in:       []float64{12.5},
			expected: "[12.5]",
		},
		{
			name:     "Two dimensions",
			in:       []float64{1.45, 2},
			expected: "[1.45, 2]",
		},
		{
			name:     "Three dimensions",
			in:       []float64{99, 12.1, 3.15666},
			expected: "[99, 12.1, 3.15666]",
		},
	}

	for _, test := range tests {
		v := NewVector(test.in...)
		actual := v.String()

		if actual != test.expected {
			t.Errorf("%s: Expected '%s', but got '%s'", test.name, test.expected, actual)
		}
	}
}

func TestVectorEquality(t *testing.T) {
	tests := []struct {
		name     string
		a        Vector
		b        Vector
		expected bool
	}{
		{
			name:     "Different dimensions (1:2)",
			a:        NewVector(1),
			b:        NewVector(1, 1),
			expected: false,
		},
		{
			name:     "Different dimensions (2:1)",
			a:        NewVector(1, 1),
			b:        NewVector(1),
			expected: false,
		},
		{
			name:     "Single dimension, different values",
			a:        NewVector(1),
			b:        NewVector(2),
			expected: false,
		},
		{
			name:     "Single dimension, same values",
			a:        NewVector(1),
			b:        NewVector(1),
			expected: true,
		},
		{
			name:     "Multiple dimensions, same values",
			a:        NewVector(1, 2, 3, 4, 5),
			b:        NewVector(1, 2, 3, 4, 5),
			expected: true,
		},
		{
			name:     "Multiple dimensions, different values",
			a:        NewVector(1, 2, 3, 4, 5),
			b:        NewVector(1, 2, 3, 4, 4),
			expected: false,
		},
	}

	for _, test := range tests {
		actual := test.a.Eq(test.b)

		if actual != test.expected {
			t.Errorf("%s: For %v and %v, expected '%v', but got '%v'", test.name, test.a, test.b, test.expected, actual)
		}
	}
}

func TestVectorEqualityWithinTolerance(t *testing.T) {
	tests := []struct {
		name      string
		a         Vector
		b         Vector
		tolerance float64
		expected  bool
	}{
		{
			name:      "Different dimensions (1:2)",
			a:         NewVector(1),
			b:         NewVector(1, 1),
			tolerance: float64(0),
			expected:  false,
		},
		{
			name:      "Different dimensions (2:1)",
			a:         NewVector(1, 1),
			b:         NewVector(1),
			tolerance: float64(0),
			expected:  false,
		},
		{
			name:      "Single dimension, different values",
			a:         NewVector(1),
			b:         NewVector(2),
			tolerance: float64(0),
			expected:  false,
		},
		{
			name:      "Single dimension, same values",
			a:         NewVector(1),
			b:         NewVector(1),
			tolerance: float64(0),
			expected:  true,
		},
		{
			name:      "Multiple dimensions, same values",
			a:         NewVector(1, 2, 3, 4, 5),
			b:         NewVector(1, 2, 3, 4, 5),
			tolerance: float64(0),
			expected:  true,
		},
		{
			name:      "Multiple dimensions, different values",
			a:         NewVector(1, 2, 3, 4, 5),
			b:         NewVector(1, 2, 3, 4, 4),
			tolerance: float64(0),
			expected:  false,
		},
		{
			name:      "One decimal place",
			a:         NewVector(1, 1),
			b:         NewVector(1.1, 1.1),
			tolerance: tolerance.OneDecimalPlace,
			expected:  true,
		},
		{
			name:      "Two decimal places",
			a:         NewVector(1, 1),
			b:         NewVector(1.1, 1.1),
			tolerance: tolerance.TwoDecimalPlaces,
			expected:  false,
		},
		{
			name:      "Three decimal places",
			a:         NewVector(1, 1),
			b:         NewVector(1.001, 1.001),
			tolerance: tolerance.ThreeDecimalPlaces,
			expected:  true,
		},
	}

	for _, test := range tests {
		actual := test.a.EqWithinTolerance(test.b, test.tolerance)

		if actual != test.expected {
			t.Errorf("%s: %v and %v within tolerance %f, expected '%v', but got '%v'", test.name, test.a, test.b, test.tolerance, test.expected, actual)
		}
	}
}

func TestVectorAddition(t *testing.T) {
	tests := []struct {
		name                 string
		a                    Vector
		b                    Vector
		expected             Vector
		expectedErrorMessage string
	}{
		{
			name:     "Positive",
			a:        NewVector(1, 1),
			b:        NewVector(1, 1),
			expected: NewVector(2, 2),
		},
		{
			name:     "Zeroes",
			a:        NewVector(0, 1),
			b:        NewVector(0, 3),
			expected: NewVector(0, 4),
		},
		{
			name:     "Negatives",
			a:        NewVector(-4, -6),
			b:        NewVector(4, 8),
			expected: NewVector(0, 2),
		},
		{
			name:     "Three dimensions",
			a:        NewVector(1, 2, 3),
			b:        NewVector(3, 2, 1),
			expected: NewVector(4, 4, 4),
		},
		{
			name:                 "Mismatched dimensions",
			a:                    NewVector(1),
			b:                    NewVector(1, 2),
			expected:             Vector{},
			expectedErrorMessage: "cannot add vectors together because they have different dimensions (1 and 2)",
		},
	}

	for _, test := range tests {
		actual, err := test.a.Add(test.b)

		if !actual.Eq(test.expected) {
			t.Errorf("%s: For %v + %v, expected '%v', but got '%v'", test.name, test.a, test.b, test.expected, actual)
		}

		if err != nil {
			if test.expectedErrorMessage == "" {
				t.Errorf("%s: For %v + %v, no error was expected, but got '%v'", test.name, test.a, test.b, err)
				continue
			}

			if test.expectedErrorMessage != err.Error() {
				t.Errorf("%s: For %v + %v, expected error message '%v', but got '%v'", test.name, test.a, test.b, test.expectedErrorMessage, err)
			}
		}
	}
}

func TestVectorSubtraction(t *testing.T) {
	tests := []struct {
		name                 string
		a                    Vector
		b                    Vector
		expected             Vector
		expectedErrorMessage string
	}{
		{
			name:     "Positive",
			a:        NewVector(1, 1),
			b:        NewVector(1, 1),
			expected: NewVector(0, 0),
		},
		{
			name:     "Zeroes",
			a:        NewVector(0, 1),
			b:        NewVector(0, 3),
			expected: NewVector(0, -2),
		},
		{
			name:     "Negatives",
			a:        NewVector(-4, -6),
			b:        NewVector(-4, -6),
			expected: NewVector(0, 0),
		},
		{
			name:     "Three dimensions",
			a:        NewVector(-1, -2, -3),
			b:        NewVector(1, 2, 3),
			expected: NewVector(-2, -4, -6),
		},
		{
			name:                 "Mismatched dimensions",
			a:                    NewVector(1),
			b:                    NewVector(1, 2),
			expected:             Vector{},
			expectedErrorMessage: "cannot subtract vectors because they have different dimensions (1 and 2)",
		},
	}

	for _, test := range tests {
		actual, err := test.a.Sub(test.b)

		if !actual.Eq(test.expected) {
			t.Errorf("%s: For %v - %v, expected '%v', but got '%v'", test.name, test.a, test.b, test.expected, actual)
		}

		if err != nil {
			if test.expectedErrorMessage == "" {
				t.Errorf("%s: For %v - %v, no error was expected, but got '%v'", test.name, test.a, test.b, err)
				continue
			}

			if test.expectedErrorMessage != err.Error() {
				t.Errorf("%s: For %v - %v, expected error message '%v', but got '%v'", test.name, test.a, test.b, test.expectedErrorMessage, err)
			}
		}
	}
}

func TestVectorMultiplication(t *testing.T) {
	tests := []struct {
		name                 string
		a                    Vector
		b                    Vector
		expected             Vector
		expectedErrorMessage string
	}{
		{
			name:     "Positive",
			a:        NewVector(3, 4),
			b:        NewVector(3, 4),
			expected: NewVector(9, 16),
		},
		{
			name:     "Zeroes",
			a:        NewVector(0, 1),
			b:        NewVector(6, 3),
			expected: NewVector(0, 3),
		},
		{
			name:     "Negatives",
			a:        NewVector(-4, -6),
			b:        NewVector(-4, 6),
			expected: NewVector(16, -36),
		},
		{
			name:     "Three dimensions",
			a:        NewVector(-1, -2, 3),
			b:        NewVector(1, 2, -3),
			expected: NewVector(-1, -4, -9),
		},
		{
			name:                 "Mismatched dimensions",
			a:                    NewVector(1),
			b:                    NewVector(1, 2),
			expected:             Vector{},
			expectedErrorMessage: "cannot multiply vectors because they have different dimensions (1 and 2)",
		},
	}

	for _, test := range tests {
		actual, err := test.a.Mul(test.b)

		if !actual.Eq(test.expected) {
			t.Errorf("%s: For %v * %v, expected '%v', but got '%v'", test.name, test.a, test.b, test.expected, actual)
		}

		if err != nil {
			if test.expectedErrorMessage == "" {
				t.Errorf("%s: For %v * %v, no error was expected, but got '%v'", test.name, test.a, test.b, err)
				continue
			}

			if test.expectedErrorMessage != err.Error() {
				t.Errorf("%s: For %v * %v, expected error message '%v', but got '%v'", test.name, test.a, test.b, test.expectedErrorMessage, err)
			}
		}
	}
}

func TestVectorScalarMultiplication(t *testing.T) {
	tests := []struct {
		name     string
		a        Vector
		scalar   float64
		expected Vector
	}{
		{
			name:     "Positive",
			a:        NewVector(3, 4),
			scalar:   3,
			expected: NewVector(9, 12),
		},
		{
			name:     "Zeroes",
			a:        NewVector(0, 1),
			scalar:   12,
			expected: NewVector(0, 12),
		},
		{
			name:     "Negatives",
			a:        NewVector(-4, 6),
			scalar:   -4,
			expected: NewVector(16, -24),
		},
		{
			name:     "Three dimensions",
			a:        NewVector(1, 2, 3),
			scalar:   0.5,
			expected: NewVector(0.5, 1, 1.5),
		},
	}

	for _, test := range tests {
		actual := test.a.Scale(test.scalar)

		if !actual.Eq(test.expected) {
			t.Errorf("%s: For %v * %f, expected '%v', but got '%v'", test.name, test.a, test.scalar, test.expected, actual)
		}
	}
}

func TestVectorMagnitudeCalculation(t *testing.T) {
	tests := []struct {
		name     string
		input    Vector
		expected float64
	}{
		{
			name:     "Pythagoran triangle",
			input:    NewVector(4, 3),
			expected: 5,
		},
		{
			name:     "Ones",
			input:    NewVector(1, 1, 1),
			expected: math.Sqrt(1 + 1 + 1),
		},
		{
			name:     "Zeroes",
			input:    NewVector(0, 0, 0, 0),
			expected: 0,
		},
		{
			name:     "Negative numbers",
			input:    NewVector(-4, 3),
			expected: 5,
		},
		{
			name:     "Negative one",
			input:    NewVector(-1, 1, 1),
			expected: math.Sqrt(3),
		},
	}

	for _, test := range tests {
		actual := test.input.Magnitude()

		if actual != test.expected {
			t.Errorf("%s: For the magnitude of %v, expected '%f', but got '%f'", test.name, test.input, test.expected, actual)
		}
	}
}

func TestIsZeroVectorFunction(t *testing.T) {
	tests := []struct {
		name     string
		input    Vector
		expected bool
	}{
		{
			name:     "Empty",
			input:    NewVector(),
			expected: true,
		},
		{
			name:     "Zeroes",
			input:    NewVector(0, 0),
			expected: true,
		},
		{
			name:     "Negatives",
			input:    NewVector(-5.581, -2.136),
			expected: false,
		},
		{
			name:     "Mixed zero",
			input:    NewVector(0, 1),
			expected: false,
		},
		{
			name:     "Close to zero",
			input:    NewVector(0.00000000001, 0.0),
			expected: true,
		},
	}

	for _, test := range tests {
		actual := test.input.IsZeroVector()

		if actual != test.expected {
			t.Errorf("%s: Expected calculating whether %v is a zero vector to return %v, but got %v", test.name, test.input, test.expected, actual)
		}
	}
}

func TestVectorNormalization(t *testing.T) {
	tests := []struct {
		name     string
		input    Vector
		expected Vector
	}{
		{
			name:     "Pythagoran triple",
			input:    NewVector(4, 3),
			expected: NewVector(0.8, 0.6), // 4*(1/5) = 4 * 0.2 = 0.8, 3*(1/5) = 3 * 0.2 = 0.6
		},
		{
			name:     "Zeroes",
			input:    NewVector(0, 0),
			expected: NewVector(0, 0),
		},
		{
			name:     "Udacity example 1",
			input:    NewVector(5.581, -2.136),
			expected: NewVector(0.934, -0.357),
		},
		{
			name:     "Udacity example 2",
			input:    NewVector(1.996, 3.108, -4.554),
			expected: NewVector(0.340, 0.530, -0.777),
		},
	}

	for _, test := range tests {
		actual := test.input.Normalize()

		if !actual.EqWithinTolerance(test.expected, tolerance.ThreeDecimalPlaces) {
			t.Errorf("%s: For the direction of %v, expected %v, but got %v", test.name, test.input, test.expected, actual)
		}

		if !actual.IsZeroVector() {
			if !tolerance.IsWithin(actual.Magnitude(), 1, tolerance.ThreeDecimalPlaces) {
				t.Errorf("%s: The magnitude for a unit vector should always be one, but got %f", test.name, actual.Magnitude())
			}
		}
	}
}

func TestDotProduct(t *testing.T) {
	tests := []struct {
		name                 string
		a                    Vector
		b                    Vector
		expected             float64
		expectedErrorMessage string
	}{
		{
			name:     "Positive",
			a:        NewVector(3, 4),
			b:        NewVector(3, 4),
			expected: (3 * 3) + (4 * 4),
		},
		{
			name:     "Zeroes",
			a:        NewVector(0, 3),
			b:        NewVector(3, 0),
			expected: 0,
		},
		{
			name:     "Negatives",
			a:        NewVector(-4, 4),
			b:        NewVector(4, 6),
			expected: (-4 * 4) + (4 * 6),
		},
		{
			name:     "Three dimensions",
			a:        NewVector(-1, -2, 3),
			b:        NewVector(1, 2, -3),
			expected: (-1 * 1) + (-2 * 2) + (3 * -3),
		},
		{
			name:                 "Mismatched dimensions",
			a:                    NewVector(1),
			b:                    NewVector(1, 2),
			expected:             0,
			expectedErrorMessage: "cannot calculate the dot product of the vectors because they have different dimensions (1 and 2)",
		},
	}

	for _, test := range tests {
		actual, err := test.a.DotProduct(test.b)

		if actual != test.expected {
			t.Errorf("%s: For the dot product of %v and %v, expected '%v', but got '%v'", test.name, test.a, test.b, test.expected, actual)
		}

		if err != nil {
			if test.expectedErrorMessage == "" {
				t.Errorf("%s: For the dot product of %v and %v, no error was expected, but got '%v'", test.name, test.a, test.b, err)
				continue
			}

			if test.expectedErrorMessage != err.Error() {
				t.Errorf("%s: For the dot proudct of %v and %v, expected error message '%v', but got '%v'", test.name, test.a, test.b, test.expectedErrorMessage, err)
			}
		}
	}
}

func TestAngleBetweenFunction(t *testing.T) {
	tests := []struct {
		name                 string
		a                    Vector
		b                    Vector
		expected             Radian
		expectedErrorMessage string
	}{
		{
			name:     "Right angled triangle (3, 4, 5)",
			a:        NewVector(3, 0),
			b:        NewVector(0, 4),
			expected: 1.5708, // 90 degrees
		},
		{
			name:     "Triangle (8, 6, 10)",
			a:        NewVector(8, 6),
			b:        NewVector(8, 0),
			expected: 0.6440265, // 36.9 degrees
		},
		{
			name:     "Triangle (8, 6, 10)",
			a:        NewVector(-8, -6),
			b:        NewVector(0, -6),
			expected: 0.9267698, // 53.1 degrees
		},
		{
			name:     "Equilateral triangle in 3 dimensions (1)",
			a:        NewVector(1, 0, 1),
			b:        NewVector(0, 1, 1),
			expected: 1.0472, // 60 degrees
		},
		{
			name:     "Equilateral triangle in 3 dimensions (2)",
			a:        NewVector(0, 1, 1),
			b:        NewVector(1, 1, 0),
			expected: 1.0472, // 60 degrees
		},
		{
			name:                 "Mismatched dimensions",
			a:                    NewVector(1),
			b:                    NewVector(1, 2),
			expected:             0,
			expectedErrorMessage: "cannot calculate the dot product of the vectors because they have different dimensions (1 and 2)",
		},
	}

	for _, test := range tests {
		actual, err := test.a.AngleBetween(test.b)

		if !tolerance.IsWithin(float64(actual), float64(test.expected), tolerance.ThreeDecimalPlaces) {
			t.Errorf("%s: For the angle between %v and %v, expected %f radians, but got %f", test.name, test.a, test.b, test.expected, actual)
		}

		if err != nil {
			if test.expectedErrorMessage == "" {
				t.Errorf("%s: For the angle between %v and %v, no error was expected, but got '%v'", test.name, test.a, test.b, err)
				continue
			}

			if test.expectedErrorMessage != err.Error() {
				t.Errorf("%s: For the angle between %v and %v, expected error message '%v', but got '%v'", test.name, test.a, test.b, test.expectedErrorMessage, err)
			}
		}
	}
}

func TestIsParallelToFunction(t *testing.T) {
	tests := []struct {
		name                 string
		a                    Vector
		b                    Vector
		expected             bool
		expectedErrorMessage string
	}{
		{
			name:     "Exact",
			a:        NewVector(0, 3),
			b:        NewVector(0, 3),
			expected: true,
		},
		{
			name:     "Twice the size",
			a:        NewVector(1, 1),
			b:        NewVector(2, 2),
			expected: true,
		},
		{
			name:     "Triple the size",
			a:        NewVector(2, 3, 1),
			b:        NewVector(6, 9, 3),
			expected: true,
		},
		{
			name:     "Opposite direction",
			a:        NewVector(-2, -2),
			b:        NewVector(2, 2),
			expected: true,
		},
		{
			name:                 "Different sizes",
			a:                    NewVector(-2, -2, -1),
			b:                    NewVector(2, 2),
			expected:             false,
			expectedErrorMessage: "cannot calculate whether the vectors are parallel because they have different dimensions (3 and 2)",
		},
		{
			name:     "Something to nothing",
			a:        NewVector(0, 3),
			b:        NewVector(0, 0),
			expected: true,
		},
		{
			name:     "Nothing to something",
			a:        NewVector(0, 0, 0),
			b:        NewVector(0, 0, 5),
			expected: true,
		},
		{
			name:     "Nothing to nothing",
			a:        NewVector(0, 0, 0, 0),
			b:        NewVector(0, 0, 0, 0),
			expected: true,
		},
	}

	for _, test := range tests {
		actual, err := test.a.IsParallelTo(test.b)

		if actual != test.expected {
			t.Errorf("%s: For %v and %v - expected %v, but got %v", test.name, test.a, test.b, test.expected, actual)
		}

		if err != nil {
			if test.expectedErrorMessage == "" {
				t.Errorf("%s: For %v and %v, no error was expected, but got '%v'", test.name, test.a, test.b, err)
				continue
			}

			if test.expectedErrorMessage != err.Error() {
				t.Errorf("%s: For %v and %v, expected error message '%v', but got '%v'", test.name, test.a, test.b, test.expectedErrorMessage, err)
			}
		}
	}
}

func TestIsOrthogonalToFunction(t *testing.T) {
	tests := []struct {
		name                 string
		a                    Vector
		b                    Vector
		expected             bool
		expectedErrorMessage string
	}{
		{
			name:     "Equal",
			a:        NewVector(1, 1),
			b:        NewVector(1, 1),
			expected: false,
		},
		{
			name:     "Parallel",
			a:        NewVector(1, 1),
			b:        NewVector(2, 2),
			expected: false,
		},
		{
			name:     "Right angle",
			a:        NewVector(5, 0),
			b:        NewVector(0, 5),
			expected: true,
		},
		{
			name:     "Three dimsional angle",
			a:        NewVector(5, 0, 5),
			b:        NewVector(0, 5, 0),
			expected: true,
		},
		{
			name:                 "Different sizes",
			a:                    NewVector(-2, -2, -1),
			b:                    NewVector(2, 2),
			expected:             false,
			expectedErrorMessage: "error calculating whether the vectors are orthogonal: cannot calculate the dot product of the vectors because they have different dimensions (3 and 2)",
		},
	}

	for _, test := range tests {
		actual, err := test.a.IsOrthogonalTo(test.b)

		if actual != test.expected {
			t.Errorf("%s: For %v and %v - expected %v, but got %v", test.name, test.a, test.b, test.expected, actual)
		}

		if err != nil {
			if test.expectedErrorMessage == "" {
				t.Errorf("%s: For %v and %v, no error was expected, but got '%v'", test.name, test.a, test.b, err)
				continue
			}

			if test.expectedErrorMessage != err.Error() {
				t.Errorf("%s: For %v and %v, expected error message '%v', but got '%v'", test.name, test.a, test.b, test.expectedErrorMessage, err)
			}
		}
	}
}

func TestProjectionFunction(t *testing.T) {
	tests := []struct {
		name                 string
		basis                Vector
		p                    Vector
		expected             Vector
		expectedErrorMessage string
	}{
		{
			name:     "Pythagoran triangle",
			basis:    NewVector(10, 0), // 10 across, but no height, just a straight line
			p:        NewVector(4, 3),  // 4 along and 3 up, forming the hypotenuse of the triangle (5 long).
			expected: NewVector(4, 0),  // 4 along and 0 up, the base of the triangle, with the 3 height made up by the orthogonal vector.
		},
		// Example from https://www.khanacademy.org/math/linear-algebra/matrix-transformations/lin-trans-examples/v/introduction-to-projections
		{
			name:     "Khan Academy #1",
			basis:    NewVector(2, 1),
			p:        NewVector(2, 3),
			expected: NewVector(2.8, 1.4),
		},
		{
			name:                 "Different sizes",
			basis:                NewVector(-2, -2, -1),
			p:                    NewVector(2, 2),
			expected:             Vector{},
			expectedErrorMessage: "error projecting [2, 2] onto [-2, -2, -1] with error: cannot calculate the dot product of the vectors because they have different dimensions (2 and 3)",
		},
	}

	for _, test := range tests {
		actual, err := test.basis.Projection(test.p)

		if !actual.Eq(test.expected) {
			t.Errorf("%s: Projecting %v onto basis %v - expected %v, but got %v", test.name, test.p, test.basis, test.expected, actual)
		}

		if err != nil {
			if test.expectedErrorMessage == "" {
				t.Errorf("%s: Projecting %v onto basis %v, no error was expected, but got '%v'", test.name, test.p, test.basis, err)
				continue
			}

			if !strings.HasSuffix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("%s: Projecting %v onto %v, expected error message to start with '%v', but got '%v'", test.name, test.p, test.basis, test.expectedErrorMessage, err)
			}
		}
	}
}

func TestProjectOrthogonalFunction(t *testing.T) {
	tests := []struct {
		name                 string
		basis                Vector
		p                    Vector
		expected             Vector
		expectedErrorMessage string
	}{
		{
			name:     "Pythagoran triangle", // a 5,4,3 triangle
			basis:    NewVector(10, 0),      // 10 across, but no height, just a straight line
			p:        NewVector(4, 3),       // 4 along and 3 up, forming the hypotenuse of the triangle.
			expected: NewVector(0, 3),       // The Project function would give (4, 0)
		},
		{
			name:     "Bigger triangle", // a 5, 12, 13 triangle
			basis:    NewVector(10, 0),  // 2 * 5 as the basis vector. The unit vector is (1, 0).
			p:        NewVector(5, 13),  // 5 across and 13 up.
			expected: NewVector(0, 13),  // Rise 13 to meet the 5, 13 vector.
		},
		{
			name:                 "Different sizes",
			basis:                NewVector(-2, -2, -1),
			p:                    NewVector(2, 2),
			expected:             Vector{},
			expectedErrorMessage: "error projecting [2, 2] onto [-2, -2, -1] with error: cannot calculate the dot product of the vectors because they have different dimensions (2 and 3)",
		},
	}

	for _, test := range tests {
		actual, err := test.basis.ProjectionOrthogonalComponent(test.p)

		if !actual.Eq(test.expected) {
			t.Errorf("%s: Projecting %v onto basis %v - expected %v, but got %v", test.name, test.p, test.basis, test.expected, actual)
		}

		if err != nil {
			if test.expectedErrorMessage == "" {
				t.Errorf("%s: Projecting %v onto basis %v, no error was expected, but got '%v'", test.name, test.p, test.basis, err)
				continue
			}

			if !strings.HasSuffix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("%s: Projecting %v onto %v, expected error message to start with '%v', but got '%v'", test.name, test.p, test.basis, test.expectedErrorMessage, err)
			}
		}
	}
}

func TestRoundingFunction(t *testing.T) {
	tests := []struct {
		input    Vector
		decimals int
		expected Vector
	}{
		{
			input:    NewVector(1.1, 2.1),
			decimals: 0,
			expected: NewVector(1.0, 2.0),
		},
		{
			input:    NewVector(1.12, 2.12),
			decimals: 1,
			expected: NewVector(1.1, 2.1),
		},
		{
			input:    NewVector(1.55, 2.55),
			decimals: 1,
			expected: NewVector(1.6, 2.6),
		},
		{
			input:    NewVector(2.4445, 2.9999),
			decimals: 3,
			expected: NewVector(2.444, 3.0),
		},
	}

	for _, test := range tests {
		actual := test.input.Round(test.decimals)

		if !actual.Eq(test.expected) {
			t.Errorf("Rounding %v to %d decimal places - expected %v, but got %v", test.input, test.decimals, test.expected, actual)
		}
	}
}

func TestCrossProductFunction(t *testing.T) {
	tests := []struct {
		name                 string
		v                    Vector
		w                    Vector
		expected             Vector
		expectedErrorMessage string
	}{
		{
			name:                 "4 dimension basis",
			v:                    NewVector(3, 3, 3, 3),
			w:                    NewVector(1, 2, 3),
			expected:             Vector{}, // Default value
			expectedErrorMessage: "the basis vector has 4 dimensions but must have 3 because cross products do not generalize to multiple dimensions",
		},
		{
			name:                 "1 dimension basis",
			v:                    NewVector(3),
			w:                    NewVector(1, 2, 3),
			expected:             Vector{}, // Default value
			expectedErrorMessage: "the basis vector has 1 dimensions but must have 3 because cross products do not generalize to multiple dimensions",
		},
		{
			name:                 "4 dimension input",
			v:                    NewVector(1, 2, 3),
			w:                    NewVector(1, 2, 3, 4),
			expected:             Vector{}, // Default value
			expectedErrorMessage: "the input vector has 4 dimensions but must have 3 because cross products do not generalize to multiple dimensions",
		},
		{
			name:                 "1 dimension input",
			v:                    NewVector(1, 2, 3),
			w:                    NewVector(1),
			expected:             Vector{}, // Default value
			expectedErrorMessage: "the input vector has 1 dimensions but must have 3 because cross products do not generalize to multiple dimensions",
		},
		{
			name:     "the cross product of two parallel vectors is the zero vector (2D)",
			v:        NewVector(3, 0, 0), // Line moving right 3 places.
			w:        NewVector(3, 0, 0), // Line moving right 3 places.
			expected: NewVector(0, 0, 0), // Zero vector
		},
		{
			name:     "the cross product of vectors at 180 degrees is the zero vector (2D)",
			v:        NewVector(3, 0, 0),  // Line moving right 3 places.
			w:        NewVector(-3, 0, 0), // Line moving left 3 places.
			expected: NewVector(0, 0, 0),  // Zero vector
		},
		{
			name:     "if either vector is zero, the result is the zero vector (2D)",
			v:        NewVector(3, 0, 0), // Line moving right 3 places.
			w:        NewVector(0, 0, 0), // Zero vector.
			expected: NewVector(0, 0, 0), // Zero vector
		},
		{
			name:     "if either vector is zero, the result is the zero vector (3D)",
			v:        NewVector(0, 0, 0), // Zero vector.
			w:        NewVector(3, 3, 3), // Non-zero vector.
			expected: NewVector(0, 0, 0), // Zero vector
		},
		{
			name:     "vectors at right angles produce a z vector (3D)",
			v:        NewVector(3, 0, 0), // 3 across.
			w:        NewVector(0, 3, 0), // 3 up.
			expected: NewVector(0, 0, 9), // z of 9.
		},
		{
			name:     "vectors at right angles produce a z vector (3D) - opposite direction",
			v:        NewVector(0, 3, 0),  // 3 up.
			w:        NewVector(3, 0, 0),  // 3 across.
			expected: NewVector(0, 0, -9), // z of -9.
		},
		{
			name:     "Udacity example at https://classroom.udacity.com/courses/ud953/lessons/4374471116/concepts/45834932820923#",
			v:        NewVector(5, 3, -2),
			w:        NewVector(-1, 0, 3),
			expected: NewVector(9, -13, 3),
		},
		{
			name:     "Khan Academy example at https://www.khanacademy.org/math/linear-algebra/vectors-and-spaces/dot-cross-products/v/linear-algebra-cross-product-introduction",
			v:        NewVector(1, -7, 1),
			w:        NewVector(5, 2, 4),
			expected: NewVector(-30, 1, 37),
		},
	}

	for _, test := range tests {
		actual, err := test.v.CrossProduct(test.w)

		if !actual.Eq(test.expected) {
			t.Errorf("%s: Calculating the cross product of %v and %v - expected %v, but got %v", test.name, test.v, test.w, test.expected, actual)
		}

		if err != nil {
			if test.expectedErrorMessage == "" {
				t.Errorf("%s: Calculating the cross product of %v and %v, no error was expected but got '%v'", test.name, test.v, test.w, err)
				continue
			}

			if !strings.HasSuffix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("%s: Calculating the cross product of %v and %v - expected error message to start with %v, but got '%v'", test.name, test.v, test.w, test.expectedErrorMessage, err)
			}
		}

		reverse, _ := test.w.CrossProduct(test.v)
		scaledReverse := reverse.Scale(-1)
		if !scaledReverse.Eq(actual) {
			t.Errorf("%s: reverse: for v (%v) and w (%v) expected (v * w) = %v = -(w * v), but got %v", test.name, test.v, test.w, actual, scaledReverse)
		}
	}
}

func TestParallelogramAreaFunction(t *testing.T) {
	tests := []struct {
		name                 string
		v                    Vector
		w                    Vector
		expected             float64
		expectedErrorMessage string
	}{
		{
			name:     "3 * 3",
			v:        NewVector(3, 0, 0),
			w:        NewVector(0, 3, 0),
			expected: 9,
		},
		{
			name:     "3 * 3 - in different orientation",
			v:        NewVector(0, 0, 3),
			w:        NewVector(0, 3, 0),
			expected: 9,
		},
		{
			name:     "Udacity example at https://classroom.udacity.com/courses/ud953/lessons/4374471116/concepts/45834932820923#",
			v:        NewVector(5, 3, -2),
			w:        NewVector(-1, 0, 3),
			expected: math.Sqrt((9 * 9) + (-13 * -13) + (3 * 3)), // 9 squared + -13 squared + 3 squared
		},
		{
			name:     "2D input",
			v:        NewVector(3, 0),
			w:        NewVector(0, 3),
			expected: 9,
		},
		{
			name:                 "4D input",
			v:                    NewVector(1, 2, 3, 4),
			w:                    NewVector(0, 3),
			expected:             0.0,
			expectedErrorMessage: "the basis vector has 4 dimensions but must have 3 because cross products do not generalize to multiple dimensions",
		},
	}

	for _, test := range tests {
		actual, err := test.v.AreaOfParallelogram(test.w)

		if err != nil {
			if test.expectedErrorMessage == "" {
				t.Errorf("%s: Calculating parallelogram area of %v and %v, no error was expected but got '%v'", test.name, test.v, test.w, err)
				continue
			}

			if !strings.HasSuffix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("%s: Calculating parallelogram area of %v and %v - expected error message to start with %v, but got '%v'", test.name, test.v, test.w, test.expectedErrorMessage, err)
			}
		}

		if actual != test.expected {
			t.Errorf("%s: Calculating parallelogram area of %v and %v - expected %v, but got %v", test.name, test.v, test.w, test.expected, actual)
		}
	}
}

func TestTriangleAreaFunction(t *testing.T) {
	tests := []struct {
		name                 string
		v                    Vector
		w                    Vector
		expected             float64
		expectedErrorMessage string
	}{
		{
			name:     "3 * 3",
			v:        NewVector(3, 0, 0),
			w:        NewVector(0, 3, 0),
			expected: 4.5,
		},
		{
			name:     "Udacity example at https://classroom.udacity.com/courses/ud953/lessons/4374471116/concepts/45834932820923#",
			v:        NewVector(5, 3, -2),
			w:        NewVector(-1, 0, 3),
			expected: math.Sqrt((9*9)+(-13*-13)+(3*3)) / 2, // 9 squared + -13 squared + 3 squared
		},
		{
			name:     "2D input",
			v:        NewVector(3, 0),
			w:        NewVector(0, 3),
			expected: 4.5,
		},
		{
			name:                 "4D input",
			v:                    NewVector(1, 2, 3, 4),
			w:                    NewVector(0, 3),
			expected:             0.0,
			expectedErrorMessage: "the basis vector has 4 dimensions but must have 3 because cross products do not generalize to multiple dimensions",
		},
	}

	for _, test := range tests {
		actual, err := test.v.AreaOfTriangle(test.w)

		if err != nil {
			if test.expectedErrorMessage == "" {
				t.Errorf("%s: Calculating triangle area of %v and %v, no error was expected but got '%v'", test.name, test.v, test.w, err)
				continue
			}

			if !strings.HasSuffix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("%s: Calculating triangle area of %v and %v - expected error message to start with %v, but got '%v'", test.name, test.v, test.w, test.expectedErrorMessage, err)
			}
		}

		if actual != test.expected {
			t.Errorf("%s: Calculating triangle area of %v and %v - expected %v, but got %v", test.name, test.v, test.w, test.expected, actual)
		}
	}
}
