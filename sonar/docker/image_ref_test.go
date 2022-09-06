package docker

import (
	"reflect"
	"testing"
)

func TestNewImageRef(t *testing.T) {

	testCases := []struct {
		input    []string
		expected *ImageRef
	}{
		{
			[]string{"hubci", "gotham", "0.10.1"},
			&ImageRef{
				Namespace: "hubci",
				Name:      "gotham",
				Tag:       "0.10.1",
				ShowTag:   true,
			},
		},
		{
			[]string{"goreleaser", "goreleaser", ""},
			&ImageRef{
				Namespace: "goreleaser",
				Name:      "goreleaser",
				Tag:       "latest",
				ShowTag:   true,
			},
		},
		{
			[]string{"", "ubuntu", "20.04"},
			&ImageRef{
				Namespace: "library",
				Name:      "ubuntu",
				Tag:       "20.04",
				ShowTag:   true,
			},
		},
	}

	for i, tc := range testCases {

		actual, _ := NewImageRef(tc.input[0], tc.input[1], tc.input[2])

		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("NewImageRef test[%d]: expected %v, actual %v", i, tc.expected, actual)
		}
	}
}

func TestString(t *testing.T) {

	testCases := []struct {
		input    string
		expected string
	}{
		{
			"hubci/gotham:0.10.1",
			"hubci/gotham:0.10.1",
		},
		{
			"goreleaser/goreleaser",
			"goreleaser/goreleaser:latest",
		},
		{
			"ubuntu:20.04",
			"library/ubuntu:20.04",
		},
	}

	for i, tc := range testCases {

		image, _ := ParseImageRef(tc.input)
		actual := image.String()

		if actual != tc.expected {
			t.Errorf("String test[%d]: expected %v, actual %v", i, tc.expected, actual)
		}
	}
}
