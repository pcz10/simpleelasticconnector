package server

import (
	"net/http"
	"log"
	"fmt"
	"todo/service"
)

func Run() {
	fmt.Printf("Starting server at port 8080\n")
	
	handleRequests()
	
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

func handleRequests() {
	http.HandleFunc("/", helloServer)
	http.HandleFunc("/get", getNote)
	http.HandleFunc("/add", addNote)
}

func helloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "welcome page!")
}

func getNote(w http.ResponseWriter, r *http.Request) {
	notes := service.FindAll()
	w.Header().Set("Content-Type", "application/json")
	w.Write(notes)
}

func addNote(w http.ResponseWriter, r *http.Request) {
	js := service.Add(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
