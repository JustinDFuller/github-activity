package main

import (
	// "fmt"
	"github/user"
	"github/repos"
);

func main() {
	repoUrl := user.FetchReposUrl();
	repos.FetchRepos(repoUrl);
}
