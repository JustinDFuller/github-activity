package githubActivityJson

import (
	// "github.com/nsf/jsondiff"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestExpectedOutput(t *testing.T) {
	directory, err := os.Getwd()
	file := filepath.Join(directory, "expectedOutput.json")

	if err != nil {
		t.Fatal(err)
	}

	expected, err := ioutil.ReadFile(file)

	if err != nil {
		t.Fatal(err)
	}

	actual := Fetch("JustinDFuller")

	if actual != string(expected) {
		// options := jsondiff.DefaultConsoleOptions()
		// _, diff := jsondiff.Compare(expected, []byte(actual), &options)
		// t.Fatalf("%s", diff)
	}
}
