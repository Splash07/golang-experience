package main

import "fmt"

type helloStruct struct {
	a int
	b string
	c bool
	// _ [0][]byte - uncomment to prevent comparison
}

func main() {
	hello := helloStruct{
		a: 10,
		b: "lol",
		c: true,
	}

	hello2 := helloStruct{
		a: 10,
		b: "lol",
		c: true,
	}

	if hello == hello2 {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}
