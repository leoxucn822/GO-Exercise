package main

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	var urlBase = "https://10.58.171.217:5665"
	var apiUser = "admin"
	var apiPass = "admin"
	var currentTime = time.Now().Unix()
	var duration = 3600

	var requestBody = []byte(`{
		"author": "admin",
		"host": "vm001",
		"comment": "test",
		"all_services": true,
		"start_time": currentTime,
		"end_time": currentTime+int64(duration)
    }`)

	urlEndpoint := urlBase + "/v1/actions/schedule-downtime"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", urlEndpoint, bytes.NewBuffer(requestBody))

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-HTTP-Method-Override", "POST")

	req.SetBasicAuth(apiUser, apiPass)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal("Server error:", err)
		return
	}
	defer resp.Body.Close()

	log.Print("Response status:", resp.Status)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	if resp.StatusCode == http.StatusOK {
		log.Print("Result: " + bodyString)
	} else {
		log.Fatal(bodyString)
	}
}
