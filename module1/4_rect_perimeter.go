package module1

import "fmt"

// Find the rectangle´s perimeter

func Exercise4() {
	var (
		width      int
		height     int
		perimieter int
	)
	width = 5
	height = 6

	perimieter = (width + height) * 2
	fmt.Println(perimieter)
}

func Exercise4a() {
	width := 5
	height := 6

	perimeter := (width + height) * 2
	fmt.Println(perimeter)
}
