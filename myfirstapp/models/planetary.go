package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Planetary struct {
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	Hdurl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	URL            string `json:"url"`
}

func GetPlanetary() Planetary {
	url := "https://api.nasa.gov/planetary/apod?api_key=e2iQ0x7l5wDJz7bKZBkB5DL8DI5f6dSK9fFVGdNZ"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "bearer")

	client := &http.Client{}
	resp, _ := client.Do(req)

	data, _ := ioutil.ReadAll(resp.Body)

	var planetary Planetary
	json.Unmarshal(data, &planetary)

	return planetary
}