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
	temp := ins
	fmt.Println(temp)

	seps := strings.Fields(temp)
	k := len(seps)
	var stop int
	for i := 0; i < k; i++ {
		if seps[i] == "id:" {
			stop = i

		}
	}

	seps = seps[0:stop]

	devName := strings.Join(seps, " ")
	fmt.Println(devName)
	return devName
}

// Returns prepared slice of strings to ttd.go (input the returned values from the above functions)
func NameData(loadout []string) []string {
	var names []string
	//fmt.Println(loadout)
	//fmt.Println(loadout[0])
	//fmt.Println(loadout[1])
	//words := strings.Fields(loadout[0])
	//var stop int
	//for i := 0; i < len(words); i++ {
	//	if words[i] == "id:" {
	//		stop = i
	//	}
	//}
	//words = words[0:stop]
	//fmt.Println(words)
	for i := 0; i < len(loadout); i++ {
		insert := getName(loadout[i])
		fmt.Println(insert)
		names = append(names, insert)
	}

	return names
}

// Parse through a device and obtain button info

// Parse through a device and obtain parameter info

// Parse through the list of devices
