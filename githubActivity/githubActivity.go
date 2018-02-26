package githubActivity

import (
	"github/commitActivity"
	"github/repos"
	"github/user"
	"github/types"
)

func Activity(userName types.User) repos.GithubReposResponse {
	repoUrl := user.FetchReposUrl(userName)
	repos := repos.FetchRepos(repoUrl, userName)
	return commitActivity.FetchCommitActivity(repos)
}
