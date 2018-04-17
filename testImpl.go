package main

import (
	"fmt"
	"github-activity/githubActivity"
	"github-activity/types"
)

func main() {
	activity := githubActivity.Activity(types.User("JustinDFuller"))
	fmt.Print(activity)
}
