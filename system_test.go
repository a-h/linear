package linear

import (
	"reflect"
	"strings"
	"testing"
)

func TestSystemEqualFunction(t *testing.T) {
	tests := []struct {
		name                 string
		a                    System
		b                    System
		expected             bool
		expectedErrorMessage string
	}{
		{
			name: "equal systems",
			a: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			b: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			expected: true,
		},
		{
			name: "unequal systems",
			a: NewSystem(
				NewEquation(NewVector(1, 1, 2), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			b: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			expected: false,
		},
		{
			name: "different sizes of systems",
			a: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3)),
			b: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			expected: false,
		},
		{
			name: "different dimensions of systems",
			a: NewSystem(
				NewEquation(NewVector(1, 1), 1),
				NewEquation(NewVector(1, 4), 2),
				NewEquation(NewVector(1, 1, -1), 3)),
			b: NewSystem(
				NewEquation(NewVector(0, 0), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			expected: false,
		},
	}

	for _, test := range tests {
		actual, err := test.a.Eq(test.b)
		if err != nil && !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
			t.Errorf("%s: unexpected error comparing s1 and s2: %v\n", test.name, err)
		}
		if actual != test.expected {
			t.Errorf("%s: expected %v, but got %v", test.name, test.expected, actual)
		}
	}
}

func TestSystemSwapFunction(t *testing.T) {
	tests := []struct {
		name                 string
		input                System
		moveFrom             int
		moveTo               int
		expected             System
		expectedErrorMessage string
	}{
		{
			name: "move from zero to one",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			moveFrom: 0,
			moveTo:   1,
			expected: NewSystem(
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
		},
		{
			name: "switch one and three",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			moveFrom: 1,
			moveTo:   3,
			expected: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(1, 0, -2), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(0, 1, 0), 2)),
		},
		{
			name: "switch three and one",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			moveFrom: 3,
			moveTo:   1,
			expected: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(1, 0, -2), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(0, 1, 0), 2)),
		},
		{
			name: "a - out of range",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			moveFrom: 4,
			moveTo:   0,
			expected: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(1, 0, -2), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(0, 1, 0), 2)),
			expectedErrorMessage: "index 4 is not present in the system",
		},
		{
			name: "b - out of range",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			moveFrom: 0,
			moveTo:   6,
			expected: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(1, 0, -2), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(0, 1, 0), 2)),
			expectedErrorMessage: "index 6 is not present in the system",
		},
	}

	for _, test := range tests {
		actual, err := test.input.Swap(test.moveFrom, test.moveTo)
		if err != nil {
			if !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("%s: unexpected error switching: %v\n", test.name, err)
			}
			continue
		}

		eq, err := actual.Eq(test.expected)
		if err != nil && !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
			t.Errorf("%s: unexpected error comparing s1 and s2: %v\n", test.name, err)
		}
		if !eq {
			t.Errorf("%s: expected %v, but got %v", test.name, test.expected, actual)
		}
	}
}

func TestSystemStringFunction(t *testing.T) {
	tests := []struct {
		name     string
		input    System
		expected string
	}{
		{
			name: "single equation",
			input: NewSystem(
				NewEquation(NewVector(1, 2, 3), 4)),
			expected: "{ 1x₁ + 2x₂ + 3x₃ = 4 }",
		},
		{
			name: "two equations",
			input: NewSystem(
				NewEquation(NewVector(1, 2, 3), 4),
				NewEquation(NewVector(5, 6, 7), 8)),
			expected: "{ 1x₁ + 2x₂ + 3x₃ = 4, 5x₁ + 6x₂ + 7x₃ = 8 }",
		},
	}

	for _, test := range tests {
		actual := test.input.String()
		if actual != test.expected {
			t.Errorf("%s: expected '%v', but got '%v'", test.name, test.expected, actual)
		}
	}
}

