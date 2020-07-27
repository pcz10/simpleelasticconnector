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

func (downloader *URLDownloader) Run(urlChannel <-chan string, returnChannel chan models.Note) {
	for {
		select {
		case url := <-urlChannel:
			note := GetNoteService().GetFromUrl(url)
			fmt.Printf("\nProcessing in: %v    Value in URL channel: %v    Recieved note in return channel: %v", downloader.GetId(), url, note)
			go getDataFromReturnChannel(returnChannel)
			returnChannel <- note
		default:
		}
	}
}

func getDataFromReturnChannel(ch chan models.Note) {
	fmt.Printf("\nRecieved note in return channel: %v ", <-ch)

}