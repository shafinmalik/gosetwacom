package main

import (
	"fmt"

	"github.com/shafinmalik/gosetwacom/tgsh"
	"github.com/shafinmalik/gosetwacom/ttd"
)

// FIXED - send to UAT and then merge into main

var (
	currView      = 0
	currSelection = -1
	rep           = make([]Entry, 0) // Repository of Entries
)

type Entry struct {
	name string
	catg int // category (change name later?)
	pakg ttd.Device
}

func main() {
	// Get devices
	devs := tgsh.DeviceData()

	nems := tgsh.NameData(devs)
	fmt.Println(nems)
	for i := 0; i < len(nems); i++ {
		fmt.Println(nems[i])
	}
	// EDIT --> BUG HAS BEEN FIXED. SEND TO UAT AND REINSERT
	// CUI MATERIALS + REMOVE COMMENTS

	for i := 0; i < len(nems); i++ {
		dp := ttd.NewDevice(nems[i])
		sampled := Entry{name: "name", catg: 0, pakg: *dp}
		rep = append(rep, sampled)
	}

}
