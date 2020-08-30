package docker

import (
	"encoding/json"
	"net/http"
)

func ImagePulls(image string) (uint, error) {

	req, err := http.NewRequest("GET", "https://hub.docker.com/v2/repositories/"+image+"/", nil)
	if err != nil {
		return 0, err
	}

	resp, err := SendRequest(req, "", "")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result2 map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result2)

	return uint(result2["pull_count"].(float64)), nil
}

func ImageStars(image string) (uint, error) {

	req, err := http.NewRequest("GET", "https://hub.docker.com/v2/repositories/"+image+"/", nil)
	if err != nil {
		return 0, err
	}

	resp, err := SendRequest(req, "", "")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result2 map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result2)

	return uint(result2["star_count"].(float64)), nil
}

func ImageList(namespace string) ([]string, error) {

	req, err := http.NewRequest("GET", "https://hub.docker.com/v3/repositories/"+namespace+"/", nil)
	if err != nil {
		return nil, err
	}

	resp, err := SendRequest(req, "", "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result2 map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result2)

	var images []string

	for _, v := range result2["results"].([]interface{}) {
		images = append(images, v.(map[string]interface{})["name"].(string))
	}

	return images, nil
}
