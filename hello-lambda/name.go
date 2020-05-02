package main

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/gojektech/heimdall/httpclient"
	log "github.com/sirupsen/logrus"
)

type nameFakeBodyResponse struct {
	Name string `json:"name"`
}

func randomName() string {
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(time.Second * 10))
	resp, err := client.Get("https://api.namefake.com", nil)
	if err != nil {
		log.Errorf("Failed to find a name: %s", err)
		return "Donald Trump"
	}
	body, _ := ioutil.ReadAll(resp.Body)

	response := nameFakeBodyResponse{}
	json.Unmarshal(body, &response)
	log.Infof("Fake name: %s", response.Name)
	return response.Name
}
