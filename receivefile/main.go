package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("."))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	r.HandleFunc("/", Index).Methods(http.MethodGet)
	r.HandleFunc("/receive-file", ReceiveFile).Methods(http.MethodPost)

	srv := http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	log.Println("charging...")
	srv.ListenAndServe()
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/index.html")
}

func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	_, fileInfo, err := r.FormFile("document")
	fmt.Println(err)
	if err != nil {
		panic(err)
	}

	fmt.Println(fileInfo.Filename)

	fmt.Fprintf(w, fileInfo.Filename)
}
