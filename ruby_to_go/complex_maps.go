package main

import (
	"encoding/json"
	"fmt"
)

type animal struct {
	Name string
	Year int
}

type author struct {
	Name    string `json: "name"`
	Country string `json: "country"`
}

type game struct {
	Title       string `json: "title"`
	ReleaseYear int    `json: "release_year"`
	Author      author
}

type footballTeams struct {
	Name string
}

func main() {

	var fruits = make(map[string]int)
	fruits["peras"] = 5
	fruits["manzanas"] = 15
	fmt.Println(fruits)

	characters := map[string]string{
		"ryu": "SF",
		"kyo": "KOF",
	}
	fmt.Println(characters)

	countries := map[string]string{}
	countries["chile"] = "curico"
	countries["argentina"] = "mendoza"

	fmt.Println(countries)

	animals := map[string]animal{}
	animals["carnivoros"] = animal{Name: "Dog", Year: 2}
	animals["hervivoros"] = animal{Name: "Planta", Year: 2}
	animals["reptiles"] = animal{Name: "Cocodrilo", Year: 5}
	fmt.Println(animals)

	teams := map[string]footballTeams{
		"España": {
			Name: "Real Madrid",
		},
		"Chile": {
			Name: "Universidad de chile",
		},
		"Argentina": {
			Name: "Boca Juniors",
		},
		"Brasil": {
			Name: "Sao Paulo",
		},
	}

	for key, val := range teams {
		fmt.Printf("Football Country: %s, Team Name:  %s\n", key, val.Name)
	}

	mapwitharr := map[string][]string{
		"fruits": []string{
			"apple", "mango", "avocado",
		},
		"veggies": []string{
			"carrot", "cucumber", "kale",
		},
	}

	for key, val := range mapwitharr {
		fmt.Printf("key %s, val %s\n", key, val)
	}

	type Data struct{ Games map[string][]game }
	games := map[string][]game{
		"Fighting": []game{
			game{Title: "KOF", ReleaseYear: 1932, Author: author{Name: "SNK", Country: "Japan"}},
			game{Title: "SF", ReleaseYear: 1912, Author: author{Name: "CAPCOM", Country: "Japan"}},
		},
		"Terror": []game{
			game{Title: "resident evil", ReleaseYear: 2001, Author: author{Name: "Capcom", Country: "USA"}},
		},
		"Adventure": []game{
			game{Title: "Sonic Adventure 2", ReleaseYear: 2000, Author: author{Name: "SEGA", Country: "JAPAN"}},
		},
		"Sports": []game{
			game{Title: "Sega Sport Tennis", ReleaseYear: 1998, Author: author{Name: "SEGA", Country: "JAPAN"}},
			game{Title: "FIFA", ReleaseYear: 1998, Author: author{Name: "EA", Country: "USA"}},
		},
	}

	data := Data{Games: games}
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(string(bytes))

	fmt.Println("\n*****GAMES*****")
	for key, game := range games {
		fmt.Printf("\nCategory %s\n", key)
		for _, val := range game {
			fmt.Printf("Game Name: %s, ReleaseYear: %d, Author Name: %s, Author Country: %s\n", val.Title, val.ReleaseYear, val.Author.Name, val.Author.Country)
		}
	}

	letters := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	if exists := letters["z"] != 0; exists == false {
		fmt.Println("letter Z doesn't exists")
	}

	fmt.Println("before delete", letters)
	delete(letters, "d")
	fmt.Println("after delete", letters)
}
