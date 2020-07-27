package service

import (
	"fmt"
	"time"
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
			fmt.Printf("Processing in: %v", downloader.GetId())
			note := GetNoteService().GetFromUrl(<-urlChannel)
			fmt.Printf("Note: %v", note)
			returnChannel <- note
		default:
			fmt.Printf("URLDownloader - waiting for input")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
