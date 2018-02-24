package main

import (
	"github/user"
	"github/repos"
	"github/commitActivity"
);

func main() {
	repoUrl := user.FetchReposUrl("JustinDfuller");
	repos := repos.FetchRepos(repoUrl);
	commitActivity.FetchCommitActivity(repos);
}
