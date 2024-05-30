package main

import "fmt"

func main() {

	event := "Open door"

	if event == "Open door" {
		fmt.Println("gate kholo")
	} else if event == "close door" {
		fmt.Println("door close")
	} else {
		fmt.Println("pata nahi kya karna hai, bola hai", event)
	}

	num := -1

	if num == 0 {
		fmt.Println("value zero hai")
	} else if num < 0 {
		fmt.Println("Negative value hai")
	} else {
		fmt.Println("positive value hai", num)
	}

	number := 11
	if (number == 0 || num > 0) && (num % 2 == 0) {
		fmt.Println("value postive hai aur even bhi hai ++++++++")
	}else if (number == 0 || number > 0) && (number % 2 != 0) {
		fmt.Println("value positive hai aur odd bhi hai +++++++")
	}else if number < 0{
         fmt.Println("positive value hai", number)
	}else {
		fmt.Println("positive value hai -----------------", num)
	}
}
