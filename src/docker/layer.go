package docker

import (
	"encoding/json"
	"net/http"
)

type Layer struct {
	Digest      string
	Size        uint64
	Instruction string
}

func GetAllLayers(image, tag string) ([]Layer, error) {

	var results []Layer

	reqURL := "https://hub.docker.com/v2/repositories/" + image + "/tags/" + tag + "/images"

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := SendRequest(req, "", "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respPage []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&respPage)
	if err != nil {
		return nil, err
	}

	for _, v := range respPage[0]["layers"].([]interface{}) {

		var aLayer Layer

		if v.(map[string]interface{})["digest"] != nil {
			aLayer.Digest = v.(map[string]interface{})["digest"].(string)
		}
		aLayer.Size = uint64(v.(map[string]interface{})["size"].(float64))
		aLayer.Instruction = v.(map[string]interface{})["instruction"].(string)

		results = append(results, aLayer)
	}

	return results, nil
}
