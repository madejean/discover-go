package main

import "fmt"

func main() {
	type user struct {
		Name string `json: "name"`
		DOB  string `json: "date_of_birth, string"`
		City string `json: "city"`
	}
	fmt.Println("Hello")
}
