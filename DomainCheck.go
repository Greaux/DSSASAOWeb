package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Whois struct {
	WhoisRecord WhoisRecord `json:"WhoisRecord"`
}

type WhoisRecord struct {
	DomainName   string       `json:"domainName"`
	RegistryData RegistryData `json:"registryData"`
}

type RegistryData struct {
	CreatedDate time.Time   `json:"createdDate"`
	ExpiresDate time.Time   `json:"expiresDate"`
	DomainName  string      `json:"domainName"`
	NameServers NameServers `json:"nameServers"`
}

type NameServers struct {
	RawText   string   `json:"rawText"`
	HostNames []string `json:"hostNames"`
	Ips       []string `json:"ips"`
}

func DomainCheck(FinalURL string, url string) (Whois, error) {

	var whois Whois
	var apiKey = "at_fL3kKfXzqRceyaenMliViKdAlma8t"
	DomainInfo, err := http.Get("https://www.whoisxmlapi.com/whoisserver/WhoisService?outputFormat=JSON&apiKey=" + apiKey +"&domainName=" + FinalURL)
	if err != nil {
		return whois, err
	}
	defer DomainInfo.Body.Close()
	json.NewDecoder(DomainInfo.Body).Decode(&whois)

	timeTest(url)
	return whois, nil
}

func timeTest(url string) {
	registry := RegistryData{CreatedDate: time.Time{}, ExpiresDate: time.Time{}}

	fmt.Println(registry.CreatedDate)
	fmt.Println(registry.ExpiresDate)

	tenDays := time.Now().Add(time.Duration(10*24) * time.Hour)
	thirtyDays := time.Now().Add(time.Duration(30*24) * time.Hour)

	fmt.Println(tenDays)
	fmt.Println(thirtyDays)

	passed := registry.CreatedDate.Before(tenDays)
	left := registry.ExpiresDate.After(thirtyDays)

	if passed {
		// WriteAnswerToData(url, Answer1)
		fmt.Println("created > 10 days")
	}
	if left {
		// do something else
		fmt.Println("expires > 30 days")
	}
}

// DomainCheckModule checks dates and domain name
func DomainCheckModule(whois Whois, url string) {
	if whois.WhoisRecord.RegistryData.CreatedDate.After(time.Now().Add(time.Duration(10*24) * time.Hour * (-1))) {
		WriteAnswerToData(url, "CreatedDate error")
		fmt.Println("CreatedDate error")
	} else {
		WriteAnswerToData(url, "CreatedDate OK")
		fmt.Println("CreatedDate OK")
	}

	if whois.WhoisRecord.RegistryData.ExpiresDate.Before(time.Now().Add(time.Duration(30*24) * time.Hour * (-1))) {
		WriteAnswerToData(url, "ExpiresDate error")
		fmt.Println("ExpiresDate error")
	} else {
		WriteAnswerToData(url, "ExpiresDate OK")
		fmt.Println("ExpiresDate OK")
	}

	if strings.Contains(url, whois.WhoisRecord.DomainName) {
		WriteAnswerToData(url, "DomainName is OK")
		fmt.Println("DomainName is OK")
	} else {
		WriteAnswerToData(url, "DomainName is NOT K")
		fmt.Println("DomainName is NOT K")
	}
}
