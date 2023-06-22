package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Show struct {
	ID   string `json:"dizi_id"`
	Name string `json:"dizi_isim"`
}

func main() {
	url := "https://www.diziarsivi.com/json.php"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var shows []Show
	err = json.NewDecoder(response.Body).Decode(&shows)
	if err != nil {
		log.Fatal(err)
	}

	for _, show := range shows {
		fmt.Printf("ID: %s\nName: %s\n\n", show.ID, show.Name)
	}
}
