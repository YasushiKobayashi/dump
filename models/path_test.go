package models

import (
	"fmt"
	"testing"
)

func Test_GetFileName(t *testing.T) {
	test := "test"
	p := Path{
		FileName: test,
	}

	fileName := p.GetFileName()
	if fileName != fmt.Sprintf("%s.sql", test) {
		t.Fatalf("faile test")
	}
}

func Test_GetFullFilePath(t *testing.T) {
	path := "path"
	test := "test"
	p := Path{
		FilePath: path,
		FileName: test,
	}

	filePath := p.GetFullFilePath()
	if filePath != fmt.Sprintf("%s/%s.sql", path, test) {
		t.Fatalf("faile test")
	}
}
