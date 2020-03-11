package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// ONLY changed by flags at startup.
var dry bool
var packageFile string
var packagesOnly bool
var settingsDir string
var settingsOnly bool
var systemName string

// Setup the flags for the rest of the program.
func init() {
	flag.BoolVar(&dry, "dryrun", false,
		"Prevents the programs form being installed")
	flag.StringVar(&packageFile, "package", "./package.yaml",
		"File that lists packages to be installed")
	flag.BoolVar(&packagesOnly, "pkg", false,
		"Only install packages.")
	flag.StringVar(&settingsDir, "settings", "./settings/",
		"Location of settings to be installed.")
	flag.BoolVar(&settingsOnly, "set", false,
		"Only install settings.")
	flag.StringVar(&systemName, "OS", "",
		"Override the autodetect OS")
}

func testPath(path string) bool {
	_, err := os.Stat(path)
	return os.IsExist(err)
}

func main() {
	flag.Parse()

	var system OS
	if len(systemName) > 0 {
		system = getOSFromName(systemName)
	} else {
		system = detectOS()
	}

	if system == other {
		fmt.Println("OS is not supported.")
	} else if !settingsOnly {
		cleanPackageFile := filepath.Clean(packageFile)
		fmt.Println("Installing packages from ", cleanPackageFile)
		installPackages(packageFile, system)
	}

	if !packagesOnly {
		cleanSettingsDir := filepath.Clean(settingsDir)
		fmt.Println("Installing settings from ", cleanSettingsDir)
		installSettings(settingsDir)
	}

}
