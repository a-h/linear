package linear

import "testing"

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
		name     string
		a        Line
		b        Line
		expected bool
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
			name:     "3x + 2y = 1 and 1x + 2y = -2", // Different size vectors.
			a:        NewLine(NewVector(3, 2), 1),
			b:        NewLine(NewVector(3, 2, 1), 1),
			expected: false,
		},
	}

	for _, test := range tests {
		actual, err := test.a.Eq(test.b)

		if actual != test.expected {
			t.Errorf("%s: Expected '%v', but got '%v'", test.name, test.expected, actual)
		}

		if err != nil {
			t.Errorf("%s: Unexpected error '%v'", test.name, err)
		}
	}
}
