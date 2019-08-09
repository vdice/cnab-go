package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestCheckDriverExists(t *testing.T) {
	name := "missing-driver"
	cmddriver := &Driver{Name: name}
	if cmddriver.CheckDriverExists() {
		t.Errorf("Expected driver %s not to exist", name)
	}

	name = "existing-driver"
	cmddriver = &Driver{Name: name}
	dirname, err := ioutil.TempDir("", "cnab")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dirname)
	filename := fmt.Sprintf("%s/cnab-%s", dirname, name)
	newfile, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}

	newfile.Chmod(0755)
	path := os.Getenv("PATH")
	pathlist := []string{dirname, path}
	newpath := strings.Join(pathlist, string(os.PathListSeparator))
	defer os.Setenv("PATH", path)
	os.Setenv("PATH", newpath)
	if !cmddriver.CheckDriverExists() {
		t.Fatalf("Expected driver %s to exist", name)
	}

}