func TestSystemMultiplyFunction(t *testing.T) {
	tests := []struct {
		name                 string
		input                System
		coefficient          float64
		index                int
		expected             System
		expectedErrorMessage string
	}{
		{
			name: "multiply by 0",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			index:       0,
			coefficient: 0,
			expected: NewSystem(
				NewEquation(NewVector(0, 0, 0), 0),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
		},
		{
			name: "multiply by -1",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			index:       0,
			coefficient: -1,
			expected: NewSystem(
				NewEquation(NewVector(-1, -1, -1), -1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
		},
		{
			name: "multiply by 3",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			index:       3,
			coefficient: 3,
			expected: NewSystem(
				NewEquation(NewVector(-1, -1, -1), -1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(3, 0, -6), 6)),
		},
		{
			name: "above range",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			index:                4,
			coefficient:          3,
			expectedErrorMessage: "index 4 is not present in the system",
		},
		{
			name: "under range",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			index:                -1,
			coefficient:          3,
			expectedErrorMessage: "index -1 is not present in the system",
		},
	}

	for _, test := range tests {
		actual, err := test.input.Multiply(test.index, test.coefficient)
		if err != nil {
			if !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("%s: unexpected error multiplying: %v\n", test.name, err)
			}
			continue
		}

		eq, err := actual.Eq(test.expected)
		if err != nil && !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
			t.Errorf("%s: unexpected error comparing s1 and s2: %v\n", test.name, err)
		}
		if !eq {
			t.Errorf("%s: expected %v, but got %v", test.name, test.expected, actual)
		}
	}
}

func TestSystemAddFunction(t *testing.T) {
	tests := []struct {
		name                 string
		input                System
		addIndex             int
		toIndex              int
		coefficient          int
		expected             System
		expectedErrorMessage string
	}{
		{
			name: "add the first to the second once",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			addIndex:    0,
			toIndex:     1,
			coefficient: 1,
			expected: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(1, 2, 1), 3),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
		},
		{
			name: "add the first to the third 3 times",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			addIndex:    0,
			toIndex:     2,
			coefficient: 3,
			expected: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(4, 4, 2), 6),
				NewEquation(NewVector(1, 0, -2), 2)),
		},
		{
			name: "source out of range",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			addIndex:             -1,
			toIndex:              2,
			coefficient:          3,
			expectedErrorMessage: "source index -1 is not present in the system",
		},
		{
			name: "destination out of range",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			addIndex:             0,
			toIndex:              10,
			coefficient:          3,
			expectedErrorMessage: "destination index 10 is not present in the system",
		},
		{
			name: "mismatched dimensions",
			input: NewSystem(
				NewEquation(NewVector(1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			addIndex:             0,
			toIndex:              1,
			coefficient:          2,
			expectedErrorMessage: "cannot add vectors together because they have different dimensions (3 and 2)",
		},
		{
			name: "subtract the first one from the second once",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			addIndex:    0,
			toIndex:     1,
			coefficient: -1,
			expected: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(-1, 0, -1), 1),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
		},
		{
			name: "subtract the first from the third 3 times",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(6, 5, 4), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			addIndex:    0,
			toIndex:     2,
			coefficient: -3,
			expected: NewSystem(
				NewEquation(NewVector(3, 3, 3), 3),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(3, 2, 1), 0),
				NewEquation(NewVector(1, 0, -2), 2)),
		},
	}

	for _, test := range tests {
		actual, err := test.input.Add(test.addIndex, test.toIndex, test.coefficient)
		if err != nil {
			if test.expectedErrorMessage == "" || !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("%s: unexpected error adding index %d * %d to index %d: %v\n", test.name, test.addIndex, test.coefficient, test.toIndex, err)
			}
			continue
		}

		eq, err := actual.Eq(test.expected)
		if err != nil && !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
			t.Errorf("%s: unexpected error adding index %d * %d to index %d: %v\n", test.name, test.addIndex, test.coefficient, test.toIndex, err)
		}
		if !eq {
			t.Errorf("%s: expected %v, but got %v", test.name, test.expected, actual)
		}
	}
}

