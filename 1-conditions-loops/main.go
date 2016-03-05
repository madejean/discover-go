package main

import "fmt"
import "math/rand"

func main() {
	randomNumber := rand.Intn(100)
	if randomNumber > 50 {
		fmt.Println("my random number is", randomNumber, "and is greater than 50")
	} else {
		fmt.Println("my random number is", randomNumber, "and is not greater than 50")
	}
	var school = "Holberton School"
	if school == "Holberton School" {
		fmt.Println("I am a student of the Holberton School")
	} else {
		fmt.Println("I am not a student of the Holberton School")
	}
	var beautifulWeather = true
	if beautifulWeather == true {
		fmt.Println("It's a beautiful weather!")
	} else {
		fmt.Println("It's not")
	}
	holbertonFounders := []string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}
	for i := 0; i <= 2; i++ {
		fmt.Println(holbertonFounders[i], "is a founder")
	}
}
