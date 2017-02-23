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
