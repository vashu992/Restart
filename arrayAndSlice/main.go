package main

import "fmt"

func main() {
	// p := 1
	// q := 2
	// r := 3
	// var a, b, c, d, f, g, h, i int

	// a = 10

	var list [10]int
	list[0] = 10
	list[1] = 11
	list[2] = 12
	list[3] = 13
	list[4] = 14
	list[5] = 15
	list[6] = 16
	list[7] = 17
	list[8] = 18
	list[9] = 19

	fmt.Println("list of Number", list)

	list[5] = 1500
	fmt.Println("list of Number", list)

	var Player [11]string
	Player[0] = "Rohit Sharma (Captains)"
	Player[1] = "Hardik Pandya (Vice-captain)"
	Player[2] = "Virat Kohali"
	Player[3] = "Sanju Samson"
	Player[4] = "Y. Jaiswal"
	Player[5] = "Shivam Dubey"
	Player[6] = "Jasprit Bumrah"
	Player[7] = "MD. Shami"
	Player[8] = "Surya Kumar Yadav"
	Player[9] = "K L Rahul"

	fmt.Println("indian Cricket Team", Player)


	var Experiment [4]interface{}
	Experiment[0] = 1
	Experiment[1] = "Hello"
	Experiment[2] = 12.9
	Experiment[3] = true

	fmt.Println("This is experiment values",Experiment)

	var students []string

	students = append(students, "balak")
	students = append(students, "ddkh")

	fmt.Println("stuents = ",students)



}
