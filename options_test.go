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

func TestLoadMultipleConfigFiles(t *testing.T) {
	dir, err := os.MkdirTemp("", "goconfig-multiple-config")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	base := filepath.Join(dir, "base.json")
	if err := os.WriteFile(base, []byte(`{"value":"base","keep":"base","nested":{"name":"base","port":8080}}`), 0644); err != nil {
		t.Fatal(err)
	}

	override := filepath.Join(dir, "override.json")
	if err := os.WriteFile(override, []byte(`{"value":"override","nested":{"port":9090}}`), 0644); err != nil {
		t.Fatal(err)
	}

	c := struct {
		Value  string
		Keep   string
		Nested struct {
			Name string
			Port int
		}
	}{}

	err = Load(&c,
		WithArgs([]string{"--config", base + "," + override}),
		WithEnvLookup(func(string) (string, bool) { return "", false }),
		WithoutImplicitConfigFile(),
	)
	AssertNil(t, err)

	AssertEqual(t, c.Value, "override")
	AssertEqual(t, c.Keep, "base")
	AssertEqual(t, c.Nested.Name, "base")
	AssertEqual(t, c.Nested.Port, 9090)
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
