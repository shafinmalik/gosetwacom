package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
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

// Displays main layout with instructions
// Text updates should be made in this function
func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("Help", maxX-25, maxY*0+1, maxX-1, 9); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "↑ ↓: Scroll Selection")
		fmt.Fprintln(v, "← →: Change View")
		fmt.Fprintln(v, "^C: Exit")
		v.Title = "Help"
	}

	if currView == 0 {
		// Initialize Library and clear the others
		LibraryView(g)

	} else if currView == 1 {
		// Initialize Informatics and clear the others
		InformaticsView(g)

	} else if currView == 2 {
		// Initialize Data and clear the others
		DataView(g)

	} else {
		// Initialize error/panic. This should not happen
		panic("Defer")
	}

	return nil
}

// Need Keybinds for left, right, enter, ctrl-C initially
func initKeybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		return err
	}

	// Down/Up Arrows
	if err := g.SetKeybinding("Library", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		LibraryView(g)
		return err
	}

	if err := g.SetKeybinding("Library", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		LibraryView(g)
		return err
	}

	return nil
}

// TODO:
// Library Layout Template Function
func LibraryView(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("Library", maxX*0+1, maxY*0+5, maxX/2-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Library"
		if len(rep) > 0 {
			currSelection = 0
			v.Highlight = true
		}
		v.SelBgColor = gocui.ColorWhite
		for i := 0; i < len(rep); i++ {
			fmt.Fprintln(v, rep[i].name)
		}

	}

	v, err := g.SetView("Entries", maxX/2+1, maxY*0+5, maxX-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Info"
		fmt.Fprintln(v, currSelection)
	}

	// Focuses on Library
	if _, err := g.SetCurrentView("Library"); err != nil {
		return err
	}

	return nil
}

// Library View Clear Function
func ClearLibrary(g *gocui.Gui) error {
	if err := g.DeleteView("Library"); err != nil {
		return err
	}

	if err := g.DeleteView("Entries"); err != nil {
		return err
	}

	return nil
}

// Informatics Layout Template Function
func InformaticsView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("Functions", maxX*0+1, maxY*0+5, maxX/2-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Informatics"
		fmt.Fprintln(v, "Function 1")
	}
	return nil
}

// Informatics View Clear Function

// Database Layout Template Function
func DataView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("SQL", maxX*0+1, maxY*0+5, maxX/2-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "SQL"
		fmt.Fprintln(v, "Function 1")
	}
	return nil
}

// Database View Clear Function

// Arrow Key functions
func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()

		// prevent scroll down if length of rep will be exceeded
		if cy+1 > len(rep)-1 {
			return nil
		} else {
			currSelection = cy + 1
		}
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}

		}
	}

	if err := g.DeleteView("Entries"); err != nil {
		return err
	}

	maxX, maxY := g.Size()
	v, err := g.SetView("Entries", maxX/2+1, maxY*0+5, maxX-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Info"
		fmt.Fprintln(v, currSelection)
	}
	//ClearLibrary(g)
	//LibraryView(g)
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()

		if cy-1 >= 0 {
			currSelection = cy - 1
		}

		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
			//currSelection = cy - 1
		}
	}

	if err := g.DeleteView("Entries"); err != nil {
		return err
	}

	maxX, maxY := g.Size()
	v, err := g.SetView("Entries", maxX/2+1, maxY*0+5, maxX-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Info"
		fmt.Fprintln(v, currSelection)
	}
	//ClearLibrary(g)
	//LibraryView(g)
	return nil
}

// Add Entry. Use this (incomplete) function to add to the repository
func addEntry(s string, r *[]Entry) {
	nr := Entry{name: s}
	nr.catg = 0

	*r = append(*r, nr)
}

// Returns false if unable to remove
func remEntry(s string, r *[]Entry) bool {
	return false
}
