package ttd

// Device Struct
type Device struct {
	Name   string
	id     int
	d_type string

	// profiles
	btnprofiles   []btnmaps
	paramprofiles []params
}

type btnmaps struct {
	btn []rune
}

type params struct {
	MapToOutput     string
	Area            string
	Suppress        int
	ToolDebugLevel  string
	RawSample       int
	Pressure        [4]int
	Mode            string
	Touch           string
	Gesture         string
	ZoomDistance    int
	ScrollDistance  int
	TapTime         int
	CursorProximity int
	Rotate          string
}

func NewDevice(dName string) *Device {
	d := Device{Name: dName}
	return &d
}
