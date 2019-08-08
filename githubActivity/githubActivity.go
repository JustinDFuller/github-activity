package githubActivity

import (
	"github.com/JustinDFuller/github-activity/commitActivity"
	"github.com/JustinDFuller/github-activity/repos"
	"github.com/JustinDFuller/github-activity/types"
	"github.com/JustinDFuller/github-activity/user"
)

func Activity(userName types.User) repos.GithubReposResponse {
	repoUrl := user.FetchReposUrl(userName)
	repos := repos.FetchRepos(repoUrl, userName)
	return commitActivity.FetchCommitActivity(repos, userName)
}
