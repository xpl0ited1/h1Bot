package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	headers := map[string][]string{
		"Accept": []string{"application/json"},
	}

	req, err := http.NewRequest("GET", "https://api.hackerone.com/v1/hackers/me/reports", nil)
	req.Header = headers
	req.SetBasicAuth("username", "TOKEN")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(string(body))
}
