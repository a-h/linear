package bigfloat

// Taken from https://github.com/robpike/ivy/

import (
	"fmt"
	"math/big"
)

type loop struct {
	name          string     // The name of the function we are evaluating.
	i             uint64     // Loop count.
	maxIterations uint64     // When to give up.
	arg           *big.Float // original argument to function; only used for diagnostic.
	prevZ         *big.Float // Result from the previous iteration.
	delta         *big.Float // |Change| from previous iteration.
}

// newLoop returns a new loop checker. The arguments are the name
// of the function being evaluated, the argument to the function, and
// the maximum number of iterations to perform before giving up.
// The last number in terms of iterations per bit, so the caller can
// ignore the precision setting.
func newLoop(name string, x *big.Float, itersPerBit uint) *loop {
	return &loop{
		name:          name,
		arg:           newFloat().Set(x),
		maxIterations: 10 + uint64(itersPerBit*prec),
		prevZ:         newFloat(),
		delta:         newFloat(),
	}
}

// done reports whether the loop is done. If it does not converge
// after the maximum number of iterations, it errors out.
func (l *loop) done(z *big.Float) bool {
	l.delta.Sub(l.prevZ, z)
	sign := l.delta.Sign()
	if sign == 0 {
		return true
	}
	if sign < 0 {
		l.delta.Neg(l.delta)
	}
	// Check if delta is no bigger than the smallest change in z that can be
	// represented with the given precision.
	var eps big.Float
	eps.SetMantExp(eps.SetUint64(1), z.MantExp(nil)-int(z.Prec()))
	if l.delta.Cmp(&eps) <= 0 {
		return true
	}
	l.i++
	if l.i == l.maxIterations {
		// Users should never see this.
		panic(fmt.Sprintf("%s %s: did not converge after %d iterations; prev,last result %s,%s delta %s", l.name, l.arg, l.maxIterations, z, l.prevZ, l.delta))
	}
	l.prevZ.Set(z)
	return false
}
