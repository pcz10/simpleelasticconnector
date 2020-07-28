package service

import (
	"fmt"
	models "todo/model"
)

type URLDownloader struct {
	Id string
}

type URLPerNoteDataModel struct {
	URL string
	ReturnChannel chan models.Note
}

var URLDownloaderClient = GetURLDownloader()

func GetURLDownloader() *URLDownloader {
	return &URLDownloader{}
}

func (downloader URLDownloader) GetId() string {
	return downloader.Id
}

func (downloader *URLDownloader) Run(channel chan URLPerNoteDataModel) {
	for {
		select {
		case ch := <-channel:
			note := GetNoteService().GetFromUrl(ch.URL)
			fmt.Printf("\nProcessing in: %v    Value in URL channel: %v    Recieved note in return channel: %v", downloader.GetId(), ch.URL, note)
			//go getDataFromReturnChannel(returnChannel)
			ch.ReturnChannel <- note
		default:
		}
	}
}

func getDataFromReturnChannel(ch chan models.Note) {
	fmt.Printf("\nRecieved note in return channel: %v ", <-ch)
}