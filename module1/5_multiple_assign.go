package module1

import (
	"fmt"
)

func Exercise5() {
	var (
		lang    string
		version int
	)

	lang, version = "go", 2

	fmt.Println(lang, "version", version)
}

func Exercise5a() {
	// Expected output
	// Air is good on Mars
	// It´s true
	// It is 19.5 degrees

	planet := "Mars"
	isTrue := true
	temp := 19.5

	fmt.Printf("Air is good on %s\n", planet)
	fmt.Printf("It´s %t\n", isTrue)
	fmt.Printf("It is %f degrees\n", temp)
}
