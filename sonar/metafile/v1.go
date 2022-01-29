package metafile

import (
	"log"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
)

type DockerInfo struct {
	Namespace string
	Name      string
	Tags      []string
}

type GitInfo struct {
	Remote   string
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
	Home        string
	Docs        string
	Docker      *DockerInfo
	Git         *GitInfo
	Packages    []PackageInfo
}

func Generate() Metafile {

	mf := Metafile{
		Version:   "0.1.0",
		Generator: "sonar",
		Date:      time.Now(),
	}

	// missing
	// description
	// home
	// docs
	// package

	repository, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{
		DetectDotGit: true,
	})
	if err != nil {
		log.Print("Couldn't open repo")
	}

	remote, err := repository.Remote("origin")

	remoteURL := strings.Split(strings.Split(remote.String(), "\n")[0], "\t")[1]
	remoteURL = strings.Split(remoteURL, ":")[1]
	org := strings.Split(remoteURL, "/")[0]
	repo := strings.Split(remoteURL, "/")[1]
	repo = repo[0 : len(repo)-12]
	ref, err := repository.Head()

	mf.Git = &GitInfo{
		Remote: "https://github.com/" + org + "/" + repo,
		Hash:   ref.Hash().String(),
	}

	log.Print("org: " + org)
	log.Print("repo: " + repo)

	return mf
}
