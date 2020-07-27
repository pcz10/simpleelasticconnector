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
var returnChannel = make(chan models.Note)

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

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
	

	go downloaderWorker1.Run(informationChannel, returnChannel)
	go downloaderWorker2.Run(informationChannel, returnChannel)
	go downloaderWorker3.Run(informationChannel, returnChannel)
	go downloaderWorker4.Run(informationChannel, returnChannel)
	go downloaderWorker5.Run(informationChannel, returnChannel)

	for {
		select {
		case note := <-returnChannel:
			fmt.Printf("Recieved note: %v", note)
		}
	}

}

func getFromUrls(w http.ResponseWriter, r *http.Request) {
	urls := []string{"http://localhost:8080/get/1", "http://localhost:8080/get/2", "http://localhost:8080/get/3", "http://localhost:8080/get/4",
		"http://localhost:8080/get/5", "http://localhost:8080/get/6", "http://localhost:8080/get/7", "http://localhost:8080/get/8",
		"http://localhost:8080/get/9", "http://localhost:8080/get/10", "http://localhost:8080/get/11", "http://localhost:8080/get/12"}
	for _, url := range urls {
		log.Printf("revieved URL %v from urls %v", url, urls)
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
