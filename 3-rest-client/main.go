package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	type movie struct {
		Title      string `json:"Title"`
		Year       string `json:"Year"`
		Rated      string `json:"Rated"`
		Released   string `json:"Released"`
		Runtime    string `json:"Runtime"`
		Genre      string `json:"Genre"`
		Director   string `json:"Director"`
		Writer     string `json:"Writer"`
		Actors     string `json:"Actors"`
		Plot       string `json:"Plot"`
		Language   string `json:"Language"`
		Country    string `json:"Country"`
		Awards     string `json:"Awards"`
		Poster     string `json:"Poster"`
		Metascore  string `json:"Metascore"`
		ImdbRating string `json:"imdbRating"`
		ImdbVotes  string `json:"imdbVotes"`
		ImdbID     string `json:"imdbID"`
		Type       string `json:"Type"`
		Response   string `json:"Response"`
	}

	resp, err := http.Get("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println("status code is", resp.StatusCode, "OK")

	var u movie
	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		fmt.Println("err", err)
		return
	}
	p, err := strconv.ParseFloat(u.ImdbRating, 32)
	p = (p / 10) * 100
	fmt.Printf("The movie : %s was released in %s - the IMBD rating is %d%% with %s votes\n", u.Title, u.Year, int(p), u.ImdbVotes)
}
