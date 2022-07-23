package comms

// use interface to communicate with os. Send info to tabfetch

type receiver struct {
	communicationName string
	comm              string
}

type sender struct {
	command string
	exec    bool
	e       error
}
