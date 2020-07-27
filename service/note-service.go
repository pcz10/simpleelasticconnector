package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"todo/model"
	"todo/persistance"
)

type NoteService struct{}

var NoteServiceClient = GetNoteService()

func GetNoteService() *NoteService {
	return &NoteService{}
}

func (noteService *NoteService) FindAll() []byte {
	var notes = persistance.GetElasticClient().FindAll()
	var js []byte
	js, err := json.Marshal(notes)
	if err != nil {
		log.Println("Marshall error. Err= ", err)
	}
	return js
}

func (noteService *NoteService) Add(w http.ResponseWriter, r *http.Request) []byte {
	var note model.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	js, err := json.Marshal(note)
	if err != nil {
		log.Println("Marshall error. Err= ", err)
	}
	err = persistance.GetElasticClient().Add(js)
	return js

}

func (noteService *NoteService) GetFromUrl(url string) model.Note {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println("\nget from url URL: ", url)
	data, _ := ioutil.ReadAll(res.Body)
	var note model.Note
	err1 := json.Unmarshal(data, &note)
	if err1 != nil {
		log.Println("Unmarshall note error. Err=", err1)
	}
	// log.Println("\nget from url note: ", note)
	return note
}
