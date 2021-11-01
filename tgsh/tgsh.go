package tgsh

import (
	"fmt"
	"os/exec"
	"strings"
)

// Obtain device data as a slice of strings
func DeviceData() []string {
	app := "xsetwacom"

	arg0 := "list"
	arg1 := "devices"

	cmd := exec.Command(app, arg0, arg1)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	// split devices into array of strings
	loadout := strings.Split(string(stdout), "\n")
	return loadout
}

// Convert slice of strings into individual strings
func getName(ins string) string {
	var name []string

	temp := ins
	seps := strings.Fields(temp)
	k := len(seps)
	for i := 0; i < k; i++ {
		if seps[i] == "id:" {
			name = append(seps[0:i])
			fmt.Println(name)
		}
	}

	devName := strings.Join(name, " ")
	return devName
}

// Returns prepared slice of strings to ttd.go (input the returned values from the above functions)
func NameData(loadout []string) []string {
	var names []string
	for i := 0; i < len(names); i++ {
		insert := getName(loadout[i])
		names = append(names, insert)
	}

	return names
}

// Parse through a device and obtain button info

// Parse through a device and obtain parameter info

// Parse through the list of devices
