package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/joho/godotenv"
)

const red = "\033[0;31m"
const yellow = "\033[0;33m"
const green = "\033[0;32m"
const blue = "\033[0;34m"
const purple = "\033[0;35m"
const boldRed = "\033[1;31m"
const boldYellow = "\033[1;33m"
const boldPurple = "\033[1;35m"
const boldWhite = "\033[1;37m"
const reset = "\033[0m"

const cat = `   \    /\    %s              
    )  ( %s)   %s
   (  /  )    %s
    \(__)|    %s 
`

func main() {
	godotenv.Load("/etc/os-release")

	username := getCurrentUsername()
	hostname := getHostname()

	formattedHostName := fmt.Sprintf("%s %s%s@%s%s %s", boldRed, username, red, boldRed, hostname, reset)

	operatingSystem := getOperatingSystem()
	formattedOperatingSystem := fmt.Sprintf("%s OS: %s %s", yellow, operatingSystem, reset)

	kernelVersion := getKernelVersion()
	formattedKernelVersion := fmt.Sprintf("%s Kernel: %s %s", green, kernelVersion, reset)

	shell := os.Getenv("SHELL")
	formattedShell := fmt.Sprintf("%s Shell: %s %s", blue, shell, reset)

	formattedEye := fmt.Sprintf("%s'%s", boldYellow, reset)

	fmt.Printf(cat, formattedHostName, formattedEye, formattedOperatingSystem, formattedKernelVersion, formattedShell)
}

func getCurrentUsername() string {
	username, err := user.Current()

	if err != nil {
		return "unknown"
	}

	return username.Username
}

func getHostname() string {
	hostname, err := os.Hostname()

	if err != nil {
		return "unknown"
	}

	return hostname
}

func getOperatingSystem() string {
	return os.Getenv("NAME")
}

func getKernelVersion() string {
	file, err := os.ReadFile("/proc/version")

	if err != nil {
		return "unknown"
	}

	contents := string(file)

	return strings.Split(contents, " ")[2]
}
