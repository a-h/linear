package main

import "fmt"
import "github.com/a-h/linear"

func main() {
	quiz1()
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
