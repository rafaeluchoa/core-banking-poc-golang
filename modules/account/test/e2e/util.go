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
	reqJson, err := json.Marshal(req)
	if err != nil {
		log.Panic(err)
	}

	var reqMap map[string]interface{}
	err = json.Unmarshal(reqJson, &reqMap)
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

	resp, err := http.Get(baseURL)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	return fromJson[T](resp, uri, reqJson)
}

func Post[T any](uri string, req any) *T {
	reqJson, err := json.Marshal(req)
	if err != nil {
		log.Panic(err)
	}

	resp, err := http.Post(URL+uri, "application/json",
		bytes.NewBuffer([]byte(reqJson)))
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	return fromJson[T](resp, uri, reqJson)
}

func fromJson[T any](resp *http.Response, uri string, reqJson []byte) *T {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	respJson := string(body)
	log.Printf("%s\n>> %s\n<< %s\n", uri, reqJson, respJson)

	var result T
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Panic(err)
	}

	return &result
}
