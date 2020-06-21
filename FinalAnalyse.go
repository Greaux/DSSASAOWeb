package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func FinalAnalyse(url string) {
	file, err := os.Open(url + ".data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var points = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		switch str := scanner.Text(); str {
		case "Type=Forum":
			WriteResultData(url, "Тип - форум")
			switch str := scanner.Text(); str {
			case "Filter=0":
				WriteResultData(url, "Плоха")
				points = +5
			case "Filter=1":
				WriteResultData(url, "Наличие фильтра устраняет большую часть")
			}
		case "Type=Cloud":
			WriteResultData(url, "Тип - облако")

		case "Type=Bank":
			WriteResultData(url, "Тип - банк")

		case "Type=SW":
			WriteResultData(url, "Тип - социальная сеть")

		}
	var fpoints = strconv.Itoa(points)
	WriteResultData(url, "Количество очков" + fpoints )
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}