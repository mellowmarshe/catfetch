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
const gray = "\033[0;90m"
const purple = "\033[0;35m"
const cyan = "\033[0;36m"
const boldRed = "\033[1;31m"
const boldYellow = "\033[1;33m"
const boldBlue = "\033[1;34m"
const boldGreen = "\033[1;32m"
const boldPurple = "\033[1;35m"
const boldWhite = "\033[1;37m"
const reset = "\033[0m"

// We cant escape backticks in Go..
const cat = `.       .
\` + "`" + `-"'"-'/     %s
 } 6 6 {      %s
=.  Y  ,=     %s
  /^^^\  .    %s
 /     \  )   %s
(..)-(..)/    
`

func main() {
	godotenv.Load("/etc/os-release")

	username := getCurrentUsername()
	hostname := getHostname()

	formattedHostName := fmt.Sprintf("%s%s%s@%s%s%s", boldWhite, username, reset, boldWhite, hostname, reset)

	operatingSystem := getOperatingSystem()
	formattedOperatingSystem := fmt.Sprintf("%sos %s%s", red, reset, strings.ToLower(operatingSystem))

	kernelVersion := getKernelVersion()
	formattedKernelVersion := fmt.Sprintf("%skernel %s%s", green, reset, kernelVersion)

	shell := os.Getenv("SHELL")
	formattedShell := fmt.Sprintf("%sshell %s%s", blue, reset, shell)

	formattedColors := fmt.Sprintf("%s■■■%s■■■%s■■■%s■■■%s■■■%s■■■ %s", red, green, yellow, blue, purple, cyan, reset)

	fmt.Printf(cat, formattedHostName, formattedOperatingSystem, formattedKernelVersion, formattedShell, formattedColors)
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
