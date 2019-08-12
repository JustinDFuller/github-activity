package main

import (
	"fmt"
	"github.com/JustinDFuller/github-activity/internal/user"
)

func main() {
	activity := user.GetActivity("JustinDFuller")
	fmt.Println(activity)
}
