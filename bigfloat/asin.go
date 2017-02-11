package bigfloat

import "math/big"

// Taken from https://github.com/robpike/ivy/blob/8c30a212d60844424baec62cf4d44c378ac6c615/value/asin.go and modified
// to export key functions.

const prec = 256
const strPi = "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679821480865132823066470938446095505822317253594081284811174502841027019385211055596446229489549303819644288109756659334461284756482337867831652712019091456485669234603486104543266482133936072602491412737245870066063155881748815209209628292540917153643678925903600113305305488204665213841469519415116094330572703657595919530921861173819326117931051185480744623799627495673518857527248912279381830119491298336733624406566430860213949463952247371907021798609437027705392171762931767523846748184676694051320005681271452635608277857713427577896091736371787214684409012249534301465495853710507922796892589235420199561121290219608640344181598136297747713099605187072113499999983729780499510597317328160963185950244594553469083026425223082533446850352619311881710100031378387528865875332083814206171776691473035982534904287554687311595628638823537875937519577818577805321712268066130019278766111959092164201989380952572010654858632788659361533818279682303019520353018529689957736225994138912497217752834791315155748572424541506959508295331168617278558890750983817546374649393192550604009277016711390098488240128583616035637076601047101819429555961989467678374494482553797747268471040475346462080466842590694912933136770289891521047521620569660240580381501935112533824300355876402474964732639141992726042699227967823547816360093417216412199245863150302861829745557067498385054945885869269956909272107975093029553211653449872027559602364806654991198818347977535663698074265425278625518184175746728909777727938000816470600161452491921732172147723501414419735685481613611573525521334757418494684385233239073941433345477624168625189835694855620992192221842725502542568876717904946016534668049886272327917860857843838279679766814541009538837863609506800642251252051173929848960841284886269456042419652850222106611863067442786220391949450471237137869609563643719172874677646575739624138908658326459958133904780275900994657640789512694683983525957098258226205224894077267194782684826014769909026401363944374553050682034962524517493996514314298091906592509372216964615157098583874105978859597729754989301617539284681382686838689427741559918559252459539594310499725246808459872736446958486538367362226260991246080512438843904512441365497627807977156914359977001296160894416948685558484063534220722258284886481584560285060168427394522674676788952521385225499546667278239864565961163548862305774564980355936345681743241125150760694794510965960940252288797108931456691368672287489405601015033086179286809208747609178249385890097149096759852613655497818931297848216829989487226588048575640142704775551323796414515237462343645428584447952658678210511413547357395231134271661021359695362314429524849371871101457654035902799344037420073105785390621983874478084784896833214457138687519435064302184531910484810053706146806749192781911979399520614196634287544406437451237181921799983910159195618146751426912397489409071864942319615679452080"

var floatOne = newFloat().SetInt64(1)
var floatTwo = newFloat().SetInt64(2)
var floatMinusOne = newFloat().SetInt64(-1)

func newFloat() *big.Float {
	return new(big.Float).SetPrec(prec)
}

var floatPi, _ = newFloat().SetString(strPi)

// Acos computes acos(x) as π/2 - asin(x).
func Acos(x *big.Float) *big.Float {
	// acos(x) = π/2 - asin(x)
	z := newFloat().Set(floatPi)
	z.Quo(z, newFloat().SetInt64(2))
	return z.Sub(z, floatAsin(x))
}

// floatAsin computes asin(x) using the formula asin(x) = atan(x/sqrt(1-x²)).
func floatAsin(x *big.Float) *big.Float {
	// The asin Taylor series converges very slowly near ±1, but our
	// atan implementation converges well for all values, so we use
	// the formula above to compute asin. But be careful when |x|=1.
	if x.Cmp(floatOne) == 0 {
		z := newFloat().Set(floatPi)
		return z.Quo(z, floatTwo)
	}
	if x.Cmp(floatMinusOne) == 0 {
		z := newFloat().Set(floatPi)
		z.Quo(z, floatTwo)
		return z.Neg(z)
	}
	z := newFloat()
	z.Mul(x, x)
	z.Sub(floatOne, z)
	z = Sqrt(z)
	z.Quo(x, z)
	return floatAtan(z)
}

