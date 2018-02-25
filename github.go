package main

import (
	"github/commitActivity"
	"github/repos"
	"github/user"
)

func main() {
	repoUrl := user.FetchReposUrl("JustinDfuller")
	repos := repos.FetchRepos(repoUrl)
	commitActivity.FetchCommitActivity(repos)
}