func TestSystemFindFirstNonZeroCoefficientsFunction(t *testing.T) {
	tests := []struct {
		input      System
		expected   []int
		expectedOK bool
	}{
		{
			input: NewSystem(
				NewEquation(NewVector(1, 0, 0), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(0, 0, 1), 3)),
			expected: []int{0, 1, 2},
		},
		{
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(1, 1, -1), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			expected: []int{0, 1, 0, 0},
		},
		{
			input:    NewSystem(NewEquation(NewVector(0, 0, 0), 1)),
			expected: []int{-1},
		},
		{
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(0, 0, 0), 3),
				NewEquation(NewVector(1, 0, -2), 2)),
			expected: []int{0, 1, -1, 0},
		},
	}

	for i, test := range tests {
		actual := test.input.FindFirstNonZeroCoefficients()

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("%d: expected %v, but got %v\n", i, test.expected, actual)
		}
	}
}

func TestSystemIsTriangularFormFunction(t *testing.T) {
	tests := []struct {
		name                          string
		input                         System
		expected                      bool
		expectedErrorMessage          string
		expectedAllLeadingTermsAreOne bool
	}{
		{
			name: "opposite of triangular form",
			input: NewSystem(
				NewEquation(NewVector(0, 0, 0), 1),
				NewEquation(NewVector(0, 1, 1), 2),
				NewEquation(NewVector(1, 1, 1), 3)),
			expected:                      false,
			expectedAllLeadingTermsAreOne: true,
		},
		{
			name: "starts well, but the second item is incorrect",
			input: NewSystem(
				NewEquation(NewVector(0, 1, 1), 1),
				NewEquation(NewVector(0, 0, 0), 2),
				NewEquation(NewVector(1, 1, 1), 3)),
			expected:                      false,
			expectedAllLeadingTermsAreOne: true,
		},
		{
			name: "perfect",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 1), 2),
				NewEquation(NewVector(0, 0, 1), 3)),
			expected:                      true,
			expectedAllLeadingTermsAreOne: true,
		},
		{
			name: "zeroes",
			input: NewSystem(
				NewEquation(NewVector(1, 0, 0), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(0, 0, 1), 3)),
			expected:                      true,
			expectedAllLeadingTermsAreOne: true,
		},
		{
			name: "already triangular",
			input: NewSystem(
				NewEquation(NewVector(5, 4, -1), 0),
				NewEquation(NewVector(0, 10, 3), 11),
				NewEquation(NewVector(0, 0, 3), 3)),
			expected:                      true,
			expectedAllLeadingTermsAreOne: false,
		},
		{
			name: "mismatched term counts",
			input: NewSystem(
				NewEquation(NewVector(0, 0), 1),
				NewEquation(NewVector(0, 1, 1), 2),
				NewEquation(NewVector(1, 1, 1), 3)),
			expected:                      false,
			expectedErrorMessage:          "all equations in a system need to have the same number of terms",
			expectedAllLeadingTermsAreOne: true,
		},
		{
			name: "more terms than equations",
			input: NewSystem(
				NewEquation(NewVector(0, 0, 1, 1), 1),
				NewEquation(NewVector(0, 1, 1, 1), 2),
				NewEquation(NewVector(1, 1, 1, 1), 3)),
			expected:                      false,
			expectedErrorMessage:          "the number of terms in each equation needs to match the number of terms in the system",
			expectedAllLeadingTermsAreOne: true,
		},
		{
			name: "more equations than terms",
			input: NewSystem(
				NewEquation(NewVector(1, 2), 1),
				NewEquation(NewVector(0, 0), 1),
				NewEquation(NewVector(0, 0), 2)),
			expected:                      true,
			expectedAllLeadingTermsAreOne: true,
		},
		{
			name: "zeroes in all the right places. the leading term is index 0 for them all, so triangular form is faked",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1, 1), 1),
				NewEquation(NewVector(1, 0, 1, 1), 2),
				NewEquation(NewVector(1, 1, 0, 1), 3)),
			expected:                      true,
			expectedAllLeadingTermsAreOne: true,
		},
		{
			name: "zeroes in the right places, but the leading term is out of order",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 1, 1), 2),
				NewEquation(NewVector(1, 0, 1, 1), 3)),
			expected:                      false,
			expectedAllLeadingTermsAreOne: true,
		},
		{
			name: "system with zero coefficient for second row, second term",
			input: NewSystem(
				NewEquation(NewVector(0, 1, 1), 1),
				NewEquation(NewVector(0, 0, 2), 2),
				NewEquation(NewVector(0, 0, 0), 3)),
			expected:                      true,
			expectedAllLeadingTermsAreOne: false,
		},
		{
			name: "system with non-zero coefficient for second row, first term",
			input: NewSystem(
				NewEquation(NewVector(0, 1, 1), 1),
				NewEquation(NewVector(1, 0, 2), 2),
				NewEquation(NewVector(0, 0, 0), 3)),
			expected:                      false,
			expectedAllLeadingTermsAreOne: true,
		},
		{
			name: "system with no solution for third term",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 0), 1),
				NewEquation(NewVector(0, 2, 0), 2),
				NewEquation(NewVector(0, 0, 0), 3)),
			expected:                      true,
			expectedAllLeadingTermsAreOne: false,
		},
		{
			name: "it's OK to have multiple leading zeroes",
			input: NewSystem(
				NewEquation(NewVector(1, 0, 0, 0), 1),
				NewEquation(NewVector(0, 0, 1, 0), 2),
				NewEquation(NewVector(0, 0, 0, 1), 3)),
			expected:                      true,
			expectedAllLeadingTermsAreOne: true,
		},
		{
			name: "the leading term can be followed by anything",
			input: NewSystem(
				NewEquation(NewVector(1, 5, 5, 5), 1),
				NewEquation(NewVector(0, 0, 1, 2), 2),
				NewEquation(NewVector(0, 0, 0, 1), 3)),
			expected:                      true,
			expectedAllLeadingTermsAreOne: true,
		},
		{
			name: "equations with more variables must be at the top",
			input: NewSystem(
				NewEquation(NewVector(0, 0, 0, 0), 1),
				NewEquation(NewVector(0, 0, 1, 2), 2),
				NewEquation(NewVector(0, 0, 0, 1), 3)),
			expected:                      false,
			expectedAllLeadingTermsAreOne: true,
		},
		{
			name: "leading terms don't need to be zero",
			input: NewSystem(
				NewEquation(NewVector(2, 0, 0, 0), 1),
				NewEquation(NewVector(0, 3, 1, 2), 2),
				NewEquation(NewVector(0, 0, 4, 1), 3)),
			expected:                      true,
			expectedAllLeadingTermsAreOne: false,
		},
	}

	for _, test := range tests {
		actual, actualAllLeadingTermsAreOne, err := test.input.IsTriangularForm()
		if err != nil {
			if test.expectedErrorMessage == "" || !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("%s: unexpected error: %v\n", test.name, err)
			}
			continue
		}

		if actual != test.expected {
			t.Errorf("%s: expected %v, but got %v", test.name, test.expected, actual)
		}

		if actualAllLeadingTermsAreOne != test.expectedAllLeadingTermsAreOne {
			t.Errorf("%s: expected all leading terms to be one: %v, but got %v", test.name, test.expectedAllLeadingTermsAreOne, actualAllLeadingTermsAreOne)
		}
	}
}

