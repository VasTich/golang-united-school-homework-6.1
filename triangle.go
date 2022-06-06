package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (tr Triangle) CalcPerimeter() float64 {
	return 3 * tr.Side
}

func (tr Triangle) CalcArea() float64 {
	return (tr.Side * tr.Side) * (math.Sqrt(3) / 4)
}
