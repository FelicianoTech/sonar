package docker

import (
	"time"
)

/*
 * Represents our lowest level object. An image is arch specific and may or may
 * not be under a Tag. What end users might typically consider to be an image
 * we instead would call an ImageRef or Image Reference.
 */
type Image struct {
	Arch       string    `json:"architecture"`
	Digest     string    `json:"digest"`
	LastPulled time.Time `json:"last_pulled"`
	LastPushed time.Time `json:"last_pushed"`
	OS         string    `json:"os"`
	Size       int64     `json:"size"`
	Status     string    `json:"status"`
}