func TestSystemTriangularFormFunction(t *testing.T) {
	tests := []struct {
		name                 string
		input                System
		expected             System
		expectedErrorMessage string
	}{
		{
			name: "swap so that equations with non-zero coefficients are at the top (1)",
			input: NewSystem(
				NewEquation(NewVector(0, 0, 0), 1),
				NewEquation(NewVector(0, 1, 1), 2),
				NewEquation(NewVector(1, 1, 1), 3)),
			expected: NewSystem(
				NewEquation(NewVector(1, 1, 1), 3),
				NewEquation(NewVector(0, 1, 1), 2),
				NewEquation(NewVector(0, 0, 0), 1)),
		},
		{
			name: "swap so that equations with non-zero coefficients are at the top (2)",
			input: NewSystem(
				NewEquation(NewVector(0, 1, 1), 1),
				NewEquation(NewVector(0, 0, 0), 2),
				NewEquation(NewVector(1, 1, 1), 3)),
			expected: NewSystem(
				NewEquation(NewVector(1, 1, 1), 3),
				NewEquation(NewVector(0, 1, 1), 1),
				NewEquation(NewVector(0, 0, 0), 2)),
		},
		{
			name: "no changes needed, it's already triangular",
			input: NewSystem(
				NewEquation(NewVector(5, 4, -1), 0),
				NewEquation(NewVector(0, 10, 3), 11),
				NewEquation(NewVector(0, 0, 3), 3)),
			expected: NewSystem(
				NewEquation(NewVector(5, 4, -1), 0),
				NewEquation(NewVector(0, 10, 3), 11),
				NewEquation(NewVector(0, 0, 3), 3)),
		},
		{
			name: "elimination",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(1, 2, 2), 2),
				NewEquation(NewVector(1, 2, 3), 3)),
			expected: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 1), 1),
				NewEquation(NewVector(0, 0, 1), 1)),
		},
		{
			name: "mismatched term count",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(1, 2), 2),
				NewEquation(NewVector(1, 2, 3), 3)),
			expectedErrorMessage: "all equations in a system need to have the same number of terms",
		},
		{
			name: "coefficients cancel, constant term is different",
			input: NewSystem(
				NewEquation(NewVector(1, 2), 1),
				NewEquation(NewVector(1, 2), 2),
				NewEquation(NewVector(1, 2), 3)),
			expected: NewSystem(
				NewEquation(NewVector(1, 2), 1),
				NewEquation(NewVector(0, 0), 1),
				NewEquation(NewVector(0, 0), 2)),
		},
		{
			name: "system with no solution",
			input: NewSystem(
				NewEquation(NewVector(1, 0, 1), 1),
				NewEquation(NewVector(1, 0, 1), 2),
				NewEquation(NewVector(1, 0, 1), 3)),
			expected: NewSystem(
				NewEquation(NewVector(1, 0, 1), 1),
				NewEquation(NewVector(0, 0, 0), 1),
				NewEquation(NewVector(0, 0, 0), 2)),
		},
		{
			name: "system with 3 variables and only two equations",
			input: NewSystem(
				NewEquation(NewVector(1, 2, 3), 1),
				NewEquation(NewVector(0, 1, 2), 2)),
			expected: NewSystem(
				NewEquation(NewVector(1, 2, 3), 1),
				NewEquation(NewVector(0, 1, 2), 2)),
		},
	}

	for _, test := range tests {
		actual, err := test.input.TriangularForm()
		if err != nil {
			if test.expectedErrorMessage == "" || !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("%s: unexpected error: %v\n", test.name, err)
			}
			continue
		}

		eq, err := actual.Eq(test.expected)
		if err != nil && !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
			t.Errorf("%s: unexpected error comparing %v and %v: %v\n", test.name, test.input, actual, err)
		}
		if !eq {
			t.Errorf("%s: expected %v, but got %v", test.name, test.expected, actual)
		}
	}
}

