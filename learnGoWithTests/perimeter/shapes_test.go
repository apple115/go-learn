package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f wnat %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Area(rectangle)
	want := 100.00
	if got != want {
		t.Errorf("got %.2f wnat %.2f", got, want)
	}
}
