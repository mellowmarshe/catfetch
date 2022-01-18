package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const red = "\033[0;31m"
const yellow = "\033[0;33m"
const green = "\033[0;32m"
const blue = "\033[0;34m"
const purple = "\033[0;35m"
const cyan = "\033[0;36m"
const boldRed = "\033[1;31m"
const boldYellow = "\033[1;33m"
const boldPurple = "\033[1;35m"
const boldWhite = "\033[1;37m"
const reset = "\033[0m"

const cat = `              %s   
   \    /\    %s             
    )  ( %s)   %s
   (  /  )    %s
    \(__)|    %s 
              
`
const compact = `%s
%s
%s
%s
%s
`

func main() {

	var isCompact bool

	flag.BoolVar(&isCompact, "compact", false, "If this flag is present the ASCII art cat wont be printed")
	flag.Parse()

	godotenv.Load("/etc/os-release")

	username := getCurrentUsername()
	hostname := getHostname()

	formattedHostName := fmt.Sprintf("%s%s%s@%s%s %s", boldRed, username, red, boldRed, hostname, reset)

	operatingSystem := getOperatingSystem()
	formattedOperatingSystem := fmt.Sprintf("%sOS: %s%s", yellow, reset, operatingSystem)

	kernelVersion := getKernelVersion()
	formattedKernelVersion := fmt.Sprintf("%sKernel: %s%s", green, reset, kernelVersion)

	shell := os.Getenv("SHELL")
	formattedShell := fmt.Sprintf("%sShell: %s%s", blue, reset, shell)

	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(10)

	var eyeCharacter string

	if rand > 7 {
		eyeCharacter = "o"
	} else {
		eyeCharacter = "'"
	}

	formattedEye := fmt.Sprintf("%s%s%s", boldYellow, eyeCharacter, reset)

	formattedColors := fmt.Sprintf("%s▇▇%s▇▇%s▇▇%s▇▇%s▇▇%s▇▇ %s", red, green, yellow, blue, purple, cyan, reset)

	if isCompact {
		fmt.Printf(compact, formattedHostName, formattedOperatingSystem, formattedKernelVersion, formattedShell, formattedColors)
	} else {
		fmt.Printf(cat, formattedHostName, formattedOperatingSystem, formattedEye, formattedKernelVersion, formattedShell, formattedColors)
	}

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