func TestSystemReducedRowEchelonFormFunction(t *testing.T) {
	tests := []struct {
		name                 string
		input                System
		expected             System
		expectedSuccess      bool
		expectedErrorMessage string
	}{
		{
			name: "remove equation 2 from equation 1",
			input: NewSystem(
				NewEquation(NewVector(0, 0, 0), 1),
				NewEquation(NewVector(0, 1, 1), 2),
				NewEquation(NewVector(1, 1, 1), 3)),
			expected: NewSystem(
				NewEquation(NewVector(1, 0, 0), 1),
				NewEquation(NewVector(0, 1, 1), 2),
				NewEquation(NewVector(0, 0, 0), 1)),
			expectedSuccess: false,
		},
		{
			name: "mismatched term counts",
			input: NewSystem(
				NewEquation(NewVector(0, 0), 1),
				NewEquation(NewVector(0, 1, 1), 2),
				NewEquation(NewVector(1, 1, 1), 3)),
			expected: NewSystem(
				NewEquation(NewVector(1, 0, 0), 1),
				NewEquation(NewVector(0, 1, 1), 2),
				NewEquation(NewVector(0, 0, 0), 1)),
			expectedSuccess:      false,
			expectedErrorMessage: "all equations in a system need to have the same number of terms",
		},
		{
			name: "perfect already",
			input: NewSystem(
				NewEquation(NewVector(1, 0, 0), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(0, 0, 1), 3)),
			expected: NewSystem(
				NewEquation(NewVector(1, 0, 0), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(0, 0, 1), 3)),
			expectedSuccess: true,
		},
		{
			name: "remove 3rd equation from the first to complete",
			input: NewSystem(
				NewEquation(NewVector(1, 0, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(0, 0, 1), 3)),
			expected: NewSystem(
				NewEquation(NewVector(1, 0, 0), -2),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(0, 0, 1), 3)),
			expectedSuccess: true,
		},
		{
			name: "ensure terms become one",
			input: NewSystem(
				NewEquation(NewVector(2, 0, 2), 2),
				NewEquation(NewVector(0, 2, 0), 2),
				NewEquation(NewVector(0, 0, 2), 3)),
			expected: NewSystem(
				NewEquation(NewVector(1, 0, 0), -0.5),
				NewEquation(NewVector(0, 1, 0), 1),
				NewEquation(NewVector(0, 0, 1), 1.5)),
			expectedSuccess: true,
		},
		{
			name: "remove multiples",
			input: NewSystem(
				NewEquation(NewVector(3, 1), 3),
				NewEquation(NewVector(0, 2), 2)),
			expected: NewSystem(
				NewEquation(NewVector(1, 0), 2.0/3.0),
				NewEquation(NewVector(0, 1), 1)),
			expectedSuccess: true,
		},
		{
			name: "not enough non-zero terms to be able to be in RREF",
			input: NewSystem(
				NewEquation(NewVector(1, 0, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(0, 1, 0), 2)),
			expected: NewSystem(
				NewEquation(NewVector(1, 0, 1), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(0, 0, 0), 0)),
			expectedSuccess: false,
		},
		{
			name: "parallel lines",
			input: NewSystem(
				NewEquation(NewVector(1, 2), 18),
				NewEquation(NewVector(1, 2), 12)),
			expected: NewSystem(
				NewEquation(NewVector(1, 2), 18),
				NewEquation(NewVector(0, 0), -6)),
			expectedSuccess: false, // You can't solve parallel lines
		},
	}

	for _, test := range tests {
		actual, actualSuccess, err := test.input.ComputeRREF()
		if err != nil {
			if test.expectedErrorMessage == "" || !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("%s: unexpected error: %v\n", test.name, err)
			}
			continue
		}
		if actualSuccess != test.expectedSuccess {
			t.Errorf("%s: expected success %v, but got %v\n", test.name, test.expectedSuccess, actualSuccess)
		}

		eq, err := actual.Eq(test.expected)
		if err != nil && !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
			t.Errorf("%s: unexpected error comparing %v and %v: %v\n", test.name, test.input, actual, err)
		}
		if !eq {
			t.Errorf("%s: expected %v, but got %v", test.name, test.expected, actual)
		}
	}
}

