package golang_united_school_homework

import (
	"errors"
)

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
	if len(b.shapes) == b.shapesCapacity {
		return errors.New("shapes capacity is full")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i > len(b.shapes)-1 {
		return nil, errors.New("index is out of range")
	}
	if i > b.shapesCapacity {
		return nil, errors.New("index doesn't exist")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i > len(b.shapes)-1 {
		return nil, errors.New("index is out of range")
	}
	if i > b.shapesCapacity {
		return nil, errors.New("index doesn't exist")
	}
	extracted := b.shapes[i]
	updated := b.shapes[:i]
	if i != len(b.shapes)-1 {
		updated = append(updated, b.shapes[i+1:]...)
		b.shapes = updated
		return extracted, nil
	}
	b.shapes = updated
	return extracted, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i > len(b.shapes)-1 {
		return nil, errors.New("index is out of range")
	}
	if i > b.shapesCapacity {
		return nil, errors.New("index doesn't exist")
	}
	extracted := b.shapes[i]
	b.shapes[i] = shape
	return extracted, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sumP float64 = 0
	for i := 0; i < len(b.shapes); i++ {
		sumP += b.shapes[i].CalcPerimeter()
	}
	return sumP
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sumA float64
	for i := 0; i < len(b.shapes); i++ {
		sumA += b.shapes[i].CalcArea()
	}
	return sumA
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	updated := make([]Shape, 0, len(b.shapes))
	for i := 0; i < len(b.shapes); i++ {
		extracted, ok := b.shapes[i].(*Circle)
		if !ok {
			updated = append(updated, extracted)
		}
	}
	if len(updated) == len(b.shapes) {
		return errors.New("there are no circles in list")
	}
	b.shapes = updated
	return nil
}
