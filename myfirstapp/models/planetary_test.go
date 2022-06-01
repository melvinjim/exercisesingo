package models_test

import (
	"encoding/json"
	"io/ioutil"
	"myfirstapp/models"
	"net/http"
	"testing"
)

func TestDate(t *testing.T) {
	url := "https://api.nasa.gov/planetary/apod?api_key=e2iQ0x7l5wDJz7bKZBkB5DL8DI5f6dSK9fFVGdNZ"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "bearer")

	client := &http.Client{}
	resp, _ := client.Do(req)

	data, _ := ioutil.ReadAll(resp.Body)

	var planetary models.Planetary
	json.Unmarshal(data, &planetary)

	result := planetary.Date
	expectedResult := "2022-06-01"
	if result != expectedResult {
		t.Errorf("Expected %s but result was %s", expectedResult, result)
	}
}

func TestExplanation(t *testing.T) {
	url := "https://api.nasa.gov/planetary/apod?api_key=e2iQ0x7l5wDJz7bKZBkB5DL8DI5f6dSK9fFVGdNZ"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "bearer")

	client := &http.Client{}
	resp, _ := client.Do(req)

	data, _ := ioutil.ReadAll(resp.Body)

	var planetary models.Planetary
	json.Unmarshal(data, &planetary)

	result := planetary.Explanation
	expectedResult := "It wasn't the storm of the century -- but it was a night to remember. Last night was the peak of the Tau Herculids meteor shower, a usually modest dribble of occasional meteors originating from the disintegrating Comet 73P/Schwassmann-Wachmann 3. This year, calculations showed that the Earth might be passing through a particularly dense stream of comet debris -- at best creating a storm of bright meteors streaking out from the constellation of Hercules. What actually happened fell short of a meteor storm, but could be called a decent meteor shower. Featured here is a composite image taken at Kitt Peak National Observatory in Arizona, USA accumulated over 2.5 hours very late on May 30.  Over that time, 19 Tau Herculids meteors were captured, along with 4 unrelated meteors. (Can you find them?)  In the near foreground is the Bok 2.3-meter Telescope with the 4.0-meter Mayall Telescope just behind it. Next year, the annual Tau Herculids are expected to return to its normal low rate, with the next active night forecast for 2049."
	if result != expectedResult {
		t.Errorf("Expected %s but result was %s", expectedResult, result)
	}
}

// fmt.Println(string(data))

func TestHdurl(t *testing.T) {
	url := "https://api.nasa.gov/planetary/apod?api_key=e2iQ0x7l5wDJz7bKZBkB5DL8DI5f6dSK9fFVGdNZ"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "bearer")

	client := &http.Client{}
	resp, _ := client.Do(req)

	data, _ := ioutil.ReadAll(resp.Body)

	var planetary models.Planetary
	json.Unmarshal(data, &planetary)

	result := planetary.Hdurl
	expectedResult := "https://apod.nasa.gov/apod/image/2206/TauHerMeteors_Lyu_1280.jpg"
	if result != expectedResult {
		t.Errorf("Expected %s but result was %s", expectedResult, result)
	}
}

func TestMediaType(t *testing.T) {
	url := "https://api.nasa.gov/planetary/apod?api_key=e2iQ0x7l5wDJz7bKZBkB5DL8DI5f6dSK9fFVGdNZ"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "bearer")

	client := &http.Client{}
	resp, _ := client.Do(req)

	data, _ := ioutil.ReadAll(resp.Body)

	var planetary models.Planetary
	json.Unmarshal(data, &planetary)

	result := planetary.MediaType
	expectedResult := "image"
	if result != expectedResult {
		t.Errorf("Expected %s but result was %s", expectedResult, result)
	}
}

func TestServiceVersion(t *testing.T) {
	url := "https://api.nasa.gov/planetary/apod?api_key=e2iQ0x7l5wDJz7bKZBkB5DL8DI5f6dSK9fFVGdNZ"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "bearer")

	client := &http.Client{}
	resp, _ := client.Do(req)

	data, _ := ioutil.ReadAll(resp.Body)

	var planetary models.Planetary
	json.Unmarshal(data, &planetary)

	result := planetary.ServiceVersion
	expectedResult := "v1"
	if result != expectedResult {
		t.Errorf("Expected %s but result was %s", expectedResult, result)
	}
}

func TestTitle(t *testing.T) {
	url := "https://api.nasa.gov/planetary/apod?api_key=e2iQ0x7l5wDJz7bKZBkB5DL8DI5f6dSK9fFVGdNZ"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "bearer")

	client := &http.Client{}
	resp, _ := client.Do(req)

	data, _ := ioutil.ReadAll(resp.Body)

	var planetary models.Planetary
	json.Unmarshal(data, &planetary)

	result := planetary.Title
	expectedResult := "Tau Herculids Meteors over Kitt Peak Telescopes"
	if result != expectedResult {
		t.Errorf("Expected %s but result was %s", expectedResult, result)
	}
}

func TestURl(t *testing.T) {
	url := "https://api.nasa.gov/planetary/apod?api_key=e2iQ0x7l5wDJz7bKZBkB5DL8DI5f6dSK9fFVGdNZ"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "bearer")

	client := &http.Client{}
	resp, _ := client.Do(req)

	data, _ := ioutil.ReadAll(resp.Body)

	var planetary models.Planetary
	json.Unmarshal(data, &planetary)

	result := planetary.URL
	expectedResult := "https://apod.nasa.gov/apod/image/2206/TauHerMeteors_Lyu_1080.jpg"
	if result != expectedResult {
		t.Errorf("Expected %s but result was %s", expectedResult, result)
	}
}
