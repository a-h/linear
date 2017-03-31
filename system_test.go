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
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			b: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			expected: true,
		},
		{
			name: "unequal systems",
			a: NewSystem(
				NewLine(NewVector(1, 1, 2), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			b: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			expected: false,
		},
		{
			name: "different sizes of systems",
			a: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3)),
			b: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			expected: false,
		},
		{
			name: "different dimensions of systems",
			a: NewSystem(
				NewLine(NewVector(1, 1), 1),
				NewLine(NewVector(1, 4), 2),
				NewLine(NewVector(1, 1, -1), 3)),
			b: NewSystem(
				NewLine(NewVector(0, 0), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
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
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			moveFrom: 0,
			moveTo:   1,
			expected: NewSystem(
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
		},
		{
			name: "switch one and three",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			moveFrom: 1,
			moveTo:   3,
			expected: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(1, 0, -2), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(0, 1, 0), 2)),
		},
		{
			name: "switch three and one",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			moveFrom: 3,
			moveTo:   1,
			expected: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(1, 0, -2), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(0, 1, 0), 2)),
		},
		{
			name: "a - out of range",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			moveFrom: 4,
			moveTo:   0,
			expected: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(1, 0, -2), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(0, 1, 0), 2)),
			expectedErrorMessage: "index 4 is not present in the system",
		},
		{
			name: "b - out of range",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			moveFrom: 0,
			moveTo:   6,
			expected: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(1, 0, -2), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(0, 1, 0), 2)),
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
				NewLine(NewVector(1, 2, 3), 4)),
			expected: "{ 1x₁ + 2x₂ + 3x₃ = 4 }",
		},
		{
			name: "two equations",
			input: NewSystem(
				NewLine(NewVector(1, 2, 3), 4),
				NewLine(NewVector(5, 6, 7), 8)),
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
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			index:       0,
			coefficient: 0,
			expected: NewSystem(
				NewLine(NewVector(0, 0, 0), 0),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
		},
		{
			name: "multiply by -1",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			index:       0,
			coefficient: -1,
			expected: NewSystem(
				NewLine(NewVector(-1, -1, -1), -1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
		},
		{
			name: "multiply by 3",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			index:       3,
			coefficient: 3,
			expected: NewSystem(
				NewLine(NewVector(-1, -1, -1), -1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(3, 0, -6), 6)),
		},
		{
			name: "above range",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			index:                4,
			coefficient:          3,
			expectedErrorMessage: "index 4 is not present in the system",
		},
		{
			name: "under range",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
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
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			addIndex:    0,
			toIndex:     1,
			coefficient: 1,
			expected: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(1, 2, 1), 3),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
		},
		{
			name: "add the first to the third 3 times",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			addIndex:    0,
			toIndex:     2,
			coefficient: 3,
			expected: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(4, 4, 2), 6),
				NewLine(NewVector(1, 0, -2), 2)),
		},
		{
			name: "source out of range",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			addIndex:             -1,
			toIndex:              2,
			coefficient:          3,
			expectedErrorMessage: "source index -1 is not present in the system",
		},
		{
			name: "destination out of range",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			addIndex:             0,
			toIndex:              10,
			coefficient:          3,
			expectedErrorMessage: "destination index 10 is not present in the system",
		},
		{
			name: "mismatched dimensions",
			input: NewSystem(
				NewLine(NewVector(1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			addIndex:             0,
			toIndex:              1,
			coefficient:          2,
			expectedErrorMessage: "cannot add vectors together because they have different dimensions (3 and 2)",
		},
		{
			name: "subtract the first one from the second once",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			addIndex:    0,
			toIndex:     1,
			coefficient: -1,
			expected: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(-1, 0, -1), 1),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
		},
		{
			name: "subtract the first from the third 3 times",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(6, 5, 4), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			addIndex:    0,
			toIndex:     2,
			coefficient: -3,
			expected: NewSystem(
				NewLine(NewVector(3, 3, 3), 3),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(3, 2, 1), 0),
				NewLine(NewVector(1, 0, -2), 2)),
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
		input                System
		expected             []int
		expectedErrorMessage string
	}{
		{
			input: NewSystem(
				NewLine(NewVector(1, 0, 0), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(0, 0, 1), 3)),
			expected: []int{0, 1, 2},
		},
		{
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			expected: []int{0, 1, 0, 0},
		},
		{
			input:                NewSystem(NewLine(NewVector(0, 0, 0), 1)),
			expectedErrorMessage: "failed to find a non-zero coefficient for equation at index 0",
		},
		{
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(0, 0, 0), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			expectedErrorMessage: "failed to find a non-zero coefficient for equation at index 2",
		},
	}

	for i, test := range tests {
		actual, err := test.input.FindFirstNonZeroCoefficients()
		if err != nil {
			if test.expectedErrorMessage == "" || !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("%d: unexpected error: %v\n", i, err)
			}
			continue
		}

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("%d: expected %v, but got %v\n", i, test.expected, actual)
		}
	}
}

