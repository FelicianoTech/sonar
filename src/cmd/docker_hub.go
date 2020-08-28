package cmd

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type dockerLayer struct {
	digest      string
	size        uint64
	instruction string
}

type dockerTag struct {
	name string
	size uint64
	date time.Time
}

func getAllLayers(image, tag string) []dockerLayer {

	var results []dockerLayer

	reqURL := "https://hub.docker.com/v2/repositories/" + image + "/tags/" + tag + "/images"

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := sendRequest(req, "", "")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var respPage []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&respPage)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range respPage[0]["layers"].([]interface{}) {

		var aLayer dockerLayer

		if v.(map[string]interface{})["digest"] != nil {
			aLayer.digest = v.(map[string]interface{})["digest"].(string)
		}
		aLayer.size = uint64(v.(map[string]interface{})["size"].(float64))
		aLayer.instruction = v.(map[string]interface{})["instruction"].(string)

		results = append(results, aLayer)
	}

	return results
}

func getAllTags(image string) []dockerTag {

	var results []dockerTag

	reqURL := "https://hub.docker.com/v2/repositories/" + image + "/tags/?page_size=100"

	for {

		req, err := http.NewRequest("GET", reqURL, nil)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := sendRequest(req, "", "")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		var respPage map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&respPage)

		for _, v := range respPage["results"].([]interface{}) {

			var aTag dockerTag

			aTag.name = v.(map[string]interface{})["name"].(string)
			aTag.size = uint64(v.(map[string]interface{})["full_size"].(float64))
			aTag.date, err = time.Parse(time.RFC3339, v.(map[string]interface{})["last_updated"].(string))
			if err != nil {
				log.Fatal(err)
			}

			results = append(results, aTag)
		}

		if respPage["next"] == nil {
			break
		} else {
			reqURL = respPage["next"].(string)
		}
	}

	return results
}
