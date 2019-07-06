package command

import (
	"os"
	"testing"
)

func TestRoot(t *testing.T) {
	_ = Execute()
}

func TestParse(t *testing.T) {
	if _, err := os.Stat("../file/GeoLite2-City-Locations-en.csv"); !os.IsNotExist(err) {
		parseCmd := NewParseCommand()
		parseCmd.SetArgs([]string{"--root", "..", "--command", "generate", "--location", "../file/GeoLite2-City-Locations-en.csv"})
		_ = parseCmd.Execute()
		parseCmd.SetArgs([]string{"--root", "..", "--command", "validate"})
		_ = parseCmd.Execute()
	}
}
