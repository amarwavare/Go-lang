package main

import "fmt"

var PublicVar string = "A public variable can be var & const, starts with uppercase in naming convention."

func main() {
	const user string = "some user"
	const booleanData bool = bool(user != "some user") || true
	const smallInt int8 = 120
	const smallUint uint8 = 200
	var initInt int    // Cannot use const for declaring. Hence var works
	var initStr string // Cannot use const for declaring. Hence var works
	PublicVar = "Modified public variable"
	unsignedVariable := "likely a variable of type any"

	fmt.Printf("user `%v` is of type `%T` \n", user, user)
	fmt.Printf("booleanData `%v` is of type `%T` \n", booleanData, booleanData)
	fmt.Printf("smallInt `%v` is of type `%T` \n", smallInt, smallInt)
	fmt.Printf("smallUint `%v` is of type `%T` \n", smallUint, smallUint)
	fmt.Printf("initInt `%v` is of type `%T` \n", initInt, initInt)
	fmt.Printf("initStr `%v` is of type `%T` \n", initStr, initStr)
	fmt.Printf("PublicVar `%v` is of type `%T` \n", PublicVar, PublicVar)
	fmt.Printf("unsignedVariable `%v` is of type `%T` \n", unsignedVariable, unsignedVariable)
}
