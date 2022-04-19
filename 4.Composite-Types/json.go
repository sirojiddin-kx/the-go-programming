package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {

	var movies = []Movie{
		{Title: "Osmondagi bolalar", Year: 2000, Color: true,
			Actors: []string{"Sirojiddin Khomidov", "Aziz Karimov"}},
		{Title: "Yerdagi bolalar", Year: 2022, Color: false, 
	        Actors: []string{"Sirojiddin Odamboev", "Jahongir Abdullayev"}},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed : %s", err)
	}

	fmt.Printf("%s\n", data)

	data, err = json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshling failed : %s", err)
	}

	fmt.Printf("%s\n", data)

	var titles []struct {Title string}
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshling failed: %s", err)
	}

	fmt.Println(titles)

}
