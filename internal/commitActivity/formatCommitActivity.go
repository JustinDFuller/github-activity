package commitActivity

import (
	"sort"
	"strings"
	"time"
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

func parseDate(date string) time.Time {
	parsedDate, _ := time.Parse("2006-01-02", date)
	return parsedDate
}

func sortByCommitDate(commits []types.FormattedCommitResponse) []types.FormattedCommitResponse {
	sort.SliceStable(commits, func(i, j int) bool {
		dateOne := parseDate(commits[i].Date)
		dateTwo := parseDate(commits[j].Date)
		return dateOne.After(dateTwo)
	})
	return commits
}

func FormatCommitActivity(commits GithubCommitResponse) []types.FormattedCommitResponse {
	return sortByCommitDate(toArray(getCommitsPerDate(commits)))
}
