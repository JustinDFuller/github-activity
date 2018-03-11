package main

import (
  "fmt"
	"github-activity/types"
	"github-activity/githubActivity"
)

func main() {
	activity := githubActivity.Activity(types.User("JustinDFuller"))
	fmt.Print(activity)
}