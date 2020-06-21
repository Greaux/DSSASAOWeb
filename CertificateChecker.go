package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type SSLTest []struct {
	DNSNames  []string  `json:"dns_names"`
	NotBefore time.Time `json:"not_before"`
	NotAfter  time.Time `json:"not_after"`
	Issuer    struct {
		Name         string `json:"name"`
		PubkeySha256 string `json:"pubkey_sha256"`
	} `json:"issuer"`
}

func CertificateChecker(url string) (SSLTest, error) {

	var SSLTest SSLTest
	CertificateInfo, err := http.Get("https://api.certspotter.com/v1/issuances?domain=" + url + "&expand=dns_names&expand=issuer&expand=cert")
	if err != nil {
		return SSLTest, err
	}
	defer CertificateInfo.Body.Close()

	json.NewDecoder(CertificateInfo.Body).Decode(&SSLTest)
	return SSLTest, nil

}

// CertificateCheckModule checks dnsNames and ssl times
func CertificateCheckModule(sslTests SSLTest, url string) {
	havDNSInURL := false
	notBefore := false
	notAfter := 0
	issuerName := false

	for _, sslTest := range sslTests {
		haveDNSInURL := false
		for _, dnsName := range sslTest.DNSNames {
			if strings.Contains(url, dnsName) {
				haveDNSInURL = true
				break
			}
		}

		if haveDNSInURL {
			havDNSInURL = true
		}

		if sslTest.NotBefore.Before(time.Now()) {
			notBefore = true
		}

		if sslTest.NotAfter.Add(time.Duration(30*24)*time.Hour*(-1)).Before(time.Now()) && sslTest.NotAfter.After(time.Now()) {
			notAfter = 2
		} else if sslTest.NotAfter.Add(time.Duration(30*24) * time.Hour * (-1)).After(time.Now()) {
			notAfter = 1
		}

		if strings.Contains(sslTest.Issuer.Name, "VMware Installer") || strings.Contains(sslTest.Issuer.Name, "Let's Encrypt") {

		} else {
			issuerName = true
		}
	}

	if havDNSInURL {
		WriteAnswerToData(url, "DNSNames OK")
		fmt.Println("DNSNames OK")
	} else {
		WriteAnswerToData(url, "DNSNames error")
		fmt.Println("DNSNames error")
	}

	if notBefore {
		WriteAnswerToData(url, "NotBefore OK")
		fmt.Println("NotBefore OK")
	} else {
		WriteAnswerToData(url, "NotBefore error")
		fmt.Println("NotBefore error")
	}

	if notAfter == 2 {
		WriteAnswerToData(url, "NotAfter Warning")
		fmt.Println("NotAfter Warning")
	} else if notAfter == 1 {
		WriteAnswerToData(url, "NotAfter OK")
		fmt.Println("NotAfter OK")
	} else if notAfter == 0 {
		WriteAnswerToData(url, "NotAfter error")
		fmt.Println("NotAfter error")
	}

	if !issuerName {
		WriteAnswerToData(url, "IssuerName Warning")
		fmt.Println("IssuerName Warning")
	} else {
		WriteAnswerToData(url, "IssuerName OK")
		fmt.Println("IssuerName OK")
	}
}
