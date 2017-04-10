package linear

// Parameterization provides a way of enumerating the many possible solutions to a system of equations
// which has infinite solutions.
type Parameterization struct {
	Basepoint        Vector
	DirectionVectors []Vector
}
