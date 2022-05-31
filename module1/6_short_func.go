package module1

import "fmt"

func Exercise6() {
	// Declare 2 vars using short declaration
	// use multi to assign values
	// print only 2nd arg
	_, arg2 := multi()

	fmt.Println(arg2)
}

func multi() (int, int) {
	return 5, 6
}
