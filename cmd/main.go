package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AllPr struct {
	Items []Items
}

type Items struct {
	Title        string
	Url          string
	Pull_Request PullRequest `json: "pull_request"`
}

type PullRequest struct {
	Url      string `json: "url"`
	Html_Url string `json: "html_url"`
}

func main() {

	url := "https://api.github.com/search/issues?q=repo%3Abasecamp/local_time+state%3Aopen+type%3Apr+sort%3Acreated"
	fmt.Print("hello", url)
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
		log.Printf("failed to get pr list %v", err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(responseData))
	var pr AllPr
	json.Unmarshal(responseData, &pr)
	fmt.Printf("\n\n\n\n>> %+v", pr)
}
