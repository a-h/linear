package linear

import (
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
		times                int
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
			addIndex: 0,
			toIndex:  1,
			times:    1,
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
			addIndex: 0,
			toIndex:  2,
			times:    3,
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
			times:                3,
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
			times:                3,
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
			times:                2,
			expectedErrorMessage: "cannot add vectors together because they have different dimensions (3 and 2)",
		},
	}

	for _, test := range tests {
		actual, err := test.input.Add(test.addIndex, test.toIndex, test.times)
		if err != nil {
			if test.expectedErrorMessage == "" || !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("%s: unexpected error adding index %d to index %d %d times: %v\n", test.name, test.addIndex, test.toIndex, test.times, err)
			}
			continue
		}

		eq, err := actual.Eq(test.expected)
		if err != nil && !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
			t.Errorf("%s: unexpected error adding index %d to index %d %d times: %v\n", test.name, test.addIndex, test.toIndex, test.times, err)
		}
		if !eq {
			t.Errorf("%s: expected %v, but got %v", test.name, test.expected, actual)
		}
	}
}

func TestSystemSubFunction(t *testing.T) {
	tests := []struct {
		name                 string
		input                System
		subtractIndex        int
		fromIndex            int
		times                int
		expected             System
		expectedErrorMessage string
	}{
		{
			name: "subtract the first one from the second once",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			subtractIndex: 0,
			fromIndex:     1,
			times:         1,
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
			subtractIndex: 0,
			fromIndex:     2,
			times:         3,
			expected: NewSystem(
				NewLine(NewVector(3, 3, 3), 3),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(3, 2, 1), 0),
				NewLine(NewVector(1, 0, -2), 2)),
		},
		{
			name: "source out of range",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			subtractIndex:        -1,
			fromIndex:            2,
			times:                3,
			expectedErrorMessage: "source index -1 is not present in the system",
		},
		{
			name: "destination out of range",
			input: NewSystem(
				NewLine(NewVector(1, 1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			subtractIndex:        0,
			fromIndex:            10,
			times:                3,
			expectedErrorMessage: "destination index 10 is not present in the system",
		},
		{
			name: "mismatched dimensions",
			input: NewSystem(
				NewLine(NewVector(1, 1), 1),
				NewLine(NewVector(0, 1, 0), 2),
				NewLine(NewVector(1, 1, -1), 3),
				NewLine(NewVector(1, 0, -2), 2)),
			subtractIndex:        0,
			fromIndex:            1,
			times:                2,
			expectedErrorMessage: "cannot subtract vectors because they have different dimensions (3 and 2)",
		},
	}

	for _, test := range tests {
		actual, err := test.input.Sub(test.subtractIndex, test.fromIndex, test.times)
		if err != nil {
			if test.expectedErrorMessage == "" || !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("%s: unexpected error subtracting index %d to index %d %d times: %v\n", test.name, test.subtractIndex, test.fromIndex, test.times, err)
			}
			continue
		}

		eq, err := actual.Eq(test.expected)
		if err != nil && !strings.HasPrefix(err.Error(), test.expectedErrorMessage) {
			t.Errorf("%s: unexpected error subtracting index %d to index %d %d times: %v\n", test.name, test.subtractIndex, test.fromIndex, test.times, err)
		}
		if !eq {
			t.Errorf("%s: expected %v, but got %v", test.name, test.expected, actual)
		}
	}
}
