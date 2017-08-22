package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	new_map := make(map[string]int)

	new_map["John"] = 12
	new_map["Susan"] = 29

	fmt.Println(new_map["John"])

	full := map[string]int{"Jose": 23, "Gemima": 67}

	fmt.Println(full["Jose"])
}
