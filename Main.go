package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var url string

	// Getting URL
	fmt.Print("Enter Link (w/o http): ")
	fmt.Fscan(os.Stdin, &url)
	fmt.Printf("Analyzing %s ...\n", url)

	// Creating Data file for domain
	CreateNewFiles(url) // finished

	// Getting https request
	resp, err := http.Get("http://" + url)
	if err != nil {
		panic(err)
	}

	// Save web to temp file
	f, err := os.Create("WebPage.html")
	resp.Write(f)
	f.Close()
	fmt.Println("html file created and wrote successfully")
	// save final URL (to check redirects)
	FinalURL := resp.Request.URL.String()

	// do this now so it won't be forgotten
	defer resp.Body.Close()

	// reads html as a slice of bytes
	if err != nil {
		panic(err)
	}

	// check forcing https
	HTTPS(FinalURL, url)

	// Whois check (via API)
	whois, err := DomainCheck(FinalURL, url)
	if err == nil {
		DomainCheckModule(whois, url)
	} else {
		fmt.Println("error getting domain info")
	}

	// Check Certificates (via API)
	sslTests, err := CertificateChecker(url)
	if err == nil {
		CertificateCheckModule(sslTests, url)
	}

	// Опросник
	MainQuestions(url)

	// Making analyze + save results
	FinalAnalyse(url)

	// Removing trash files
	DeleteDataFiles(url)
}
