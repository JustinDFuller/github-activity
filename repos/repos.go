package repos

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"sort"
	"time"
);

type GithubReposResponse []struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Url string `json:"url"`
	Language string `json:"language"`
	UpdatedAt string `json:"updated_at"`
}

func (repos *GithubReposResponse) parseJSON(jsonToParse []byte) {
	if !json.Valid(jsonToParse) {
		log.Fatal("Invalid JSON received");
	}

	err := json.Unmarshal(jsonToParse, &repos);

	if err != nil {
		log.Fatal("An error occured while parsing JSON.");
	}
}

func (repos *GithubReposResponse) httpRequest(url string) {
	res, err := http.Get(url);

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	
	if err != nil {
		log.Fatal(err)
	}

	repos.parseJSON(body);
}

func (repos GithubReposResponse) prettyPrintJSON() {
	formattedJSON, err := json.MarshalIndent(repos, "", "  ")
	
	if err != nil {
			fmt.Println("error:", err)
	}
	
	fmt.Print(string(formattedJSON))
}

func parseSortByDate(date string) time.Time {
	parsedTime, _ := time.Parse(time.RFC3339, date);
	return parsedTime;
}

func (repos GithubReposResponse) sortByDate() {
	sort.Slice(repos, func(i, j int) bool {
		firstDate := parseSortByDate(repos[i].UpdatedAt);
		secondDate := parseSortByDate(repos[j].UpdatedAt);
		return firstDate.After(secondDate);
	});
}

func FetchRepos(url string) GithubReposResponse {
	var repos GithubReposResponse;
	fmt.Print(url);
	repos.httpRequest(url);
	repos.sortByDate();
	repos.prettyPrintJSON();
	return repos;
}
