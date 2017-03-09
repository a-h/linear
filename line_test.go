package linear

import (
	"strings"
	"testing"
)

func TestSubscript(t *testing.T) {
	actual := getSubscript(9876543210)
	expected := "₉₈₇₆₅₄₃₂₁₀"
	if actual != expected {
		t.Errorf("Expected %s, but got %s.", expected, actual)
	}
}

func TestLineStringRepresentation(t *testing.T) {
	tests := []struct {
		name     string
		in       Line
		expected string
	}{
		{
			name:     "3x - 2y = 1",
			in:       NewLine(NewVector(3, -2), 1),
			expected: "3x₁ - 2x₂ = 1",
		},
		{
			name:     "-2x + 2x = 3",
			in:       NewLine(NewVector(-2, 2), 3),
			expected: "-2x₁ + 2x₂ = 3",
		},
		{
			name:     "Zero Vector Line",
			in:       NewLine(NewVector(0, 0), 3),
			expected: "x₁ + x₂ = 3",
		},
		{
			name:     "-4x -3y = -12",
			in:       NewLine(NewVector(-4, -3), -12),
			expected: "-4x₁ - 3x₂ = -12",
		},
		{
			name:     "-24x -93y = -12",
			in:       NewLine(NewVector(-24, -93), -12),
			expected: "-24x₁ - 93x₂ = -12",
		},
	}

	for _, test := range tests {
		actual := test.in.String()

		if actual != test.expected {
			t.Errorf("%s: Expected '%v', but got '%v'", test.name, test.expected, actual)
		}
	}
}

func TestParallelLineFunction(t *testing.T) {
	tests := []struct {
		name     string
		a        Line
		b        Line
		expected bool
	}{
		{
			name:     "3x + 2y = 12 and 3x + 2y = 18", // Shifted by +6 on the x axis.
			a:        NewLine(NewVector(3, 2), 12),
			b:        NewLine(NewVector(3, 2), 18),
			expected: true,
		},
		{
			name:     "Scaled - 3x + 2y = 12 and 6x + 4y = 24", // Double the size.
			a:        NewLine(NewVector(3, 2), 12),
			b:        NewLine(NewVector(6, 4), 24),
			expected: true,
		},
		{
			name:     "Non-parallel lines", // Double the size.
			a:        NewLine(NewVector(2, 1), 12),
			b:        NewLine(NewVector(3, 1), 12),
			expected: false,
		},
		{
			name:     "Points", // Shifted by +6 on the x axis.
			a:        NewLine(NewVector(0, 0), 12),
			b:        NewLine(NewVector(0, 0), 18),
			expected: true,
		},
		{
			name:     "Example from Udacity (1)", // https://classroom.udacity.com/courses/ud953/lessons/4624329808/concepts/49417987180923#
			a:        NewLine(NewVector(3, -2), 1),
			b:        NewLine(NewVector(-6, 4), 0),
			expected: true,
		},
		{
			name:     "Example from Udacity (2)", // https://classroom.udacity.com/courses/ud953/lessons/4624329808/concepts/49417987180923#
			a:        NewLine(NewVector(0, 2), 3),
			b:        NewLine(NewVector(1, 1), 2),
			expected: false,
		},
	}

	for _, test := range tests {
		actual, err := test.a.IsParallelTo(test.b)

		if actual != test.expected {
			t.Errorf("%s: Expected '%v' and '%v' parallel test to be '%v', but got '%v'", test.name, test.a, test.b, test.expected, actual)
		}

		if err != nil {
			t.Errorf("%s: Got error '%v'.", test.name, err)
		}
	}
}

