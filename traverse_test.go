package goconfig

import (
	"fmt"
	"testing"
)

func TestTraverse(t *testing.T) {

	s := struct {
		Name     string
		Age      int
		MyStruct struct {
			Subname string
		}
		MyList []string
	}{
		MyList: []string{"one", "two", "three"},
	}

	traverse(&s, func(i item) {
		fmt.Println(i.Path)
	})

	t.Error()

}
