package goconfig

import (
	"errors"
	"strings"
	"path/filepath"
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	"gopkg.in/yaml.v2"
)

type providerType int

const (
	typeJson providerType = iota
	typeYaml
)

var providers map[providerType]func([]byte) provider

func init() {
	providers = make(map[providerType]func([]byte) provider)
	providers[typeJson] = func(data []byte) provider {
		return &jsonProvider{data}
	}
	providers[typeYaml] = func(data []byte) provider {
		return &ymlProvider{data}
	}
}

type (
	provider interface {
		which() providerType
		fill(c interface{})
	}

	jsonProvider struct {
		data []byte
	}

	ymlProvider struct {
		data []byte
	}
)

func (pJson *jsonProvider) which() providerType {
	return typeJson
}

func (pJson *jsonProvider) fill(c interface{}) {
	err := json.Unmarshal(pJson.data, &c)
	if nil != err {
		fmt.Println("Config file should be a valid JSON")
		os.Exit(1)
	}
}

func (json *ymlProvider) which() providerType {
	return typeYaml
}
func (pYml *ymlProvider) fill(c interface{}) {
	err := yaml.Unmarshal(pYml.data, &c)
	if nil != err {
		fmt.Println("Config file should be a valid YAML")
		os.Exit(1)
	}
}

func getProvider(configPath string) (provider, error) {
	if !isFile(configPath) {
		return nil, errors.New("a single file must be provided")
	}
	pType, err := getProviderType(configPath)
	if err != nil {
		return nil, err
	}
	pFn, ok := providers[pType]
	if !ok {
		return nil, errors.New("not recognized provider type")
	}
	data, err := ioutil.ReadFile(configPath)
	if nil != err {
		fmt.Println("Unable to read config file `" + configPath + "`!")
		os.Exit(1)
	}
	return pFn(data), nil
}

func getProviderType(path string) (providerType, error) {
	ext := strings.ToLower(filepath.Ext(path))
	if isJson(ext) {
		return typeJson, nil
	} else if isYml(ext) {
		return typeYaml, nil
	}
	return -1, errors.New("unsupported file format")
}

func isJson(ext string) bool {
	return "json" == ext
}

func isYml(ext string) bool {
	return "yaml" == ext || "yml" == ext
}
