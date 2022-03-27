package docker

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"time"
)

type Tag struct {
	Name   string
	Size   int64
	Date   time.Time
	Digest string
}

func GetAllTags(imageStr string) ([]Tag, error) {

	var results []Tag
	var image *ImageRef

	image, err := ParseImageRef(imageStr)
	if err != nil {
		return nil, errors.New("Image name not parsable.")
	}
	image.ShowTag = false

	reqURL := "https://hub.docker.com/v2/repositories/" + image.String() + "/tags/?page_size=100"

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

			// if a tag is inactive, skip
			if v.(map[string]interface{})["tag_status"].(string) == "inactive" {
				continue
			}

			var aTag Tag

			aTag.Name = v.(map[string]interface{})["name"].(string)
			aTag.Size = int64(v.(map[string]interface{})["full_size"].(float64))
			aTag.Date, err = time.Parse(time.RFC3339, v.(map[string]interface{})["last_updated"].(string))
			anImage := v.(map[string]interface{})["images"].([]interface{})[0]

			// There are cases where the digest for a tag can be missing. Not
			// sure why. Until then, check for this edge case and set to an
			// empty string (don't set) when digest isn't available.
			if anImage.(map[string]interface{})["digest"] != nil {
				aTag.Digest = anImage.(map[string]interface{})["digest"].(string)
			}

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

func GetFilteredTags(imageStr string, nameRegex string) ([]Tag, error) {

	allTags, err := GetAllTags(imageStr)
	if err != nil {
		return nil, err
	}

	var filteredTags []Tag
	var invert bool

	if nameRegex[0] == '!' {
		invert = true
		nameRegex = nameRegex[1:]
	}

	for _, tag := range allTags {

		matched, err := regexp.MatchString(nameRegex, tag.Name)
		if err != nil {
			return nil, err
		}

		if matched != invert {
			filteredTags = append(filteredTags, tag)
		}
	}

	return filteredTags, nil
}
