package main

import (
	"fmt"
)

type Movie struct {
	Title  string
	Year   int
	Color  bool
	Actors []string
}

func main() {
	fmt.Println("vim-go")
	var movies = []Movie{
		{
			Title: "Casablanca", Year: 1942, Color: false, Actors: []string{
				"Humphrey Bogart", "Ingrid Bergman",
			},
		},
		{
			Title: "Casablanca", Year: 1967, Color: true, Actors: []string{
				"Paul Newman",
			},
		},
	}
	fmt.Println(movies)

	var movie1 Movie
	movie1.Year = 1996

	color = &movie1.Color
	*color = true
}
