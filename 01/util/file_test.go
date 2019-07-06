package util

import (
	"testing"
)

func TestCreateDirIfNotExist(t *testing.T) {
	CreateDirIfNotExist("../docs/json")
}

func TestMarshalJsonAndWriteFile(t *testing.T) {
	MarshalJsonAndWriteFile("test", "../docs/json/test.json", true)
}
