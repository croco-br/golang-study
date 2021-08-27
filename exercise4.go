package main

import (
	"fmt"
)

type myType int

var x myType

//exercise 4
func exercise4() {

	begin(4)
	fmt.Printf("value %v\n", x)
	fmt.Printf("type %T\n", x)
	x = 42
	fmt.Printf("new value %v\n", x)
	end(4)
}
