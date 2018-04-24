package githubActivity

import (
	"fmt"
	"testing"

	"github.com/amekss/assert"
)

func TestActivity(t *testing.T) {
	data := Activity("JustinDFuller")

	assert.True(t, len(data) >= 1)

	for _, activity := range data {
		// assert.True(t);
		fmt.Println(activity)
	}
}
