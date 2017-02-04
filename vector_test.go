package linear

import "testing"

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