func TestSystemIsReducedRowEchelonFormFunction(t *testing.T) {
	tests := []struct {
		name                 string
		input                System
		expected             bool
		expectedErrorMessage string
	}{
		{
			name: "not even triangular form",
			input: NewSystem(
				NewEquation(NewVector(0, 0, 0), 1),
				NewEquation(NewVector(0, 1, 1), 2),
				NewEquation(NewVector(1, 1, 1), 3)),
			expected: false,
		},
		{
			name: "triangular form, but not RREF",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 1), 2),
				NewEquation(NewVector(1, 1, 1), 3)),
			expected: false,
		},
		{
			name: "mismatched terms in the system",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 1, 1, 2), 2),
				NewEquation(NewVector(1, 1, 1), 3)),
			expectedErrorMessage: "all equations in a system need to have the same number of terms",
		},
		{
			name: "ideal case",
			input: NewSystem(
				NewEquation(NewVector(1, 0, 0), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(0, 0, 1), 3)),
			expected: true,
		},
		{
			name: "each equation has a leading term",
			input: NewSystem(
				NewEquation(NewVector(1, 0, 0), 1),
				NewEquation(NewVector(0, 1, 1), 2),
				NewEquation(NewVector(0, 0, 1), 3)),
			expected: true,
		},
		{
			name: "each equation has a leading term",
			input: NewSystem(
				NewEquation(NewVector(0, 1, 0), 1),
				NewEquation(NewVector(1, 0, 1), 2),
				NewEquation(NewVector(0, 0, 1), 3)),
			expected: false,
		},
		{
			name: "the leading term of each equation must be one",
			input: NewSystem(
				NewEquation(NewVector(1, 0, 0), 1),
				NewEquation(NewVector(0, 2, 0), 2),
				NewEquation(NewVector(0, 0, 1), 3)),
			expected: false,
		},
		{
			name: "parallel lines are in RREF",
			input: NewSystem(
				NewEquation(NewVector(1, 0), 12),
				NewEquation(NewVector(0, 1), 18)),
			expected: true,
		},
		{
			name: "it's OK to have some coefficients with non-zero terms, it just means that there are infinite solutions",
			input: NewSystem(
				NewEquation(NewVector(1, 0), 12),
				NewEquation(NewVector(0, 0), 18)),
			expected: true,
		},
	}

	for _, test := range tests {
		actual, err := test.input.IsRREF()
		if err != nil {
			if test.expectedErrorMessage == "" || !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("%s: unexpected error: %v\n", test.name, err)
			}
			continue
		}

		if actual != test.expected {
			t.Errorf("%s: expected %v, but got %v", test.name, test.expected, actual)
		}
	}
}

