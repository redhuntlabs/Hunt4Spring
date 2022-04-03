package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func heuristicCheck(url string) bool {
	log.Println("Checking:", url)
	targeturl := url + "?class.module.abcd.URLs%5B0%5D=0"
	resp, err := http.Get(targeturl)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == 400 {
		return false
	} else if heuristicScan(url) {
		return true
	}

	return false
}

func heuristicScan(url string) bool {
	targeturl := url + "?class.module.classLoader.URLs%5B0%5D=0"
	resp, err := http.Get(targeturl)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == 400 {
		log.Println(url, "[Seems to be vulnerable!]")
		return true
	}
	return false
}

func serializeJSON(fname string) bool {
	data, err := json.MarshalIndent(finalData, "", "  ")
	if err != nil {
		return false
	}
	ioutil.WriteFile(fname, data, 0644)
	return true
}
