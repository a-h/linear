package main

import "fmt"
import "github.com/a-h/linear"
import "flag"
import "github.com/a-h/round"

var quiz = flag.Int("quiz", 0, "The quiz to return the answers for.")

func main() {
	flag.Parse()

	fmt.Printf("Printing answers for quiz %d.\n", *quiz)

	switch *quiz {
	case 1:
		quiz1()
	case 2:
		quiz2()
	case 3:
		quiz3()
	case 4:
		quiz4()
	case 5:
		quiz5()
	case 6:
		quiz6()
	case 7:
		quiz7()
	case 8:
		quiz8()
	case 9:
		quiz9()
	case 10:
		quiz10()
	case 11:
		quiz11()
	case 12:
		quiz12()
	case 13:
		quiz13()
	default:
		fmt.Println("Quiz not found.")
	}
}

func quiz1() {
	q1, err := linear.NewVector(8.218, -9.341).Add(linear.NewVector(-1.129, 2.111))
	if err != nil {
		fmt.Println("Failed to answer question 1")
		return
	}
	fmt.Println(q1)
	q2, err := linear.NewVector(7.119, 8.215).Sub(linear.NewVector(-8.223, 0.878))
	if err != nil {
		fmt.Println("Failed to answer question 2")
		return
	}
	fmt.Println(q2)
	q3 := linear.NewVector(1.671, -1.012, -0.318).Scale(7.41)
	fmt.Println(q3)
}

func quiz2() {
	q1 := linear.NewVector(-0.221, 7.437).Magnitude()
	fmt.Println(q1)
	q2 := linear.NewVector(8.813, -1.331, -6.247).Magnitude()
	fmt.Println(q2)
	q3 := linear.NewVector(5.581, -2.136).Normalize()
	fmt.Println(q3)
	q4 := linear.NewVector(1.996, 3.108, -4.554).Normalize()
	fmt.Println(q4)
}

func quiz3() {
	q1, err := linear.NewVector(7.887, 4.138).DotProduct(linear.NewVector(-8.802, 6.776))
	if err != nil {
		fmt.Println("Failed to answer question 1")
		return
	}
	fmt.Println(q1)
	q2, err := linear.NewVector(-5.955, -4.904, -1.874).DotProduct(linear.NewVector(-4.496, -8.755, 7.103))
	if err != nil {
		fmt.Println("Failed to answer question 2")
		return
	}
	fmt.Println(q2)
	q3, err := linear.NewVector(3.183, -7.627).AngleBetween(linear.NewVector(-2.668, 5.319))
	if err != nil {
		fmt.Println("Failed to answer question 3")
		return
	}
	fmt.Println(q3)
	q4, err := linear.NewVector(7.35, 0.221, 5.188).AngleBetween(linear.NewVector(2.751, 8.259, 3.985))
	if err != nil {
		fmt.Println("Failed to answer question 4")
		return
	}
	fmt.Println(q4.Degrees())
}

func quiz4() {
	questions := []struct {
		v linear.Vector
		w linear.Vector
	}{
		{
			v: linear.NewVector(-7.579, -7.88),
			w: linear.NewVector(22.737, 23.64),
		},
		{
			v: linear.NewVector(-2.029, 9.97, 4.172),
			w: linear.NewVector(-9.231, -6.639, -7.245),
		},
		{
			v: linear.NewVector(-2.328, -7.284, -1.214),
			w: linear.NewVector(-1.821, 1.072, -2.94),
		},
		{
			v: linear.NewVector(2.118, 4.827),
			w: linear.NewVector(0, 0),
		},
	}

	for i, q := range questions {
		isParallel, err := q.v.IsParallelTo(q.w)
		if err != nil {
			fmt.Printf("Failed to answer question %d (is parallel) with err %v", i, err)
			return
		}

		isOrthogonal, err := q.v.IsOrthogonalTo(q.w)
		if err != nil {
			fmt.Printf("Failed to answer question %d (is orthogonal) with err %v", i, err)
			return
		}

		fmt.Printf("%d: parallel: %v, orthogonal: %v\n", i, isParallel, isOrthogonal)
	}
}