func TestSystemSolveFunction(t *testing.T) {
	tests := []struct {
		name                 string
		input                System
		noSolution           bool
		hasInfiniteSolutions bool
		expected             Vector
		expectedErrorMessage string
	}{
		{
			name: "already solved",
			input: NewSystem(
				NewEquation(NewVector(1, 0, 0), 1),
				NewEquation(NewVector(0, 1, 0), 2),
				NewEquation(NewVector(0, 0, 1), 3)),
			noSolution:           false,
			hasInfiniteSolutions: false,
			expected:             NewVector(1, 2, 3),
		},
		{
			name: "requires solving, but simple",
			input: NewSystem(
				NewEquation(NewVector(2, 0, 0), 2),
				NewEquation(NewVector(0, 2, 0), 4),
				NewEquation(NewVector(0, 0, 2), 6)),
			noSolution:           false,
			hasInfiniteSolutions: false,
			expected:             NewVector(1, 2, 3), // Just everything divided by 2
		},
		{
			name: "equal lines give infinite solutions",
			input: NewSystem(
				NewEquation(NewVector(3, 2), 12),
				NewEquation(NewVector(3, 2), 12)),
			noSolution:           false,
			hasInfiniteSolutions: true,
			expected:             NewVector(),
		},
		{
			name: "equal planes give infinite solutions",
			input: NewSystem(
				NewEquation(NewVector(3, 2, 1), 12),
				NewEquation(NewVector(6, 4, 2), 24)),
			noSolution:           false,
			hasInfiniteSolutions: true,
			expected:             NewVector(),
		},
		{
			name: "all parallel lines, they'll never intersect",
			input: NewSystem(
				NewEquation(NewVector(3, 2), 12),
				NewEquation(NewVector(3, 2), 18)),
			noSolution:           true,
			hasInfiniteSolutions: false,
			expected:             NewVector(),
		},
		{
			name: "mismatched terms triggers an error",
			input: NewSystem(
				NewEquation(NewVector(3, 2), 12),
				NewEquation(NewVector(3, 2, 1), 18)),
			noSolution:           false,
			hasInfiniteSolutions: false,
			expected:             NewVector(),
			expectedErrorMessage: "all equations in a system need to have the same number of terms",
		},
	}

	for _, test := range tests {
		actualSolution, actualNoSolution, infiniteSolutions, err := test.input.Solve()
		if err != nil {
			if test.expectedErrorMessage == "" || !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("%s: unexpected error: %v\n", test.name, err)
			}
			continue
		}

		if actualNoSolution != test.noSolution {
			t.Errorf("%s: expected noSolution: %v, but was %v", test.name, test.noSolution, actualNoSolution)
		}

		if infiniteSolutions != test.hasInfiniteSolutions {
			t.Errorf("%s: expected infiniteSolutions: %v, but was %v", test.name, test.hasInfiniteSolutions, infiniteSolutions)
		}

		if !actualSolution.Eq(test.expected) {
			t.Errorf("%s: expected %v, but got %v", test.name, test.expected, actualSolution)
		}
	}
}

