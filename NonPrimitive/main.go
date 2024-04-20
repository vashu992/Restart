package main

import "fmt"

// Non Primitive

type Student struct {
	Name string
	Class int
	Subject string
	Place Address
}

type Address struct{
	Village string
	Post string
	Dist string
	StreetNo int
}

type Phone struct {
	Model string
	Brand string
	IMEI string
	Config
}

type Laptop struct{
	Model string
	Brand string
	Config Config
	SerialNo int
}

type Config struct {
	RAM int
	ROM int
	OS  string
	Proccessor string
}

type BasicDetail struct{

}

func main() {

	//Primitive DataTypes
	var Name string
	Name = "Parth"
	fmt.Println(Name)

	var Class int
	Class = 1
	fmt.Println(Class)

	var Place string
	Place = "Varanasi"
	fmt.Println(Place) 


	//Non Primitive DataType
	var Arjun Student
	Arjun.Name = "MahaBharat"
	Arjun.Class = 1
	Arjun.Subject = "Archer"
	Arjun.Place.Dist = "Varanasi"

	Karn := Student {
		Name: "veer",
		Class: 2,
		Subject: "Science",
		Place: Address{
			Post: "ugaour",
			Dist: "Bhadohi",
		},
	} 

	val := 1222

	var tech interface{}
	tech = val
	fmt.Println(tech)
	
	 
	fmt.Println("hello team")

	fmt.Println("This is karn DataType",Karn)
	fmt.Println(Arjun)
}