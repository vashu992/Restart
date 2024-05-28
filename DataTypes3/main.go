package main

import "fmt"

type Student struct {
	Name           string
	Class          int
	RollNumber     string
	StudentAddress Address
}

type Address struct {
	Lane1   string
	Lane2   string
	Post    string
	Dist    string
	Village string
}

func main() {
	// Primitive approach
	// var name_vishal string
	// var class_vishal int
	// var rollnumber_vishal int

	// var name_shivam string
	// var class_shivam int
	// var rollnumber_shivam int

	// non-Primitive

	var vishal Student
	vishal.Class = 12
	vishal.Name = "vishal Singh"
	vishal.StudentAddress.Dist = "varanasi"
	vishal.StudentAddress.Village = "Ugapur"

	shivam := Student{
		Name:       "Shivam Singh",
		Class:      12,
		RollNumber: "50",
		StudentAddress: Address{
			Dist:    "District val",
			Village: "Village Val",
			Lane1:   "NH2",
			Lane2:   "National Highway",
		},
	}

	val := 122
	val2 := "12233"

	var vishalAddress *Student
	vishalAddress = &vishal

	shivamAddress := &shivam

	fmt.Println("qqqqqqqqq",shivamAddress)
	fmt.Println(vishalAddress)

	var interfaceExample interface{}

	interfaceExample = val
	fmt.Println("Interface value = ",interfaceExample)

	interfaceExample = val2
	fmt.Println("interface value = ", interfaceExample)

	interfaceExample = false
	fmt.Println("interface value = ", interfaceExample)
	fmt.Println(shivam)
}