func TestSystemParameterizationFunction(t *testing.T) {
	tests := []struct {
		name                 string
		input                System
		expected             Parameterization
		expectedErrorMessage string
	}{
		{
			name: "the y and z coordinates are free variables, they don't affect the outcome",
			input: NewSystem(
				NewEquation(NewVector(1, 1, 1), 1),
				NewEquation(NewVector(0, 0, 0), 1)),
			expected: Parameterization{
				Basepoint: NewVector(1, 0, 0),
				DirectionVectors: []Vector{
					NewVector(-1, 1, 0),
					NewVector(-1, 0, 1),
				},
			},
		},
		{
			name: "the x coordinate is a free variable",
			input: NewSystem(
				NewEquation(NewVector(0, 1, 0), 1),
				NewEquation(NewVector(0, 0, 1), 1)),
			expected: Parameterization{
				Basepoint: NewVector(0, 1, 1),
				DirectionVectors: []Vector{
					NewVector(1, 0, 0),
				},
			},
		},
		{
			name: "freetext.org example with 2 free variables",
			input: NewSystem(
				NewEquation(NewVector(1, 2.0/3.0, 0), 2.0/3.0),
				NewEquation(NewVector(0, 0, 0), 1)),
			expected: Parameterization{
				Basepoint: NewVector(2.0/3.0, 0, 0),
				DirectionVectors: []Vector{
					NewVector(-2.0/3.0, 1, 0),
					NewVector(0, 0, 1),
				},
			},
		},
		{
			name: "input is not RREF because the leading coefficient is 2",
			input: NewSystem(
				NewEquation(NewVector(0, 2, 0), 1),
				NewEquation(NewVector(0, 0, 1), 1)),
			expectedErrorMessage: "the system is not in RREF form",
		},
		{
			name:                 "empty systems can't be parameterized",
			input:                NewSystem(),
			expectedErrorMessage: "empty systems cannot be parameterized",
		},
		{
			name: "equations in the system must all be in the same dimension",
			input: NewSystem(
				NewEquation(NewVector(0, 2, 0), 1),
				NewEquation(NewVector(0, 0, 1, 4), 1)),
			expectedErrorMessage: "all equations in a system need to have the same number of terms",
		},
	}

	for _, test := range tests {
		actual, err := test.input.Parameterize()
		if err != nil {
			if test.expectedErrorMessage == "" || !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("%s: unexpected error: %v\n", test.name, err)
			}
			continue
		}

		if !actual.Basepoint.Eq(test.expected.Basepoint) {
			t.Errorf("%s: expected basepoint %v, but got %v", test.name, test.expected.Basepoint, actual.Basepoint)
		}

		if len(actual.DirectionVectors) != len(test.expected.DirectionVectors) {
			t.Fatalf("%s: expected %d direction vectors, but got %v results of %v", test.name, len(test.expected.DirectionVectors), len(actual.DirectionVectors), actual.DirectionVectors)
		}

		for i := range actual.DirectionVectors {
			if !actual.DirectionVectors[i].Eq(test.expected.DirectionVectors[i]) {
				t.Errorf("%s: comparing direction vectors with index [%d], expected %v, but got %v", test.name, i, test.expected.DirectionVectors[i], actual.DirectionVectors[i])
			}
		}
	}
}