func quiz5() { // Coding vector projections
	questions := []struct {
		v linear.Vector
		b linear.Vector // the basis vector
	}{
		{
			v: linear.NewVector(3.039, 1.879),
			b: linear.NewVector(0.825, 2.036),
		},
		{
			v: linear.NewVector(-9.88, -3.264, -8.159),
			b: linear.NewVector(-2.155, -9.353, -9.473),
		},
		{
			v: linear.NewVector(3.009, -6.172, 3.692, -2.51),
			b: linear.NewVector(6.404, -9.144, 2.759, 8.718),
		},
	}

	for i, q := range questions {
		projection, err := q.b.Projection(q.v)
		if err != nil {
			fmt.Printf("Failed to calculate the projection for question %d with err %v", i, err)
			return
		}

		orhogonal, err := q.b.ProjectionOrthogonalComponent(q.v)
		if err != nil {
			fmt.Printf("Failed to calculate the orthogonal for question %d with err %v", i, err)
			return
		}

		fmt.Printf("%d: projection: %v, orthogonal: %v\n", i, projection.Round(3), orhogonal.Round(3))
	}
}

func quiz6() { // Coding cross products
	a, err := linear.NewVector(8.462, 7.893, -8.187).CrossProduct(linear.NewVector(6.984, -5.975, 4.778))
	if err != nil {
		fmt.Printf("Failed to calculate the cross product for question a with err %v", err)
		return
	}
	fmt.Println("a: ", a.Round(3))

	b, err := linear.NewVector(-8.987, -9.838, 5.031).AreaOfParallelogram(linear.NewVector(-4.268, -1.861, -8.866))
	if err != nil {
		fmt.Printf("Failed to calculate the cross product for question b with err %v", err)
		return
	}
	fmt.Println("b: ", round.ToEven(b, 3))

	c, err := linear.NewVector(1.5, 9.547, 3.691).AreaOfTriangle(linear.NewVector(-6.007, 0.124, 5.772))
	if err != nil {
		fmt.Printf("Failed to calculate the cross product for question c with err %v", err)
		return
	}
	fmt.Println("c: ", round.ToEven(c, 3))
}

func quiz7() { // Insections of lines
	questions := []struct {
		a linear.Line
		b linear.Line
	}{
		{
			a: linear.NewLine(linear.NewVector(4.046, 2.836), 1.21),
			b: linear.NewLine(linear.NewVector(10.115, 7.09), 3.025),
		},
		{
			a: linear.NewLine(linear.NewVector(7.204, 3.182), 8.68),
			b: linear.NewLine(linear.NewVector(8.172, 4.114), 9.883),
		},
		{
			a: linear.NewLine(linear.NewVector(1.182, 5.562), 6.744),
			b: linear.NewLine(linear.NewVector(1.773, 8.343), 9.525),
		},
	}

	for i, q := range questions {
		v, intersects, equal, err := q.a.IntersectionWith(q.b)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("%d: ", i)
		if intersects {
			fmt.Printf("%v ", v.Round(3))
		} else {
			fmt.Printf("no intersection ")
		}

		fmt.Printf("equal: %v\n", equal)
	}
}

func quiz8() { // Parallel and equal planes
	questions := []struct {
		a linear.Line
		b linear.Line
	}{
		{
			a: linear.NewLine(linear.NewVector(-0.412, 3.806, 0.728), -3.46),
			b: linear.NewLine(linear.NewVector(1.03, -9.515, -1.82), 8.65),
		},
		{
			a: linear.NewLine(linear.NewVector(2.611, 5.528, 0.283), 4.6),
			b: linear.NewLine(linear.NewVector(7.715, 8.306, 5.342), 3.76),
		},
		{
			a: linear.NewLine(linear.NewVector(-7.926, 8.625, -7.212), -7.952),
			b: linear.NewLine(linear.NewVector(-2.642, 2.875, -2.404), -2.443),
		},
	}

	for i, q := range questions {
		equal, err := q.a.Eq(q.b)
		if err != nil {
			fmt.Println(err)
			continue
		}

		parallel, err := q.a.IsParallelTo(q.b)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("%d: ", i)
		fmt.Printf("equal: %v,  parallel: %v\n", equal, parallel)
	}
}

