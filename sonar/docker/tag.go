package docker

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"time"
)

/*
 * A Tag is a group of one or more Images (multiple archs) in an ImageFamily.
 * Typically they are published around the same time with the CPU architecture
 * being the main difference between the Images.
 */
type Tag struct {
	Date       time.Time
	Digest     string    `json:"digest"`
	Images     []Image   `json:"images"`
	LastPulled time.Time `json:"tag_last_pulled"`
	LastPushed time.Time `json:"tag_last_pushed"`
	Name       string    `json:"name"`
	Size       int64     `json:"full_size"`
	Status     string    `json:"tag_status"`
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

			// only parse time when present
			if v.(map[string]interface{})["tag_last_pushed"] != nil {
				aTag.LastPushed, _ = time.Parse(time.RFC3339, v.(map[string]interface{})["tag_last_pushed"].(string))
			}

			aTag.Name = v.(map[string]interface{})["name"].(string)
			aTag.Size = int64(v.(map[string]interface{})["full_size"].(float64))
			aTag.Date = aTag.LastPushed
			if err != nil {
				return nil, err
			}

			for _, i := range v.(map[string]interface{})["images"].([]interface{}) {

				var anImage Image

				//anImage.LastPushed, err = time.Parse(time.RFC3339, i.(map[string]interface{})["last_pushed"].(string))
				//if err != nil {
				//	return nil, err
				//}

				anImage.Arch = i.(map[string]interface{})["architecture"].(string)

				if i.(map[string]interface{})["digest"] != nil {
					anImage.Digest = i.(map[string]interface{})["digest"].(string)
				}

				aTag.Images = append(aTag.Images, anImage)
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
