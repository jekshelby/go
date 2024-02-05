package main

import (
	"fmt"
)

func main() {
	tipeData()
	operator()
}

func tipeData() {
	fmt.Println("Hello Men") //Strings

	fmt.Println(12)  //int
	fmt.Println(1.2) //float

	//booleans
	fmt.Println(true)
	fmt.Println(false)
}

func operator() {
	// math
	fmt.Println("operasi matematika : ")
	fmt.Println(1 + 1)
	fmt.Println(1 - 1)
	fmt.Println(1 * 1)
	fmt.Println(1 / 1)
	// logic
	fmt.Println("operasi logika : ")
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)
}
