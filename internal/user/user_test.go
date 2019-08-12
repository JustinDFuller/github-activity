package user

import (
	"fmt"
	"testing"

	"github.com/amekss/assert"
)

// This test is completely useless..
func TestActivity(t *testing.T) {
	data := GetActivity("JustinDFuller")

	assert.True(t, len(data) >= 1)

	for _, activity := range data {
		// assert.True(t);
		fmt.Println(activity)
	}
}
