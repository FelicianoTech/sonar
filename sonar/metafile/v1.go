package metafile

import (
	"net/url"
	"time"
)

type DockerInfo struct {
	Namespace string
	Name      string
	Tags      []string
}

type GitInfo struct {
	Remote   url.URL
	Hash     string
	Filepath string
}

type PackageInfo struct {
	Name        string
	Version     string
	PackageType string
	Source      string
}

type Metafile struct {
	Version     string
	Generator   string
	Date        time.Time
	Description string
	Home        url.URL
	Docs        url.URL
	Docker      DockerInfo
	Git         GitInfo
	Packages    []PackageInfo
}
