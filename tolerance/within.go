package tolerance

import "math/big"

const (
	// OneDecimalPlace represents a distance of 1dp.
	OneDecimalPlace = float64(0.15)
	// TwoDecimalPlaces represents a distance of 2dp.
	TwoDecimalPlaces = float64(0.015)
	// ThreeDecimalPlaces represents a distance of 3dp.
	ThreeDecimalPlaces = float64(0.0015)
)

// DecimalPlaces is the number of decimal places to use.
func DecimalPlaces(number int) float64 {
	return 1.5 / float64(number*10)
}

// IsWithin returns true when the parameters are within the given tolerance from each other.
func IsWithin(a float64, b float64, tolerance float64) bool {
	return distance(a, b) <= tolerance
}

func distance(a float64, b float64) float64 {
	if a > b {
		return a - b
	}
	return b - a
}

// IsWithinBig returns true when the parameters are within the given tolerance from each other.
func IsWithinBig(a *big.Float, b *big.Float, tolerance float64) bool {
	return distanceBig(a, b).Cmp(big.NewFloat(tolerance)) <= 0
}

func distanceBig(a *big.Float, b *big.Float) *big.Float {
	r := &big.Float{}

	if a.Cmp(b) > 0 {
		return r.Sub(a, b)
	}
	return r.Sub(b, a)
}
