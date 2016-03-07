package main

import (
	"fmt"
	"time"
)

func main() {
	type user struct {
		Name string `json:"name"`
		DOB  string `json:"date_of_birth"`
		City string `json:"city"`
	}
	u := user{Name: "Betty Holberton", DOB: "March 7, 1917", City: "Philadelphia"}
	fmt.Println("Hello", u.Name)
	t, err := time.Parse("January 2, 2006", u.DOB)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s who was born in %s would be %d years old today.\n", u.Name, u.City, time.Now().Year()-t.Year())
	}
}
