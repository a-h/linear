package tolerance

import (
	"math/big"
	"testing"
)

func TestIsWithin(t *testing.T) {
	tests := []struct {
		a         float64
		b         float64
		tolerance float64
		expected  bool
	}{
		{
			a:         float64(1),
			b:         float64(1.1),
			tolerance: OneDecimalPlace,
			expected:  true,
		},
		{
			a:         float64(1),
			b:         float64(1.2),
			tolerance: OneDecimalPlace,
			expected:  false,
		},
		{
			a:         float64(1),
			b:         float64(1.11),
			tolerance: OneDecimalPlace,
			expected:  true,
		},
		{
			a:         float64(1),
			b:         float64(1.16),
			tolerance: OneDecimalPlace,
			expected:  false,
		},
		{
			a:         float64(-1),
			b:         float64(-1.1),
			tolerance: OneDecimalPlace,
			expected:  true,
		},
		{
			a:         float64(-1),
			b:         float64(-1.2),
			tolerance: OneDecimalPlace,
			expected:  false,
		},
		{
			a:         float64(-1),
			b:         float64(-1.05),
			tolerance: OneDecimalPlace,
			expected:  true,
		},
		{
			a:         float64(-1.5),
			b:         float64(1.5),
			tolerance: float64(3),
			expected:  true,
		},
		{
			a:         float64(-1.55),
			b:         float64(-1.5),
			tolerance: OneDecimalPlace,
			expected:  true,
		},
		{
			a:         float64(-1.5),
			b:         float64(-1.55),
			tolerance: OneDecimalPlace,
			expected:  true,
		},
		{
			a:         float64(3.14159265359),
			b:         float64(3.1),
			tolerance: OneDecimalPlace,
			expected:  true,
		},
		{
			a:         float64(3.14159265359),
			b:         float64(3.1),
			tolerance: TwoDecimalPlaces,
			expected:  false,
		},
		{
			a:         float64(3.14159265359),
			b:         float64(3.14),
			tolerance: TwoDecimalPlaces,
			expected:  true,
		},
		{
			a:         float64(3.14159265359),
			b:         float64(3.14),
			tolerance: ThreeDecimalPlaces,
			expected:  false,
		},
		{
			a:         float64(3.14159265359),
			b:         float64(3.141),
			tolerance: ThreeDecimalPlaces,
			expected:  true,
		},
	}

	for _, test := range tests {
		actual := IsWithin(test.a, test.b, test.tolerance)

		if actual != test.expected {
			t.Errorf("For ||%f-%f|| within %f, expected %v, but got %v", test.a, test.b, test.tolerance, test.expected, actual)
		}
	}
}

func TestIsWithinBig(t *testing.T) {
	tests := []struct {
		a         *big.Float
		b         *big.Float
		tolerance float64
		expected  bool
	}{
		{
			a:         big.NewFloat(1.0),
			b:         big.NewFloat(1.1),
			tolerance: OneDecimalPlace,
			expected:  true,
		},
		{
			a:         big.NewFloat(1),
			b:         big.NewFloat(1.2),
			tolerance: OneDecimalPlace,
			expected:  false,
		},
		{
			a:         big.NewFloat(1),
			b:         big.NewFloat(1.11),
			tolerance: OneDecimalPlace,
			expected:  true,
		},
		{
			a:         big.NewFloat(1),
			b:         big.NewFloat(1.16),
			tolerance: OneDecimalPlace,
			expected:  false,
		},
		{
			a:         big.NewFloat(-1),
			b:         big.NewFloat(-1.1),
			tolerance: OneDecimalPlace,
			expected:  true,
		},
		{
			a:         big.NewFloat(-1),
			b:         big.NewFloat(-1.2),
			tolerance: OneDecimalPlace,
			expected:  false,
		},
		{
			a:         big.NewFloat(-1),
			b:         big.NewFloat(-1.05),
			tolerance: OneDecimalPlace,
			expected:  true,
		},
		{
			a:         big.NewFloat(-1.5),
			b:         big.NewFloat(1.5),
			tolerance: float64(3),
			expected:  true,
		},
		{
			a:         big.NewFloat(-1.55),
			b:         big.NewFloat(-1.5),
			tolerance: OneDecimalPlace,
			expected:  true,
		},
		{
			a:         big.NewFloat(-1.5),
			b:         big.NewFloat(-1.55),
			tolerance: OneDecimalPlace,
			expected:  true,
		},
		{
			a:         big.NewFloat(3.14159265359),
			b:         big.NewFloat(3.1),
			tolerance: OneDecimalPlace,
			expected:  true,
		},
		{
			a:         big.NewFloat(3.14159265359),
			b:         big.NewFloat(3.1),
			tolerance: TwoDecimalPlaces,
			expected:  false,
		},
		{
			a:         big.NewFloat(3.14159265359),
			b:         big.NewFloat(3.14),
			tolerance: TwoDecimalPlaces,
			expected:  true,
		},
		{
			a:         big.NewFloat(3.14159265359),
			b:         big.NewFloat(3.14),
			tolerance: ThreeDecimalPlaces,
			expected:  false,
		},
		{
			a:         big.NewFloat(3.14159265359),
			b:         big.NewFloat(3.141),
			tolerance: ThreeDecimalPlaces,
			expected:  true,
		},
		{
			a:         big.NewFloat(3.14159265359),
			b:         big.NewFloat(3.1415926536),
			tolerance: DecimalPlaces(10),
			expected:  true,
		},
	}

	for _, test := range tests {
		actual := IsWithinBig(test.a, test.b, test.tolerance)

		if actual != test.expected {
			t.Errorf("For ||%f-%f|| within %f, expected %v, but got %v", test.a, test.b, test.tolerance, test.expected, actual)
		}
	}
}
