package goconfig

import (
	"os"
)

// Read loads configuration and exits with status code 1 on error.
//
// For library code, prefer Load so the caller can handle errors.
func Read(c interface{}) {
	if err := Load(c); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

// ReadWithError loads configuration and returns any error.
func ReadWithError(c interface{}) error {
	return Load(c)
}

func readWithError(c interface{}) error {
	return Load(c)
}
