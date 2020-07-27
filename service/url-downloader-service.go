package service

import (
	"fmt"
	models "todo/model"
)

type URLDownloader struct {
	Id string
}

var URLDownloaderClient = GetURLDownloader()

func GetURLDownloader() *URLDownloader {
	return &URLDownloader{}
}

func (downloader URLDownloader) GetId() string {
	return downloader.Id
}

func (downloader *URLDownloader) Run(urlChannel chan string, returnChannel chan models.Note) {
	for {
		select {
		case <-urlChannel:
			fmt.Printf("\nProcessing in: %v", downloader.GetId())
			note := GetNoteService().GetFromUrl(<-urlChannel)
			fmt.Printf("\nNote: %v", note)
			returnChannel <- note
		default:
		}
	}
}