func quiz9() { // Coding row operations
	// Converted the Python code from Udacity to match the Go I've written.
	p0 := linear.NewLine(linear.NewVector(1, 1, 1), 1)
	p1 := linear.NewLine(linear.NewVector(0, 1, 0), 2)
	p2 := linear.NewLine(linear.NewVector(1, 1, -1), 3)
	p3 := linear.NewLine(linear.NewVector(1, 0, -2), 2)

	s := linear.NewSystem(p0, p1, p2, p3)

	// Tests
	s = test(1, func() (linear.System, error) { return s.Swap(0, 1) }, linear.NewSystem(p1, p0, p2, p3))
	s = test(2, func() (linear.System, error) { return s.Swap(1, 3) }, linear.NewSystem(p1, p3, p2, p0))
	s = test(3, func() (linear.System, error) { return s.Swap(3, 1) }, linear.NewSystem(p1, p0, p2, p3))
	s = test(4, func() (linear.System, error) { return s.Multiply(0, 1) }, linear.NewSystem(p1, p0, p2, p3))
	p2_2 := linear.NewLine(linear.NewVector(-1, -1, 1), -3)
	s = test(5, func() (linear.System, error) { return s.Multiply(2, -1) }, linear.NewSystem(p1, p0, p2_2, p3))
	p1_2 := linear.NewLine(linear.NewVector(10, 10, 10), 10)
	s = test(6, func() (linear.System, error) { return s.Multiply(1, 10) }, linear.NewSystem(p1, p1_2, p2_2, p3))
	s = test(7, func() (linear.System, error) { return s.Add(0, 1, 0) }, linear.NewSystem(p1, p1_2, p2_2, p3))
	p1_3 := linear.NewLine(linear.NewVector(10, 11, 10), 12)
	s = test(8, func() (linear.System, error) { return s.Add(0, 1, 1) }, linear.NewSystem(p1, p1_3, p2_2, p3))
	p0_1 := linear.NewLine(linear.NewVector(-10, -10, -10), -10)
	s = test(9, func() (linear.System, error) { return s.Add(1, 0, -1) }, linear.NewSystem(p0_1, p1_3, p2_2, p3))
	fmt.Println("quiz 9 complete...")
}

func test(number int, operation func() (linear.System, error), expected linear.System) linear.System {
	actual, err := operation()
	if err != nil {
		fmt.Printf("test case %d failed with err: %v\n", number, err)
	}
	eq, err := actual.Eq(expected)
	if !eq {
		fmt.Printf("test case %d failed expected %v, but got %v\n", number, expected, actual)
	}
	if err != nil {
		fmt.Printf("test case %d failed to compare with %v\n", number, err)
	}
	return actual
}

