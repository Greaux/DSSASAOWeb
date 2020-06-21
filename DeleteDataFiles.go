package main

import (
	"fmt"
	"os"
)

func DeleteDataFiles(url string) {
	os.Remove(url + ".data")
	fmt.Println("Data file removed successfully")
	os.Remove("WebPage.html")
	fmt.Println("html file removed successfully")
}
