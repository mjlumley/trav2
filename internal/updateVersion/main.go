package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"time"
)

// The name of the configuration file we are updating
const configFile = "..\\..\\cmd\\traveller\\config.go"

// main runs updateVersion which increments the minor version and the datetime stamp of the
// version information in the config.go file in the traveller base code.
func main() {
	// Read in config file
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	//fmt.Println(b) // print the content as 'bytes'
	str := string(b) // convert content to a 'string'
	//fmt.Println(str) // print the content as a 'string'

	// Define regexp
	r, _ := regexp.Compile("(appVersion string = \\\"\\d+\\.\\d+)\\.(\\d+)\\.(\\d+)\"")

	locs := r.FindStringSubmatchIndex(str)
	base := str[locs[2]:locs[3]]
	minor := str[locs[4]:locs[5]]
	timePart := str[locs[6]:locs[7]]

	fmt.Println(r.FindStringSubmatchIndex(str))
	fmt.Println("Old details: minor: " + minor + ", timePart: " + timePart)

	// New details
	t := time.Now().Format("20060102150405")
	fmt.Println(" New time: " + t)
	newMinorInt, _ := strconv.Atoi(minor)
	newMinorInt++
	newMinorStr := strconv.Itoa(newMinorInt)
	fmt.Println(" New minor int: " + newMinorStr)

	newLine := base + "." + newMinorStr + "." + t + "\""
	fmt.Println("newLine: " + newLine)

	out := r.ReplaceAll(b, []byte(newLine))
	fmt.Println(string(out))

	// We are ready to write out
	err = ioutil.WriteFile(configFile, out, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Config file rewritten!")

}
