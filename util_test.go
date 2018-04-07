package goconfig

import (
	"testing"
	"io/ioutil"
)

func TestIsFile(t *testing.T) {
	name, err := ioutil.TempDir("", "goconfig")
	if err != nil {
		t.Errorf("unexpected error while creting temporary directory")
		t.Errorf(err.Error())
		t.FailNow()
	}
	if isFile(name) {
		t.Errorf("This should be a valid directory but it's a a file")
	}

	file, err := ioutil.TempFile(name, "goconfig")
	if err != nil {
		t.Errorf("unexpected error while creting temporary directory")
		t.FailNow()
	}
	if !isFile(file.Name()) {
		t.Errorf("This should be a valid directory but it's a a file")
	}

}
