package main

import (
	"fmt"
	"github.com/JustinDFuller/github-activity/internal/githubActivity"
	"github.com/JustinDFuller/github-activity/internal/types"
)

func main() {
	activity := githubActivity.Activity(types.User("JustinDFuller"))
	fmt.Println(activity)
}
