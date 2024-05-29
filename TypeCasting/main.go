package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a int
	a = 16

	var f int
	f = 45
	var b int16

	b = int16(a)

	var c string
	c = fmt.Sprintf("%d", f)
	fmt.Println("c = ", c)
	fmt.Println(b)

	num := "123123"
	numint, err := strconv.Atoi(num)
	fmt.Println("numint = ", numint, "err = ", err)

	num = "vishal"
	numint, err = strconv.Atoi(num)
	fmt.Println("numint = ", numint, "err = ", err)

}
