package main

import "fmt"

func main() {
	var age int = 30 // variable decleration explicit type
	fmt.Println("Age:", age)

	var name = "Dan" // variable with implied type
	fmt.Println("Your name is:", name)

	_ = name

	// Short decleration operator
	s := "Learning Golang!"
	fmt.Println(s)

	car, cost := "Audi", 50000
	fmt.Println(car, cost)
	car, year := "BMW", 2018
	_ = year

	var opened = false
	opened, file := true, "a.txt"

	_, _ = opened, file // Mute errors related to unused variables

	// Multiple Assignments
	var (
		salary    float64
		firstName string
		gender    bool
	)
	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8

	j, i = i, j // Swapping variables

	fmt.Println(i, j)

	// Using math operators in variables
	sum := 5 + 2.3
	fmt.Println(sum)

}