func TestEqualFunction(t *testing.T) {
	tests := []struct {
		name                 string
		a                    Line
		b                    Line
		expected             bool
		expectedErrorMessage string
	}{
		{
			name:     "x + y = 1 and -3x -3y = -3",
			a:        NewLine(NewVector(1, 1), 1),
			b:        NewLine(NewVector(-3, -3), -3),
			expected: true,
		},
		{
			name:     "x + y = 1 and -3x -3y = -2",
			a:        NewLine(NewVector(1, 1), 1),
			b:        NewLine(NewVector(-3, -3), -2),
			expected: false,
		},
		{
			name:     "0x + 0x = 1 and 0x + 1y = -2", // One zero vector, and a non-zero vector.
			a:        NewLine(NewVector(0, 0), 1),
			b:        NewLine(NewVector(0, 1), -2),
			expected: false,
		},
		{
			name:     "Input is zero, base is not.",
			a:        NewLine(NewVector(1, 2), 1),
			b:        NewLine(NewVector(0, 0), 1),
			expected: false,
		},
		{
			name:     "0x + 0y = 1 and 0x + 0y = -2", // 2 zero vectors, with different constants
			a:        NewLine(NewVector(0, 0), 1),
			b:        NewLine(NewVector(0, 0), -2),
			expected: false,
		},
		{
			name:     "3x + 2y = 1 and 1x + 2y = -2", // 2 non-parallel vectors.
			a:        NewLine(NewVector(3, 2), 1),
			b:        NewLine(NewVector(1, 2), -2),
			expected: false,
		},
		{
			name:                 "3x + 2y = 1 and 1x + 2y = -2", // Different size vectors.
			a:                    NewLine(NewVector(3, 2), 1),
			b:                    NewLine(NewVector(3, 2, 1), 1),
			expected:             false,
			expectedErrorMessage: "annot calculate whether the vectors are parallel because they have different dimensions (2 and 3)",
		},
	}

	for _, test := range tests {
		actual, err := test.a.Eq(test.b)

		if actual != test.expected {
			t.Errorf("%s: Expected '%v', but got '%v'", test.name, test.expected, actual)
		}

		if err != nil {
			if test.expectedErrorMessage == "" {
				t.Errorf("%s: Comparing '%v' and '%v', no error was expected but got '%v'", test.name, test.a, test.b, err)
				continue
			}

			if !strings.HasSuffix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("%s: Comparing '%v' and '%v' - expected error message to start with %v, but got '%v'", test.name, test.a, test.b, test.expectedErrorMessage, err)
			}
		}
	}
}

func TestIntersectionFunction(t *testing.T) {
	tests := []struct {
		name                 string
		a                    Line
		b                    Line
		expected             Vector
		intersects           bool
		equal                bool
		expectedErrorMessage string
	}{
		{
			name:       "1x+3y=9 and 3x+3y=9",
			a:          NewLine(NewVector(1, 3), 9),
			b:          NewLine(NewVector(3, 3), 9),
			expected:   NewVector(0, 3),
			intersects: true,
		},
		{
			name:       "1x+4y=9 and 3x+3y=9",       // 4y = -x + 9 and 3y = 9 - 3x
			a:          NewLine(NewVector(1, 4), 9), // y = -1/4x + 2.25
			b:          NewLine(NewVector(3, 3), 9), // y = 3 - x
			expected:   NewVector(1, 2),
			intersects: true,
		},
		{
			name:       "x+y=0 and x-y=0",
			a:          NewLine(NewVector(1, 1), 0),  // y = -1/4x + 2.25
			b:          NewLine(NewVector(1, -1), 0), // y = 3 - x
			expected:   NewVector(0, 0),
			intersects: true,
		},
		{
			name:       "3x + 2y = 12 and 3x + 2y = 18", // Shifted by +6 on the x axis.
			a:          NewLine(NewVector(3, 2), 12),
			b:          NewLine(NewVector(3, 2), 18),
			expected:   Vector{},
			intersects: false,
		},
		{
			name:                 "Mismatched term count (2 to 3)",
			a:                    NewLine(NewVector(3, 2), 12),
			b:                    NewLine(NewVector(3, 2, 1), 18),
			expected:             Vector{},
			intersects:           false,
			expectedErrorMessage: "The IntersectionWith function requires that both lines must have 2 dimensions. The base line has 2 dimensions, l2 has 3 dimensions.",
		},
		{
			name:       "Zero vector",
			a:          NewLine(NewVector(0, 0), 0),
			b:          NewLine(NewVector(1, 3), 0),
			expected:   Vector{},
			intersects: false,
		},
		{
			name:       "Equal lines: x + y = 1 and -3x -3y = -3",
			a:          NewLine(NewVector(1, 1), 1),
			b:          NewLine(NewVector(-3, -3), -3),
			expected:   NewVector(1, 0),
			equal:      true,
			intersects: true,
		},
	}

	for _, test := range tests {
		actual, intersects, equal, err := test.a.IntersectionWith(test.b)

		if !actual.Eq(test.expected) {
			t.Errorf("%s: Expected '%v', but got '%v'", test.name, test.expected, actual)
		}

		if intersects != test.intersects {
			t.Errorf("%s: Expected intersection to be %v, but was %v", test.name, test.intersects, intersects)
		}

		if equal != test.equal {
			t.Errorf("%s: Expected equal to be %v, but was %v", test.name, test.equal, equal)
		}

		if err != nil {
			if test.expectedErrorMessage == "" {
				t.Errorf("%s: Comparing '%v' and '%v', no error was expected but got '%v'", test.name, test.a, test.b, err)
				continue
			}

			if !strings.HasSuffix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("%s: Comparing '%v' and '%v' - expected error message to start with %v, but got '%v'", test.name, test.a, test.b, test.expectedErrorMessage, err)
			}
		}
	}
}

