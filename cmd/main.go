package main

import (
	"encoding/json"
	"fmt"
	"gitpr/pkg/data"
	"gitpr/pkg/pullRequest"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	url := "https://api.github.com/search/issues?q=repo%3Abasecamp/local_time+state%3Aopen+state%3Aclosed+type%3Apr+sort%3Acreated+page%3A1"
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
		log.Printf("failed to get pr list %v", err)
		return
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Failed to read the response", err)
	}

	var pr pullRequest.AllPr
	json.Unmarshal(responseData, &pr)

	data.Show_pr_data(pr)
}
