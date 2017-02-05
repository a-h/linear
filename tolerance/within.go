package tolerance

const (
	// OneDecimalPlace represents a distance of 1dp.
	OneDecimalPlace = float64(0.15)
	// TwoDecimalPlaces represents a distance of 2dp.
	TwoDecimalPlaces = float64(0.015)
	// ThreeDecimalPlaces represents a distance of 3dp.
	ThreeDecimalPlaces = float64(0.0015)
)

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
