package githubActivityJson

import (
	"encoding/json"
	"github.com/JustinDFuller/github-activity/internal/user"
)

func Fetch(username string) string {
	activity := user.GetActivity(username)
	stringified, _ := json.Marshal(activity)
	return string(stringified)
}
