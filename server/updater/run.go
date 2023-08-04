//go:build !gokrazy

package updater

import (
	"github.com/google/go-github/v32/github"
	"github.com/robotuimyhorobotuiotui/util"
	"github.com/thommyho/robotui/server"
)

// Run regularly checks version
func Run(log *util.Logger, httpd webServer, outChan chan<- util.Param) {
	u := &watch{
		log:     log,
		outChan: outChan,
		repo:    NewRepo(log, owner, repository),
	}

	c := make(chan *github.RepositoryRelease, 1)
	go u.watchReleases(server.Version, c) // endless

	for rel := range c {
		u.Send("availableVersion", *rel.TagName)
	}
}
