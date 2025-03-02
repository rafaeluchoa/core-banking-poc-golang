package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"nk/account/app"
)

const (
	URL = "http://localhost:8080"
)

var runSetup = true

func Setup() {
	if runSetup {
		runSetup = false
		app.Run("../../../")
	}
}

func Get[T any](uri string, req any) *T {
	reqJSON, err := json.Marshal(req)
	if err != nil {
		log.Panic(err)
	}

	var reqMap map[string]interface{}
	err = json.Unmarshal(reqJSON, &reqMap)
	if err != nil {
		log.Panic(err)
	}

	queryParams := url.Values{}
	for key, value := range reqMap {
		queryParams.Add(key, fmt.Sprintf("%v", value))
	}

	baseURL := URL + uri
	if len(queryParams) > 0 {
		baseURL += "?" + queryParams.Encode()
	}

	// #nosec G107: using const url base
	resp, err := http.Get(baseURL)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	return fromJSON[T](resp, uri, reqJSON)
}

func Post[T any](uri string, req any) *T {
	reqJSON, err := json.Marshal(req)
	if err != nil {
		log.Panic(err)
	}

	resp, err := http.Post(URL+uri, "application/json",
		bytes.NewBuffer([]byte(reqJSON)))
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	return fromJSON[T](resp, uri, reqJSON)
}

func fromJSON[T any](resp *http.Response, uri string, reqJSON []byte) *T {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	respJSON := string(body)
	log.Printf("%s\n>> %s\n<< %s\n", uri, reqJSON, respJSON)

	var result T
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Panic(err)
	}

	return &result
}
