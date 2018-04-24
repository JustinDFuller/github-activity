package githubActivity

import (
	"github-activity/commitActivity"
	"github-activity/repos"
	"github-activity/types"
	"github-activity/user"
)

func Activity(userName types.User) repos.GithubReposResponse {
	repoUrl := user.FetchReposUrl(userName)
	repos := repos.FetchRepos(repoUrl, userName)
	return commitActivity.FetchCommitActivity(repos, userName)
}
