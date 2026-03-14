package goconfig

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadWithOptionsPrecedence(t *testing.T) {
	dir, err := os.MkdirTemp("", "goconfig-load")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	file := filepath.Join(dir, "config.json")
	if err := os.WriteFile(file, []byte(`{"value":"file"}`), 0644); err != nil {
		t.Fatal(err)
	}

	lookup := func(name string) (string, bool) {
		if name == "VALUE" {
			return "env", true
		}
		return "", false
	}

	c := struct {
		Value string
	}{Value: "default"}

	err = Load(&c,
		WithProgramName("testbin"),
		WithArgs([]string{"-value", "arg"}),
		WithConfigFile(file),
		WithEnvLookup(lookup),
	)
	AssertNil(t, err)

	AssertEqual(t, c.Value, "arg")
}

func TestLoadWithoutImplicitConfigFile(t *testing.T) {
	dir, err := os.MkdirTemp("", "goconfig-no-implicit")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	if err := os.WriteFile(filepath.Join(dir, "config.json"), []byte(`{"value":"file"}`), 0644); err != nil {
		t.Fatal(err)
	}

	oldCwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(oldCwd)

	c := struct {
		Value string
	}{Value: "default"}

	err = Load(&c,
		WithArgs(nil),
		WithoutImplicitConfigFile(),
		WithEnvLookup(func(string) (string, bool) { return "", false }),
	)
	AssertNil(t, err)

	AssertEqual(t, c.Value, "default")
}

func TestLoadInvalidTarget(t *testing.T) {
	err := Load(struct{}{})
	AssertNotNil(t, err)
	AssertEqual(t, err.Error(), "config target must be a pointer")
}
