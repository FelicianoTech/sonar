package cmd

import (
	"bytes"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var numRequests int
var token string

func sendRequest(request *http.Request, username, password string) (*http.Response, error) {

	client := &http.Client{}

	// If username and password are blank, skip all of these, we're not doing an authenticated call.
	if username == "" && password == "" {
		return client.Do(request)
	}

	// The Docker Hub API isn't well documented. It's not clear for how long can
	// an auth token be used. For now, it's assumed that it will be good for every 6,000 requests.
	if numRequests%7000 == 0 {

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

		json.NewDecoder(resp.Body).Decode(&result)

		token = result["token"]
		numRequests++
	}

	request.Header.Set("Authorization", "JWT "+token)

	return client.Do(request)
}
