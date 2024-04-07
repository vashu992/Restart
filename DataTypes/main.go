package main

import "fmt"

func main() {

	// 12 -> int
	// 12.4 ->float
	// true/false -> boolean
	// c -> char/string
	// New Tech -> Collection of char (string)

	const new = 211143116


	var number int
	number = 120

	var news *int
	news =&number

	var decNum float32
	decNum = 12.222

	var addsfloat *float32
	addsfloat = &decNum


	var flags bool
	flags = true

	var addflages *bool
	addflages = &flags

	var names string
	names = "Tech Giants"

	var addNames *string
	addNames = &names

	fmt.Printf("Number value = %d , decimal value = %f,  flags value = %v , name value = %s \n",number,decNum,flags,names)
	fmt.Printf("Poiter values of news = %v , pointer values of float = %v,  pointer values of flags = %v , pointer values of names = %v \n",news,addsfloat,addflages,addNames)
	
}