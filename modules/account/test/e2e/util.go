package e2e

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"nk/account/app"

	"github.com/google/uuid"
)

const (
	URL = "http://localhost:8080"
)

func UUID() string {
	return uuid.Must(uuid.NewV7()).String()
}

func Setup() {
	app.Run("../../../")
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	respJson := string(body)
	log.Printf("POST %s\n>> %s\n<< %s\n", uri, reqJson, respJson)

	var result T
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Panic(err)
	}

	return &result
}
