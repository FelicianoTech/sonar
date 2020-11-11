package metafile

import "time"

type packageInfo struct {
	name        string
	version     string
	packageType string
	source      string
}

type metafile struct {
	version     string
	publishDate time.Time
	commitHash  string
	filePath    string
	fileRepo    string
	packages    []packageInfo
}
