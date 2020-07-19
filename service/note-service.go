package service

import (
	"encoding/json"
	"todo/persistance"
	"todo/model"
	"log"
	"net/http"
)

func FindAll() []byte {
	var notes = persistance.GetElasticClient().FindAll()
	var js []byte
	js, err := json.Marshal(notes)
	if err != nil {
		log.Println("Marshall error. Err= ", err)
	}
	return js
}

func Add(w http.ResponseWriter, r *http.Request) []byte {
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
