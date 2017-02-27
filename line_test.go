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
