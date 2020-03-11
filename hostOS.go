package main

import (
	"fmt"
	"os"
	"time"
)

// ENUM in go
type OS int

const (
	ubuntu OS = iota
	fedora
	other
)

// Try and detect the OS by looking for it's package manager.
func detectOS() OS {

	if testPath("/usr/bin/apt") {
		return ubuntu
	} else if testPath("/usr/bin/dnf") {
		return fedora
	}

	return other
}

func getOSName(system OS) string {
	switch system {
	case ubuntu:
		return "ubuntu"
	case fedora:
		return "fedora"
	default:
		return "other"
	}
}

func getOSFromName(osName string) OS {
	if osName == "ubuntu" {
		return ubuntu
	} else if osName == "fedora" {
		return fedora
	} else {
		return other
	}
}

func selectOS() OS {
	var numOS int
	done := false
	for done {
		fmt.Println("Select your OS")
		fmt.Println("ubuntu		[0]")
		fmt.Println("fedora		[1]")
		fmt.Println("other		[2]")
		fmt.Print("#>")
		fmt.Scanf("%d", &numOS)
		if numOS < 3 {
			done = true
		} else {
			fmt.Println("That is not a valid input.")
			time.Sleep(3 * time.Second)
		}
	}
	switch numOS {
	case 0:
		fmt.Println("Using ubuntu as OS.")
		return ubuntu
	case 1:
		fmt.Println("Using fedora as OS.")
		return fedora
	default:
		fmt.Println("Unknown OS.")
		return other
	}
}

func askOS() OS {
	myOS := detectOS()
	var input string

	fmt.Println("Is ", getOSName(myOS), " your OS?")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else if input == "Y" || input == "y" {
		return myOS
	} else {
		return selectOS()
	}

	return myOS
}
