package goconfig

import (
	"io/ioutil"
	"testing"
	"time"
)

func TestFillJsonNestedDuration(t *testing.T) {
	f, err := ioutil.TempFile("", "test-nested-*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString(`{"Server":{"Timeout":"10s"}}`); err != nil {
		t.Fatal(err)
	}

	cfg := struct {
		Server struct{ Timeout time.Duration }
	}{}

	err = FillJson(&cfg, f.Name())
	AssertNil(t, err)
	if cfg.Server.Timeout != 10*time.Second {
		t.Errorf("expected 10s, got %s", cfg.Server.Timeout)
	}
}
