package cmd

import (
	"errors"
	"fmt"

	"github.com/felicianotech/sonar/sonar/docker"
)

func ByteCountBinary(b uint64) string {

	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}

func getImageRefs(images []string) ([]*docker.ImageRef, error) {

	var imgRefs []*docker.ImageRef

	if len(images) > 25 {

		return nil, errors.New("Error: image limit is 25.")
	}

	for _, image := range images {

		imgRef, err := docker.ParseImageRef(image)
		if err != nil {
			return nil, err
		}

		imgRefs = append(imgRefs, imgRef)
	}

	return imgRefs, nil
}
