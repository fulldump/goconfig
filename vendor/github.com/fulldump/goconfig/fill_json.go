package goconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func FillJson(c interface{}, filename string) {

	if "" == filename {
		return
	}

	data, err := ioutil.ReadFile(filename)
	if nil != err {
		fmt.Println("Unable to read config file `" + filename + "`!")
		os.Exit(1)
	}

	err = json.Unmarshal(data, &c)
	if nil != err {
		fmt.Println("Config file should be a valid JSON")
		os.Exit(1)
	}
}
