package types

// Assumes no more than 255 commits in a given day..
type FormattedCommitResponse struct {
	Date string
	Commits uint8
}

type GithubCommitResponse []struct {
	Commit struct {
		Committer struct {
			Date string `json:"date"`
		} `json:"committer"`
	} `json:"commit"`
}

type User string