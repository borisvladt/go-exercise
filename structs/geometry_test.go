package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run(`rectangles`, func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run(`circles`, func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

func TestArea2(t *testing.T) {
	areaTests := []struct{
		name string
		shape Shape
		want float64
	} {
		{name: `Rectangle`, shape: Rectangle{12, 6}, want: 72.0},
		{name: `circle`, shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 5}, want: 36.0},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.want {
				t.Errorf("%v got %.2f want %.2f", test.shape, got, test.want)
			}
		})

	}
}