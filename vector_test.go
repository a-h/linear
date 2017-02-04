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