func quiz10() { // Triangular form
	p1 := linear.NewLine(linear.NewVector(1, 1, 1), 1)
	p2 := linear.NewLine(linear.NewVector(0, 1, 1), 2)
	s := linear.NewSystem(p1, p2)
	test(1, func() (linear.System, error) { return s.TriangularForm() }, linear.NewSystem(p1, p2))

	p1 = linear.NewLine(linear.NewVector(1, 1, 1), 1)
	p2 = linear.NewLine(linear.NewVector(1, 1, 1), 2)
	s = linear.NewSystem(p1, p2)
	test(2, func() (linear.System, error) { return s.TriangularForm() }, linear.NewSystem(p1, linear.NewLine(linear.NewVector(0, 0, 0), 1)))

	p1 = linear.NewLine(linear.NewVector(1, 1, 1), 1)
	p2 = linear.NewLine(linear.NewVector(0, 1, 0), 2)
	p3 := linear.NewLine(linear.NewVector(1, 1, -1), 3)
	p4 := linear.NewLine(linear.NewVector(1, 0, -2), 2)
	s = linear.NewSystem(p1, p2, p3, p4)
	expected := linear.NewSystem(p1, p2, linear.NewLine(linear.NewVector(0, 0, -2), 2), linear.NewLine(linear.NewVector(), 0))
	test(3, func() (linear.System, error) { return s.TriangularForm() }, expected)

	p1 = linear.NewLine(linear.NewVector(0, 1, 1), 1)
	p2 = linear.NewLine(linear.NewVector(1, -1, 1), 2)
	p3 = linear.NewLine(linear.NewVector(1, 2, -5), 3)
	s = linear.NewSystem(p1, p2, p3)
	expected = linear.NewSystem(linear.NewLine(linear.NewVector(1, -1, 1), 2),
		linear.NewLine(linear.NewVector(0, 1, 1), 1),
		linear.NewLine(linear.NewVector(0, 0, -9), -2))
	test(4, func() (linear.System, error) { return s.TriangularForm() }, expected)
}

func quiz11() { // Coding RREF.
	p1 := linear.NewLine(linear.NewVector(1, 1, 1), 1)
	p2 := linear.NewLine(linear.NewVector(0, 1, 1), 2)
	s := linear.NewSystem(p1, p2)
	expected := linear.NewSystem(linear.NewLine(linear.NewVector(1, 0, 0), -1), p2)
	test(1, func() (linear.System, error) { r, _, err := s.ComputeRREF(); return r, err }, expected)

	p1 = linear.NewLine(linear.NewVector(1, 1, 1), 1)
	p2 = linear.NewLine(linear.NewVector(1, 1, 1), 2)
	s = linear.NewSystem(p1, p2)
	expected = linear.NewSystem(p1, linear.NewLine(linear.NewVector(0, 0, 0), 1))
	test(2, func() (linear.System, error) { r, _, err := s.ComputeRREF(); return r, err }, expected)

	p1 = linear.NewLine(linear.NewVector(1, 1, 1), 1)
	p2 = linear.NewLine(linear.NewVector(0, 1, 0), 2)
	p3 := linear.NewLine(linear.NewVector(1, 1, -1), 3)
	p4 := linear.NewLine(linear.NewVector(1, 0, -2), 2)
	s = linear.NewSystem(p1, p2, p3, p4)
	// See https://discussions.udacity.com/t/coding-rref-test-case-3/204986
	expected = linear.NewSystem(linear.NewLine(linear.NewVector(1, 0, 0), 0), p2,
		linear.NewLine(linear.NewVector(0, 0, 1), -1),
		linear.NewLine(linear.NewVector(0, 0, 0), 0))
	test(3, func() (linear.System, error) { r, _, err := s.ComputeRREF(); return r, err }, expected)

	p1 = linear.NewLine(linear.NewVector(0, 1, 1), 1)
	p2 = linear.NewLine(linear.NewVector(1, -1, 1), 2)
	p3 = linear.NewLine(linear.NewVector(1, 2, -5), 3)
	s = linear.NewSystem(p1, p2, p3)
	expected = linear.NewSystem(linear.NewLine(linear.NewVector(1, 0, 0), 23.0/9.0),
		linear.NewLine(linear.NewVector(0, 1, 0), 7.0/9.0),
		linear.NewLine(linear.NewVector(0, 0, 1), 2.0/9.0))
	test(4, func() (linear.System, error) { r, _, err := s.ComputeRREF(); return r, err }, expected)
}

