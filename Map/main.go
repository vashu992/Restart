package main

import "fmt"

func main() {

	// map decleration
	m := make(map[string]int)
	// save data in map
	m["xyz"] = 10
	m["pqr"] = 12

	// read data from map
	val, ok := m["pqr"]
	fmt.Println("val = ", val, "ok = ", ok)

	// my
	val, ok = m["pqr"]
	fmt.Println("val = ", val, "ok = ", ok)

	// Condition statement
	val, ok = m["pqr"]
	if ok {
		fmt.Println("value found, val = ", val)
	} else {
		fmt.Println("value not found")
	}

	// Experiment No. 1
	o := make(map[int]int)

	o[10] = 5

	val, ok = o[10]
	fmt.Println(val, ok)

	// Experiment No.2

	w := make(map[string]string)

	w["india"] = "Won the Cup"

	val, ok = m["india"]
	fmt.Println("India win the word cup = ", val, "ok = ", ok)
}
