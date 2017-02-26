package linear

import "testing"

func TestLineStringRepresentation(t *testing.T) {
	tests := []struct {
		name     string
		in       Line
		expected string
	}{
		{
			name:     "3x - 2y = 1",
			in:       NewLine(NewVector(3, -2), 1),
			expected: "3x_1 - 2x_2 = 1",
		},
		{
			name:     "-2x + 2x = 3",
			in:       NewLine(NewVector(-2, 2), 3),
			expected: "-2x_1 + 2x_2 = 3",
		},
		{
			name:     "Zero Vector Line",
			in:       NewLine(NewVector(0, 0), 3),
			expected: "x_1 + x_2 = 3",
		},
		{
			name:     "-4x -3y = -12",
			in:       NewLine(NewVector(-4, -3), -12),
			expected: "-4x_1 - 3x_2 = -12",
		},
	}

	for _, test := range tests {
		actual := test.in.String()

		if actual != test.expected {
			t.Errorf("%s: Expected '%v', but got '%v'", test.name, test.expected, actual)
		}
	}
}
