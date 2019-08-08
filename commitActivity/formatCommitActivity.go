package commitActivity

import (
	"strings"
)

import (
	"github-activity/types"
)

type FormattedCommitResponseAsMap map[string]uint8

func FormatCommitActivity(commits GithubCommitResponse) []types.FormattedCommitResponse {
	var formattedCommits []types.FormattedCommitResponse
	formattedCommitsMap := make(FormattedCommitResponseAsMap)

	for _, value := range commits {
		date := strings.Split(value.Commit.Committer.Date, "T")[0]
		formattedCommitsMap[date] += 1
	}

	for key, value := range formattedCommitsMap {
		formattedCommits = append(formattedCommits, types.FormattedCommitResponse{
			Date:    key,
			Commits: value,
		})
	}

	return formattedCommits
}
