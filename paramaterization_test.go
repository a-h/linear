package linear

import "testing"

func TestParameterizationStringRepresentation(t *testing.T) {
	tests := []struct {
		input    Parameterization
		expected string
	}{
		{
			input: Parameterization{
				Basepoint:        NewVector(-1.3458774161867777, 0, 0.5847953216374271),
				DirectionVectors: []Vector{NewVector(-1, 1, -0)},
			},
			expected: "{ x = -1.346t - t, y = t, z = 0.585 }",
		},
	}

	for _, test := range tests {
		actual := test.input.String()
		if actual != test.expected {
			t.Errorf("for input of '%v', expected '%v', but got '%v'", test.input, test.expected, actual)
		}
	}
}

func TestParameterizationGetFreeVariableNameFunction(t *testing.T) {
	tests := []struct {
		index    int
		expected string
	}{
		{
			index:    0,
			expected: "t",
		},
		{
			index:    1,
			expected: "s",
		},
		{
			index:    2,
			expected: "t₁",
		},
		{
			index:    3,
			expected: "s₁",
		},
		{
			index:    4,
			expected: "t₂",
		},
		{
			index:    5,
			expected: "s₂",
		},
	}

	for _, test := range tests {
		actual := getFreeVariableName(test.index)
		if actual != test.expected {
			t.Errorf("for input index of '%v', expected '%v', but got '%v'", test.index, test.expected, actual)
		}
	}
}