func TestSystemIsTriangularFormFunction(t *testing.T) {
	tests := []struct {
		name                 string
		input                System
		expected             bool
		expectedErrorMessage string
	}{
		{
			name: "opposite of triangular form",
			input: NewSystem(
				NewLine(NewVector(0, 0, 0), 1),
				NewLine(NewVector(0, 1, 1), 2),
				NewLine(NewVector(1, 1, 1), 3)),
			expected: false,
		},
		{
			name: "starts well, but the second item is incorrect",
			input: NewSystem(
				NewLine(NewVector(0, 1, 1), 1),
				NewLine(NewVector(0, 0, 0), 2),
				NewLine(NewVector(1, 1, 1), 3)),
			expected: false,
		},
		{
			name: "perfect",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 1), 2),
				NewLine(NewVector(0, 0, 1), 3)),
			expected: true,
		},
		{
			name: "zeroes",
			input: NewSystem(
				NewLine(NewVector(1, 0, 0), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(0, 0, 1), 3)),
			expected: true,
		},
		{
			name: "already triangular",
			input: NewSystem(
				NewLine(NewVector(5, 4, -1), 0),
				NewLine(NewVector(0, 10, 3), 11),
				NewLine(NewVector(0, 0, 3), 3)),
			expected: true,
		},
		{
			name: "mismatched term counts",
			input: NewSystem(
				NewLine(NewVector(0, 0), 1),
				NewLine(NewVector(0, 1, 1), 2),
				NewLine(NewVector(1, 1, 1), 3)),
			expected:             false,
			expectedErrorMessage: "all equations in a system need to have the same number of terms",
		},
		{
			name: "more terms than equations",
			input: NewSystem(
				NewLine(NewVector(0, 0, 1, 1), 1),
				NewLine(NewVector(0, 1, 1, 1), 2),
				NewLine(NewVector(1, 1, 1, 1), 3)),
			expected:             false,
			expectedErrorMessage: "the number of terms in each equation needs to match the number of terms in the system",
		},
		{
			name: "more equations than terms",
			input: NewSystem(
				NewLine(NewVector(1, 2), 1),
				NewLine(NewVector(0, 0), 1),
				NewLine(NewVector(0, 0), 2)),
			expected: true,
		},
		{
			name: "zeroes in all the right places, but it's not triangular form",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1, 1), 1),
				NewLine(NewVector(1, 0, 1, 1), 2),
				NewLine(NewVector(1, 1, 0, 1), 3)),
			expected: false,
		},
		{
			name: "system with zero coefficient for second row, second term",
			input: NewSystem(
				NewLine(NewVector(0, 1, 1), 1),
				NewLine(NewVector(0, 0, 2), 2),
				NewLine(NewVector(0, 0, 0), 3)),
			expected: true,
		},
		{
			name: "system with non-zero coefficient for second row, first term",
			input: NewSystem(
				NewLine(NewVector(0, 1, 1), 1),
				NewLine(NewVector(1, 0, 2), 2),
				NewLine(NewVector(0, 0, 0), 3)),
			expected: false,
		},
		{
			name: "system with no solution for third term",
			input: NewSystem(
				NewLine(NewVector(1, 1, 0), 1),
				NewLine(NewVector(0, 2, 0), 2),
				NewLine(NewVector(0, 0, 0), 3)),
			expected: true,
		},
	}

	for _, test := range tests {
		actual, err := test.input.IsTriangularForm()
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
				NewLine(NewVector(0, 0, 0), 1),
				NewLine(NewVector(0, 1, 1), 2),
				NewLine(NewVector(1, 1, 1), 3)),
			expected: NewSystem(
				NewLine(NewVector(1, 1, 1), 3),
				NewLine(NewVector(0, 1, 1), 2),
				NewLine(NewVector(0, 0, 0), 1)),
		},
		{
			name: "swap so that equations with non-zero coefficients are at the top (2)",
			input: NewSystem(
				NewLine(NewVector(0, 1, 1), 1),
				NewLine(NewVector(0, 0, 0), 2),
				NewLine(NewVector(1, 1, 1), 3)),
			expected: NewSystem(
				NewLine(NewVector(1, 1, 1), 3),
				NewLine(NewVector(0, 1, 1), 1),
				NewLine(NewVector(0, 0, 0), 2)),
		},
		{
			name: "no changes needed, it's already triangular",
			input: NewSystem(
				NewLine(NewVector(5, 4, -1), 0),
				NewLine(NewVector(0, 10, 3), 11),
				NewLine(NewVector(0, 0, 3), 3)),
			expected: NewSystem(
				NewLine(NewVector(5, 4, -1), 0),
				NewLine(NewVector(0, 10, 3), 11),
				NewLine(NewVector(0, 0, 3), 3)),
		},
		{
			name: "elimination",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(1, 2, 2), 2),
				NewLine(NewVector(1, 2, 3), 3)),
			expected: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 1), 1),
				NewLine(NewVector(0, 0, 1), 1)),
		},
		{
			name: "mismatched term count",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(1, 2), 2),
				NewLine(NewVector(1, 2, 3), 3)),
			expectedErrorMessage: "all equations in a system need to have the same number of terms",
		},
		{
			name: "coefficients cancel, constant term is different",
			input: NewSystem(
				NewLine(NewVector(1, 2), 1),
				NewLine(NewVector(1, 2), 2),
				NewLine(NewVector(1, 2), 3)),
			expected: NewSystem(
				NewLine(NewVector(1, 2), 1),
				NewLine(NewVector(0, 0), 1),
				NewLine(NewVector(0, 0), 2)),
		},
		{
			name: "system with no solution",
			input: NewSystem(
				NewLine(NewVector(1, 0, 1), 1),
				NewLine(NewVector(1, 0, 1), 2),
				NewLine(NewVector(1, 0, 1), 3)),
			expected: NewSystem(
				NewLine(NewVector(1, 0, 1), 1),
				NewLine(NewVector(0, 0, 0), 1),
				NewLine(NewVector(0, 0, 0), 2)),
		},
		{
			name: "system with 3 variables and only two equations",
			input: NewSystem(
				NewLine(NewVector(1, 2, 3), 1),
				NewLine(NewVector(0, 1, 2), 2)),
			expected: NewSystem(
				NewLine(NewVector(1, 2, 3), 1),
				NewLine(NewVector(0, 1, 2), 2)),
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
				NewLine(NewVector(0, 0, 0), 1),
				NewLine(NewVector(0, 1, 1), 2),
				NewLine(NewVector(1, 1, 1), 3)),
			expected: NewSystem(
				NewLine(NewVector(1, 0, 0), 1),
				NewLine(NewVector(0, 1, 1), 2),
				NewLine(NewVector(0, 0, 0), 1)),
			expectedSuccess: false,
		},
		{
			name: "mismatched term counts",
			input: NewSystem(
				NewLine(NewVector(0, 0), 1),
				NewLine(NewVector(0, 1, 1), 2),
				NewLine(NewVector(1, 1, 1), 3)),
			expected: NewSystem(
				NewLine(NewVector(1, 0, 0), 1),
				NewLine(NewVector(0, 1, 1), 2),
				NewLine(NewVector(0, 0, 0), 1)),
			expectedSuccess:      false,
			expectedErrorMessage: "all equations in a system need to have the same number of terms",
		},
		{
			name: "perfect already",
			input: NewSystem(
				NewLine(NewVector(1, 0, 0), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(0, 0, 1), 3)),
			expected: NewSystem(
				NewLine(NewVector(1, 0, 0), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(0, 0, 1), 3)),
			expectedSuccess: true,
		},
		{
			name: "remove 3rd equation from the first to complete",
			input: NewSystem(
				NewLine(NewVector(1, 0, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(0, 0, 1), 3)),
			expected: NewSystem(
				NewLine(NewVector(1, 0, 0), -2),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(0, 0, 1), 3)),
			expectedSuccess: true,
		},
		{
			name: "ensure terms become one",
			input: NewSystem(
				NewLine(NewVector(2, 0, 2), 2),
				NewLine(NewVector(0, 2, 0), 2),
				NewLine(NewVector(0, 0, 2), 3)),
			expected: NewSystem(
				NewLine(NewVector(1, 0, 0), -0.5),
				NewLine(NewVector(0, 1, 0), 1),
				NewLine(NewVector(0, 0, 1), 1.5)),
			expectedSuccess: true,
		},
		{
			name: "remove multiples",
			input: NewSystem(
				NewLine(NewVector(3, 1), 3),
				NewLine(NewVector(0, 2), 2)),
			expected: NewSystem(
				NewLine(NewVector(1, 0), 2.0/3.0),
				NewLine(NewVector(0, 1), 1)),
			expectedSuccess: true,
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
				NewLine(NewVector(0, 0, 0), 1),
				NewLine(NewVector(0, 1, 1), 2),
				NewLine(NewVector(1, 1, 1), 3)),
			expected: false,
		},
		{
			name: "triangular form, but not RREF",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 1), 2),
				NewLine(NewVector(1, 1, 1), 3)),
			expected: false,
		},
		{
			name: "mismatched terms in the system",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 1, 2), 2),
				NewLine(NewVector(1, 1, 1), 3)),
			expectedErrorMessage: "all equations in a system need to have the same number of terms",
		},
		{
			name: "ideal case",
			input: NewSystem(
				NewLine(NewVector(1, 0, 0), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(0, 0, 1), 3)),
			expected: true,
		},
		{
			name: "each equation must only have one non-zero term",
			input: NewSystem(
				NewLine(NewVector(1, 0, 0), 1),
				NewLine(NewVector(0, 1, 1), 2),
				NewLine(NewVector(0, 0, 1), 3)),
			expected: false,
		},
		{
			name: "the leading term of each equation must be one",
			input: NewSystem(
				NewLine(NewVector(1, 0, 0), 1),
				NewLine(NewVector(0, 2, 0), 2),
				NewLine(NewVector(0, 0, 1), 3)),
			expected: false,
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
