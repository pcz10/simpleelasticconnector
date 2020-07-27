package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	models "todo/model"
	services "todo/service"

	"github.com/gorilla/mux"
)

var informationChannel = make(chan string)

var downloaderWorker1 = services.URLDownloader{Id: "urlDownloader_1"}
var downloaderWorker2 = services.URLDownloader{Id: "urlDownloader_2"}
var downloaderWorker3 = services.URLDownloader{Id: "urlDownloader_3"}
var downloaderWorker4 = services.URLDownloader{Id: "urlDownloader_4"}
var downloaderWorker5 = services.URLDownloader{Id: "urlDownloader_5"}

func Run() {
	fmt.Printf("Starting server at port 8080\n")

	router := mux.NewRouter()
	router.HandleFunc("/", helloServer).Methods("GET")
	router.HandleFunc("/get", getNotes).Methods("GET")
	router.HandleFunc("/get/{id}", getNoteById).Methods("GET")
	router.HandleFunc("/add", addNote).Methods("POST")
	router.HandleFunc("/urls", getFromUrls).Methods("GET")

	returnChannel1 := make(chan models.Note)
	returnChannel2 := make(chan models.Note)
	returnChannel3 := make(chan models.Note)
	returnChannel4 := make(chan models.Note)
	returnChannel5 := make(chan models.Note)

	go downloaderWorker1.Run(informationChannel, returnChannel1)
	go downloaderWorker2.Run(informationChannel, returnChannel2)
	go downloaderWorker3.Run(informationChannel, returnChannel3)
	go downloaderWorker4.Run(informationChannel, returnChannel4)
	go downloaderWorker5.Run(informationChannel, returnChannel5)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
	
}

func getFromUrls(w http.ResponseWriter, r *http.Request) {
	urls := []string{"http://localhost:8080/get/1", "http://localhost:8080/get/2", "http://localhost:8080/get/3", "http://localhost:8080/get/4",
		"http://localhost:8080/get/5", "http://localhost:8080/get/6", "http://localhost:8080/get/7", "http://localhost:8080/get/8",
		"http://localhost:8080/get/9", "http://localhost:8080/get/10", "http://localhost:8080/get/11", "http://localhost:8080/get/12",
		"http://localhost:8080/get/11", "http://localhost:8080/get/21", "http://localhost:8080/get/31", "http://localhost:8080/get/41",
		"http://localhost:8080/get/51", "http://localhost:8080/get/61", "http://localhost:8080/get/71", "http://localhost:8080/get/81",
		"http://localhost:8080/get/91", "http://localhost:8080/get/101", "http://localhost:8080/get/111", "http://localhost:8080/get/121"}
	for _, url := range urls {
		informationChannel <- url
	}
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome page!")
}

func getNotes(w http.ResponseWriter, r *http.Request) {
	notes := services.GetNoteService().FindAll()
	w.Header().Set("Content-Type", "application/json")
	w.Write(notes)
}

func getNoteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["id"]
	id, err := strconv.Atoi(inputOrderID)
	if err != nil {
		fmt.Println("convert error ", err)
	}
	note := &models.Note{
		ID:     id,
		Task:   "task" + inputOrderID,
		Status: true,
	}
	var js []byte
	js, err1 := json.Marshal(note)
	if err1 != nil {
		log.Println("Marshall error. Err= ", err1)
	}
	w.Write(js)
}

func addNote(w http.ResponseWriter, r *http.Request) {
	js := services.GetNoteService().Add(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
