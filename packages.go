package main

func getPackageString(system OS) string {
	switch system {
	case ubuntu:
		return "apt"
	case fedora:
		return "dnf"
	default:
		return ""
	}
}

func installPackages(packageFile string, system OS) {

}
