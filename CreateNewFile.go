package main

import (
	"fmt"
	"os"
)

func CreateNewFiles(url string) {
	f1, err := os.Create(url + ".data")
	if err != nil {
		fmt.Println(err)
		return
	}
	f1.Close()
	fmt.Println("data file created successfully")

	f2, err := os.Create(url + ".result")
	if err != nil {
		fmt.Println(err)
		return
	}
	f2.Write([]byte("Результат для " + url + "\n"))
	f2.Close()
	fmt.Println("result file started successfully")
}
