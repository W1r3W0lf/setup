package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Copy(src, dst string) error {

}

func installSettings(settingsDir string) {

	//settings := make([]string, 3)
	settings, err := filepath.Glob(settingsDir + "/*")
	if err != nil {
		fmt.Println("[ERROR] Getting settings")
		return
	}

	for _, setting := range settings {
		_, filename := filepath.Split(setting)
		dest := "~." + filename
		Copy(setting, dest)
	}

}

func copySettings(outPutDir string) {

}
