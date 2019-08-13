package main

import (
	"fmt"
	"github.com/JustinDFuller/github-activity/pkg/githubActivityJson"
)

func main() {
	fmt.Println(githubActivityJson.Fetch("JustinDFuller"))
}
