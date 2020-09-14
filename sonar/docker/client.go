package docker

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

var numRequests int
var token string

func SendRequest(request *http.Request, username, password string) (*http.Response, error) {

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
			return nil, err
		}

		// authenticate
		resp, err := http.Post("https://hub.docker.com/v2/users/login/", "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var result map[string]string

		json.NewDecoder(resp.Body).Decode(&result)

		token = result["token"]
		numRequests++
	}

	request.Header.Set("Authorization", "JWT "+token)

	resp, err := client.Do(request)

	// Docker Hub has a huge problem with their API right now with backends
	// simply ignoring/dropping requests due to a rate limit that isn' exposed
	// to the end user. Until this is fixed or more information is given, we're
	// going to impose a hard 2 second pause. This will kill the speed of Sonar
	// but will make sure that users' requests, particularly deletion requests,
	// will actually work instead of returning false positives.
	time.Sleep(time.Second * 2)

	return resp, err
}
