package cmd

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type dockerTag struct {
	name string
	size uint64
	date time.Time
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
