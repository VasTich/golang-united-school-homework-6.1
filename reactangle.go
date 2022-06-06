package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (rect Rectangle) CalcPerimeter() float64 {
	return 2 * (rect.Height + rect.Weight)
}

func (rect Rectangle) CalcArea() float64 {
	return rect.Height * rect.Weight
}
