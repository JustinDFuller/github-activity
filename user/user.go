package user

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
);

type GithubUserResponse struct {
	ReposUrl string `json:"repos_url"`
}

func (response *GithubUserResponse) parseJSON(jsonToParse []byte) {
	if !json.Valid(jsonToParse) {
		log.Fatal("Invalid JSON received");
	}

	err := json.Unmarshal(jsonToParse, &response);

	if err != nil {
		log.Fatal("An error occured while parsing JSON.");
	}
}

func (responseTarget *GithubUserResponse) httpRequest(url string) {
	res, err := http.Get(url);

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	
	if err != nil {
		log.Fatal(err)
	}

	responseTarget.parseJSON(body);
}

func (parsedJSON GithubUserResponse) prettyPrintJSON() {
	formattedJSON, err := json.MarshalIndent(parsedJSON, "", "  ")
	
	if err != nil {
			fmt.Println("error:", err)
	}
	
	fmt.Print(string(formattedJSON))
}

func FetchReposUrl() string {
	var githubApi GithubUserResponse;
	githubApi.httpRequest("https://api.github.com/users/JustinDfuller");
	return githubApi.ReposUrl;
}
