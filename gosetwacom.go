package main

import (
	"fmt"

	"github.com/shafinmalik/gosetwacom/tgsh"
	"github.com/shafinmalik/gosetwacom/ttd"
)

// FIXED - send to UAT and then merge into main.

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
	//
	// Current bug: in tgsh.go NameData: returns empty slice
	// Supposed to return slice of string with names

	for i := 0; i < len(nems); i++ {
		dp := ttd.NewDevice(nems[i])
		sampled := Entry{name: "name", catg: 0, pakg: *dp}
		rep = append(rep, sampled)
	}

	//sample0 := Entry{name: "Hello", catg: 0}
	//sample1 := Entry{name: "World", catg: 1}
	//sample2 := Entry{name: "Thing", catg: 1}
	//rep = append(rep, sample0, sample1, sample2)

	// Init CUI
}
