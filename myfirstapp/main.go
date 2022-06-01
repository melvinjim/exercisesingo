package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"myfirstapp/models"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("."))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	r.HandleFunc("/", Index).Methods(http.MethodGet)
	r.HandleFunc("/show-planetary", ShowPlanetary).Methods(http.MethodGet)
	r.HandleFunc("/planetary", Planetary).Methods("GET")
	r.HandleFunc("/profile", Profile).Methods(http.MethodGet)

	srv := http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	log.Println("listening...")
	srv.ListenAndServe()
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/index.html")
}

func ShowPlanetary(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/show_planetary.html")
}

func Planetary(w http.ResponseWriter, r *http.Request) {
	planetary := models.GetPlanetary()

	fmt.Printf("Date: %+v\n", planetary.Date)
	fmt.Printf("Explanation: %+v\n", planetary.Explanation)
	fmt.Printf("Hdurl: %+v\n", planetary.Hdurl)
	fmt.Printf("MediaType: %+v\n", planetary.MediaType)
	fmt.Printf("ServiceVersion: %+v\n", planetary.ServiceVersion)
	fmt.Printf("Title: %+v\n", planetary.Title)
	fmt.Printf("URL: %+v\n", planetary.URL)

	json.NewEncoder(w).Encode(planetary)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("./templates/profile.html")
	if err != nil {
		panic(err)
	}

	w.Write(file)
}