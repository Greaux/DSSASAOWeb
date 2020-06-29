package main

import (
	"strings"
)

func HTTPS(FinalURL string, url string) (HttpsResult string){
	if strings.Contains(FinalURL,"https") != false {
		HttpsResult = "https check positive"
		WriteAnswerToData(url,"HTTPS=1")
	}
	if strings.Contains(FinalURL,"https") != true {
		HttpsResult = "https check negative"
		WriteAnswerToData(url,"HTTPS=0")
	}
	return
}