package perimeter

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(width, height float64) float64 {
	return (width + height) * 2
}

func Area(rectangle Rectangle) float64 {
	return (rectangle.Height * rectangle.Width)
}
