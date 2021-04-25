package main

import(
	"fmt"
)

func main() {
	myArray := make([]string, 0)
	secondaryArray := make([]int, 0)

	for i :=0; i<=10; i++{
		myArray = append(myArray, "Tes")
		secondaryArray = append(secondaryArray, i)
	}

	fmt.Println(secondaryArray)

	for i:= range myArray{
		fmt.Println(myArray[i]," bre, no.", i)
	}
	
}