func quiz12() { // Coding GE Solution
	p1 := linear.NewLine(linear.NewVector(5.862, 1.178, -10.366), -8.15)
	p2 := linear.NewLine(linear.NewVector(-2.931, -0.589, 5.183), -4.075)
	s := linear.NewSystem(p1, p2)
	solution, noSolution, infiniteSolutions, _ := s.Solve()
	fmt.Printf("q1: Solution: %v No Solution: %v Infinite Solutions: %v\n", solution, noSolution, infiniteSolutions)

	p1 = linear.NewLine(linear.NewVector(8.631, 5.112, -1.816), -5.113)
	p2 = linear.NewLine(linear.NewVector(4.315, 11.132, -5.27), -6.775)
	p3 := linear.NewLine(linear.NewVector(-2.158, 3.01, -1.727), -0.831)
	s = linear.NewSystem(p1, p2, p3)
	solution, noSolution, infiniteSolutions, _ = s.Solve()
	fmt.Printf("q2: Solution: %v No Solution: %v Infinite Solutions: %v\n", solution, noSolution, infiniteSolutions)

	p1 = linear.NewLine(linear.NewVector(5.262, 2.739, -9.878), -3.441)
	p2 = linear.NewLine(linear.NewVector(5.111, 6.358, 7.638), -2.152)
	p3 = linear.NewLine(linear.NewVector(2.016, -9.924, -1.367), -9.278)
	p4 := linear.NewLine(linear.NewVector(2.167, -13.543, -18.883), -10.567)
	s = linear.NewSystem(p1, p2, p3, p4)
	solution, noSolution, infiniteSolutions, _ = s.Solve()
	fmt.Printf("q4: Solution: %v No Solution: %v Infinite Solutions: %v\n", solution, noSolution, infiniteSolutions)
}

func quiz13() { // Coding Parameterization
	// Q1
	// There appears to be a bug in the tutorial code here, see answer at:
	// https://github.com/omarrayward/Linear-Algebra-Refresher-Udacity/blob/master/linear_system.py
	system := linear.NewSystem(
		linear.NewLine(linear.NewVector(0.786, 0.786, 0.588), -0.714),
		linear.NewLine(linear.NewVector(-0.131, -0.131, 0.244), 0.319)) // The tutorial shows -0.138, not -0.131

	solution, noSolution, infiniteSolutions, _ := system.Solve()
	fmt.Printf("q1: Solution: %v No Solution: %v Infinite Solutions: %v\n", solution, noSolution, infiniteSolutions)
	if infiniteSolutions {
		s, _, _ := system.ComputeRREF()
		fmt.Println(s.Parameterize())
	}

	// Q2
	system = linear.NewSystem(
		linear.NewLine(linear.NewVector(8.631, 5.112, -1.816), -5.113),
		linear.NewLine(linear.NewVector(4.315, 11.132, -5.27), -6.775),
		linear.NewLine(linear.NewVector(-2.158, 3.01, -1.727), -0.831))
	solution, noSolution, infiniteSolutions, _ = system.Solve()
	fmt.Printf("q2: Solution: %v No Solution: %v Infinite Solutions: %v\n", solution, noSolution, infiniteSolutions)
	if infiniteSolutions {
		s, _, _ := system.ComputeRREF()
		fmt.Println(s.Parameterize())
	}

	// Q3
	system = linear.NewSystem(
		linear.NewLine(linear.NewVector(0.935, 1.76, -9.365), -9.955),
		linear.NewLine(linear.NewVector(0.187, 0.352, -1.873), -1.991),
		linear.NewLine(linear.NewVector(0.374, 0.704, -3.746), -3.982))
	solution, noSolution, infiniteSolutions, _ = system.Solve()
	fmt.Printf("q3: Solution: %v No Solution: %v Infinite Solutions: %v\n", solution, noSolution, infiniteSolutions)
	if infiniteSolutions {
		s, _, _ := system.ComputeRREF()
		fmt.Println(s.Parameterize())
	}
}
