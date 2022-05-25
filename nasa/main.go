package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"nasa/models"
	"net/http"
)

func main() {
	url := "https://api.nasa.gov/planetary/apod?api_key=e2iQ0x7l5wDJz7bKZBkB5DL8DI5f6dSK9fFVGdNZ"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "bearer")

	client := &http.Client{}
	resp, _ := client.Do(req)
	
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	
	var Planetary models.Planetary
    json.Unmarshal(data, &Planetary)
	
	fmt.Printf("Date: %+v\n", Planetary.Date)
    fmt.Printf("Explanation: %+v\n", Planetary.Explanation)
	fmt.Printf("Hdurl: %+v\n", Planetary.Hdurl)
	fmt.Printf("MediaType: %+v\n", Planetary.MediaType)
	fmt.Printf("ServiceVersion: %+v\n", Planetary.ServiceVersion)
	fmt.Printf("Title: %+v\n", Planetary.Title)
	fmt.Printf("URL: %+v\n", Planetary.URL)
}