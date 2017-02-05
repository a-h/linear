package linear

import (
	"testing"

	"github.com/a-h/linear/tolerance"
)

func TestRadianToDegreeConversion(t *testing.T) {
	actual := Radian(1).Degrees()
	expected := float64(57.3)

	if !tolerance.IsWithin(actual, expected, tolerance.OneDecimalPlace) {
		t.Errorf("1 radian should be %f, but was %f", expected, actual)
	}
}

func TestDegreeToRadianConversion(t *testing.T) {
	actual := NewRadian(57.3)
	expected := Radian(1)

	if !tolerance.IsWithin(float64(actual), float64(expected), tolerance.OneDecimalPlace) {
		t.Errorf("1 radian should be %f, but was %f", expected, actual)
	}
}
