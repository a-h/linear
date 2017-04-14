Linear Algebra Functions
========================

Provides Vector operations, Equation (Line) operations and functions to work with Systems of equations to produce Triangular Form, Reduced Row-Echelon Form and carry out Gaussian elimination.

Written to complete the Udacity Linear Algebra course at https://www.udacity.com/course/linear-algebra-refresher-course--ud953

See `cmd/main.go` for how the quizes are completed.

Examples
========

```go
result, err := linear.NewVector(8.218, -9.341).Add(linear.NewVector(-1.129, 2.111))
angle, err := linear.NewVector(3.183, -7.627).AngleBetween(linear.NewVector(-2.668, 5.319))
```

```go
line1 := linear.NewEquation(linear.NewVector(7.204, 3.182), 8.68) // 7.204x + 3.182y = 8.68
line2 := linear.NewEquation(linear.NewVector(8.172, 4.114), 9.883)
intersection, intersects, equal, err := line1.IntersectionWith(line2)
```