package main

import "fmt"

func main() {
	var pointer *int
	println("Default value of pointer ", pointer)

	number := 23

	var ptrNum = &number
	fmt.Printf("Memory address of ptrNum %v and actual value is %v \n", ptrNum, *ptrNum)

	*ptrNum = *ptrNum * 2
	fmt.Printf("Value of ptrNum after x2 = %v \n", *ptrNum)
}
