package main

import (
	"log"
	"os"
)

//Function to write result depends on analyzed info
func WriteResultData(url string, whatYouNeedToWrite string) {
	// If the file doesn't exist, create it, or append to the file
	fileName := url + ".result"
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(whatYouNeedToWrite + "\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func WriteAnswerToData(url string, AnswerToFile string) {
	fileName := url + ".data"
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(AnswerToFile + "\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}