package linear

import "testing"

func TestParameterizationStringRepresentation(t *testing.T) {
	tests := []struct {
		input    Parameterization
		expected string
	}{
		{
			input: Parameterization{
				Basepoint:        NewVector(-1.346, 0, 0.585),
				DirectionVectors: []Vector{NewVector(-1, 1, -0)},
			},
			expected: "{ x₁ = -1.346 - t, x₂ = t, x₃ = 0.585 }",
		},
		{
			input: Parameterization{
				Basepoint: NewVector(-10.647, 0, 0),
				DirectionVectors: []Vector{
					NewVector(-1.882, 1, 0),
					NewVector(10.016, 0, 1),
				},
			},
			expected: "{ x₁ = -10.647 - 1.882t + 10.016s, x₂ = t, x₃ = s }",
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
