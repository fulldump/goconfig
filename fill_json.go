package goconfig

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func FillJson(c interface{}, filename string) error {

	if "" == filename {
		return nil
	}

	data, err := ioutil.ReadFile(filename)
	if nil != err {
		return err
	}

	err = json.Unmarshal(data, &c)
	if nil != err {
		return errors.New("Bad json file: " + err.Error())
	}

	return nil
}
