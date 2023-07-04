package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	ID   string `json:"dizi_id"`
	Title string `json:"dizi_isim"`
}

func main() {
	url := "https://www.diziarsivi.com/json.php"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var articles []Article
	err = json.NewDecoder(response.Body).Decode(&articles)
	if err != nil {
		log.Fatal(err)
	}

	for _, article := range articles {
		fmt.Printf("%s\n", article.Title)
	}
}
