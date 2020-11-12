package docker

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Labels map[string]string

func getLabels(image, tag string) (Labels, error) {

	var results Labels

	reqURL := "https://registry-1.docker.io/v2/" + image + "/manifests/" + tag

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := SendRequest2(req, image)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respPage map[string]map[string]string
	json.NewDecoder(resp.Body).Decode(&respPage)

	reqURL = "https://registry-1.docker.io/v2/" + image + "/blobs/" + respPage["config"]["digest"]

	req, err = http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err = SendRequest2(req, image)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respPage2 map[string]map[string]Labels
	json.NewDecoder(resp.Body).Decode(&respPage2)

	results = respPage2["config"]["Labels"]

	return results, nil
}

func GetLabel(image, key string) (string, error) {

	var tag string

	imageParts := strings.Split(image, ":")

	image = imageParts[0]

	if len(imageParts) > 1 {
		tag = imageParts[1]
	} else {
		tag = "latest"
	}

	labels, err := getLabels(image, tag)
	if err != nil {
		return "", err
	}

	return labels[key], nil
}

func GetLabels(image string) (Labels, error) {

	var tag string

	imageParts := strings.Split(image, ":")

	image = imageParts[0]

	if len(imageParts) > 1 {
		tag = imageParts[1]
	} else {
		tag = "latest"
	}

	return getLabels(image, tag)
}
