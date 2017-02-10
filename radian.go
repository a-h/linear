package linear

import "math"

// Radian is a unit of measurement for angles: 1 radian = 180/π
type Radian float64

// NewRadian creates a radian from degrees.
func NewRadian(degrees float64) Radian {
	radians := degrees * (math.Pi / float64(180))
	return Radian(radians)
}

// Degrees converts from radians to degrees.
func (r Radian) Degrees() float64 {
	return float64(r) * (180 / math.Pi)
}
