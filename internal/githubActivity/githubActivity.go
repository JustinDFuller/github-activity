package githubActivity

import (
	"github.com/JustinDFuller/github-activity/internal/commitActivity"
	"github.com/JustinDFuller/github-activity/internal/repos"
	"github.com/JustinDFuller/github-activity/internal/types"
	"github.com/JustinDFuller/github-activity/internal/user"
)

func Activity(userName types.User) repos.GithubReposResponse {
	repoUrl := user.FetchReposUrl(userName)
	repos := repos.FetchRepos(repoUrl, userName)
	return commitActivity.FetchCommitActivity(repos, userName)
}
