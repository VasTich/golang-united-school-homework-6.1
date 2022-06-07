package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes)+1 > b.shapesCapacity {
		return fmt.Errorf("add shape: out of capacity")
	}

	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, fmt.Errorf("get by index: out of range")
	}

	shape := b.shapes[i]
	if shape == nil {
		return nil, fmt.Errorf("get by index: shape doesn't exist")
	}

	return shape, nil

}

func removeShape(src []Shape, i int) []Shape {
	var dest []Shape
	dest = src[:i+copy(src[i:], src[i+1:])]
	return dest
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, fmt.Errorf("extract by index: out of range")
	}
	var shape Shape = nil
	for j, v := range b.shapes {
		if j == i {
			shape = v
		}
	}

	if shape == nil {
		return nil, fmt.Errorf("extract by index: shape doesn't exist")
	}

	b.shapes = removeShape(b.shapes, i)

	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, fmt.Errorf("replace by index: out of range")
	}
	oldShape := b.shapes[i]
	if b.shapes[i] == nil {
		return nil, fmt.Errorf("replace by index: shape doesn't exist")
	}

	b.shapes[i] = shape
	return oldShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64 = 0.0
	for _, v := range b.shapes {
		if v == nil {
			continue
		}
		sum = sum + v.CalcPerimeter()
	}

	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64 = 0.0
	for _, v := range b.shapes {
		if v == nil {
			continue
		}
		sum = sum + v.CalcArea()
	}

	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	found := false
	for i := 0; i < len(b.shapes); i++ {
		if _, ok := b.shapes[i].(Circle); ok {
			found = true
			b.shapes = removeShape(b.shapes, i)
			i = i - 1
		}

		if _, ok := b.shapes[i].(*Circle); ok {
			found = true
			b.shapes = removeShape(b.shapes, i)
			i = i - 1
		}
	}

	if !found {
		return fmt.Errorf("remove all circles: circles are not exist")
	}

	return nil
}