func TestYFunction(t *testing.T) {
	tests := []struct {
		line                 Line
		inputX               float64
		expectedY            float64
		expectedErrorMessage string
	}{
		{
			line:      NewLine(NewVector(1, 4), 9),
			inputX:    1.0,
			expectedY: 2.0,
		},
		{
			line:      NewLine(NewVector(3, 3), 9),
			inputX:    1.0,
			expectedY: 2.0,
		},
		{
			line:      NewLine(NewVector(3, 3), 9),
			inputX:    3.0,
			expectedY: 0.0,
		},
		{
			line:                 NewLine(NewVector(1), 7),
			inputX:               3.0,
			expectedY:            0.0,
			expectedErrorMessage: "The Y function only supports lines with 2 dimensions.",
		},
	}

	for _, test := range tests {
		actualY, err := test.line.Y(test.inputX)

		if actualY != test.expectedY {
			t.Errorf("For line %v. At x=%v, expected y=%v, but got (%v, %v)", test.line, test.inputX, test.expectedY, test.inputX, actualY)
		}

		if err != nil {
			if test.expectedErrorMessage == "" {
				t.Errorf("For line '%v', no error was expected but got '%v'", test.line, err)
				continue
			}

			if !strings.HasSuffix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("For line '%v' - expected error message to start with '%v', but got '%v'", test.line, test.expectedErrorMessage, err)
			}
		}
	}
}

func TestXFunction(t *testing.T) {
	tests := []struct {
		line                 Line
		inputY               float64
		expectedX            float64
		expectedErrorMessage string
	}{
		{
			line:      NewLine(NewVector(1, 4), 9),
			inputY:    2.0,
			expectedX: 1.0,
		},
		{
			line:      NewLine(NewVector(3, 3), 9),
			inputY:    2.0,
			expectedX: 1.0,
		},
		{
			line:                 NewLine(NewVector(1, 2, 3), 7),
			inputY:               3.0,
			expectedX:            0.0,
			expectedErrorMessage: "The X function only supports lines with 2 dimensions.",
		},
	}

	for _, test := range tests {
		actualX, err := test.line.X(test.inputY)

		if actualX != test.expectedX {
			t.Errorf("For line %v. At y=%v, expected x=%v, but got (%v, %v)", test.line, test.inputY, test.expectedX, actualX, test.inputY)
		}

		if err != nil {
			if test.expectedErrorMessage == "" {
				t.Errorf("For line '%v', no error was expected but got '%v'", test.line, err)
				continue
			}

			if !strings.HasSuffix(err.Error(), test.expectedErrorMessage) {
				t.Errorf("For line '%v' - expected error message to start with '%v', but got '%v'", test.line, test.expectedErrorMessage, err)
			}
		}
	}
}
