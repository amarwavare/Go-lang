package main

import "fmt"

type User struct {
	name   string
	email  string
	status bool
	age    int
}

func main() {
	fmt.Println("Function in go")
	var numbers = []int64{10, 20}
	var sum, bool = adder(numbers)
	fmt.Printf("Basic addition of %v is %v which is %v \n", numbers, sum, bool)

	user := User{name: "Amarnath", email: "amarwavare@email.com", status: true, age: 26}
	defer fmt.Println("user", user) // This will execute last

	update1 := user.setMail("newmail@email.com")
	fmt.Println("Modified mail of the user is", update1)
}

func adder(values []int64) (int64, bool) {
	var total int64
	for _, val := range values {
		total += val
	}
	return int64(total), true
}

func (user User) setMail(mail string) User {
	user.email = mail
	return user
}