// floatAtan computes atan(x) using a Taylor series. There are two series,
// one for |x| < 1 and one for larger values.
func floatAtan(x *big.Float) *big.Float {
	// atan(-x) == -atan(x). Do this up top to simplify the Euler crossover calculation.
	if x.Sign() < 0 {
		z := newFloat().Set(x)
		z = floatAtan(z.Neg(z))
		return z.Neg(z)
	}

	// The series converge very slowly near 1. atan 1.00001 takes over a million
	// iterations at the default precision. But there is hope, an Euler identity:
	//	atan(x) = atan(y) + atan((x-y)/(1+xy))
	// Note that y is a free variable. If x is near 1, we can use this formula
	// to push the computation to values that converge faster. Because
	//	tan(π/8) = √2 - 1, or equivalently atan(√2 - 1) == π/8
	// we choose y = √2 - 1 and then we only need to calculate one atan:
	//	atan(x) = π/8 + atan((x-y)/(1+xy))
	// Where do we cross over? This version converges significantly faster
	// even at 0.5, but we must be careful that (x-y)/(1+xy) never approaches 1.
	// At x = 0.5, (x-y)/(1+xy) is 0.07; at x=1 it is 0.414214; at x=1.5 it is
	// 0.66, which is as big as we dare go. With 256 bits of precision and a
	// crossover at 0.5, here are the number of iterations done by
	//	atan .1*iota 20
	// 0.1 39, 0.2 55, 0.3 73, 0.4 96, 0.5 126, 0.6 47, 0.7 59, 0.8 71, 0.9 85, 1.0 99, 1.1 116, 1.2 38, 1.3 44, 1.4 50, 1.5 213, 1.6 183, 1.7 163, 1.8 147, 1.9 135, 2.0 125
	tmp := newFloat().Set(floatOne)
	tmp.Sub(tmp, x)
	tmp.Abs(tmp)
	if tmp.Cmp(newFloat().SetFloat64(0.5)) < 0 {
		z := newFloat().Set(floatPi)
		z.Quo(z, newFloat().SetInt64(8))
		y := Sqrt(floatTwo)
		y.Sub(y, floatOne)
		num := newFloat().Set(x)
		num.Sub(num, y)
		den := newFloat().Set(x)
		den = den.Mul(den, y)
		den = den.Add(den, floatOne)
		z = z.Add(z, floatAtan(num.Quo(num, den)))
		return z
	}

	if x.Cmp(floatOne) > 0 {
		return floatAtanLarge(x)
	}

	// This is the series for small values |x| <  1.
	// asin(x) = x - x³/3 + x⁵/5 - x⁷/7 + ...
	// First term to compute in loop will be x

	n := newFloat()
	term := newFloat()
	xN := newFloat().Set(x)
	xSquared := newFloat().Set(x)
	xSquared.Mul(x, x)
	z := newFloat()

	// n goes up by two each loop.
	for loop := newLoop("atan", x, 4); ; {
		term.Set(xN)
		term.Quo(term, n.SetUint64(2*loop.i+1))
		z.Add(z, term)
		xN.Neg(xN)

		if loop.done(z) {
			break
		}
		// xN *= x², becoming x**(n+2).
		xN.Mul(xN, xSquared)
	}

	return z
}

// floatAtanLarge computes atan(x)  for large x using a Taylor series.
// x is known to be > 1.
func floatAtanLarge(x *big.Float) *big.Float {
	// This is the series for larger values |x| >=  1.
	// For x > 0, atan(x) = +π/2 - 1/x + 1/3x³ -1/5x⁵ + 1/7x⁷ - ...
	// First term to compute in loop will be -1/x

	n := newFloat()
	term := newFloat()
	xN := newFloat().Set(x)
	xSquared := newFloat().Set(x)
	xSquared.Mul(x, x)
	z := newFloat().Set(floatPi)
	z.Quo(z, floatTwo)

	// n goes up by two each loop.
	for loop := newLoop("atan", x, 4); ; {
		xN.Neg(xN)
		term.Set(xN)
		term.Mul(term, n.SetUint64(2*loop.i+1))
		term.Quo(floatOne, term)
		z.Add(z, term)

		if loop.done(z) {
			break
		}
		// xN *= x², becoming x**(n+2).
		xN.Mul(xN, xSquared)
	}

	return z
}

// Sqrt computes the square root of x using Newton's method.
// TODO: Use a better algorithm such as the one from math/sqrt.go.
func Sqrt(x *big.Float) *big.Float {
	switch x.Sign() {
	case -1:
		panic("square root of negative number")
	case 0:
		return newFloat()
	}

	// Each iteration computes
	// 	z = z - (z²-x)/2z
	// z holds the result so far. A good starting point is to halve the exponent.
	// Experiments show we converge in only a handful of iterations.
	z := newFloat()
	exp := x.MantExp(z)
	z.SetMantExp(z, exp/2)

	// Intermediates, allocated once.
	zSquared := newFloat()
	num := newFloat()
	den := newFloat()

	for loop := newLoop("sqrt", x, 1); ; {
		zSquared.Mul(z, z)
		num.Sub(zSquared, x)
		den.Mul(floatTwo, z)
		num.Quo(num, den)
		z.Sub(z, num)
		if loop.done(z) {
			break
		}
	}
	return z
}
