package cmd

import (
	"bytes"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func sendRequest(request *http.Request, username, password string) (*http.Response, error) {

	requestBody, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		log.Fatal(err)
	}

	// authenticate
	resp, err := http.Post("https://hub.docker.com/v2/users/login/", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var result map[string]string
	var token string

	json.NewDecoder(resp.Body).Decode(&result)

	token = result["token"]

	client := &http.Client{}

	request.Header.Set("Authorization", "JWT "+token)

	return client.Do(request)
}
