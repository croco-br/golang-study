package main

import (
	"fmt"
)

var xx = 42
var yy = "James Bond"
var zz = true

//exercise 3
func exercise3() {

	begin(3)
	s := fmt.Sprintf("%v\t%v\t%v\t", xx, yy, zz)
	fmt.Print(s)
	end(3)
}
