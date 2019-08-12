package commitActivity

import (
	"strings"
)

import (
	"github.com/JustinDFuller/github-activity/internal/types"
)

type commitsPerDateMap map[string]uint8

func getCommitsPerDate(commits GithubCommitResponse) commitsPerDateMap {
	commitsPerDate := make(commitsPerDateMap)

	for _, value := range commits {
		date := strings.Split(value.Commit.Committer.Date, "T")[0]
		commitsPerDate[date] += 1
	}

	return commitsPerDate
}

func toArray(commitsPerDate commitsPerDateMap) []types.FormattedCommitResponse {
	var formattedCommits []types.FormattedCommitResponse

	for key, value := range commitsPerDate {
		formattedCommits = append(formattedCommits, types.FormattedCommitResponse{
			Date:    key,
			Commits: value,
		})
	}

	return formattedCommits

}

func FormatCommitActivity(commits GithubCommitResponse) []types.FormattedCommitResponse {
	return toArray(getCommitsPerDate(commits))
}
