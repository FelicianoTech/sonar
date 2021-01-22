package docker

import (
	"encoding/json"
	"net/http"
	"time"
)

type Tag struct {
	Name   string
	Size   uint64
	Date   time.Time
	Digest string
}

func GetAllTags(image string) ([]Tag, error) {

	var results []Tag

	reqURL := "https://hub.docker.com/v2/repositories/" + image + "/tags/?page_size=100"

	for {

		req, err := http.NewRequest("GET", reqURL, nil)
		if err != nil {
			return nil, err
		}

		resp, err := SendRequest(req, "", "")
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var respPage map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&respPage)

		for _, v := range respPage["results"].([]interface{}) {

			var aTag Tag

			aTag.Name = v.(map[string]interface{})["name"].(string)
			aTag.Size = uint64(v.(map[string]interface{})["full_size"].(float64))
			aTag.Date, err = time.Parse(time.RFC3339, v.(map[string]interface{})["last_updated"].(string))
			anImage := v.(map[string]interface{})["images"].([]interface{})[0]
			aTag.Digest = anImage.(map[string]interface{})["digest"].(string)
			if err != nil {
				return nil, err
			}

			results = append(results, aTag)
		}

		if respPage["next"] == nil {
			break
		} else {
			reqURL = respPage["next"].(string)
		}
	}

	return results, nil
}
