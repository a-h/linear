package main

import (
	"fmt"
	"math/big"

	"github.com/a-h/linear"
)
import "flag"

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
	q3 := linear.NewVector(1.671, -1.012, -0.318).Scale(big.NewFloat(7.41))
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

		isOrthogonol, err := q.v.IsOrthogonalTo(q.w)
		if err != nil {
			fmt.Printf("Failed to answer question %d (is orthogonol) with err %v", i, err)
			return
		}

		fmt.Printf("%d: parallel: %v, orthogonol: %v\n", i, isParallel, isOrthogonol)
	}
}
