package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	resp, err := http.Get("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var u movie
	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		fmt.Println("err", err)
		return
	}
	p, err := strconv.ParseFloat(u.ImdbRating, 32)
	p = p * 100
	fmt.Printf("The movie : %s was released in %s - the IMBD rating is %d%% with %s votes\n", u.Title, u.Year, int(p), u.ImdbVotes)
}
