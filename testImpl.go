package main

import (
	"fmt"
	"github.com/JustinDFuller/github-activity/githubActivity"
	"github.com/JustinDFuller/github-activity/types"
)

func main() {
	activity := githubActivity.Activity(types.User("JustinDFuller"))
	fmt.Println(activity)
}
