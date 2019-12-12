package drumpacket

const XboxHeaderLength = 22

// Drums
const (
	RedDrum = 0x20

	YellowDrum1 = 0x1
	YellowDrum2 = 0x2
	YellowDrum3 = 0x3
	YellowDrum4 = 0x4
	YellowDrum5 = 0x5
	YellowDrum6 = 0x6
	YellowDrum7 = 0x7

	BlueDrum10 = 0x10
	BlueDrum20 = 0x20
	BlueDrum30 = 0x30
	BlueDrum40 = 0x40
	BlueDrum50 = 0x50
	BlueDrum60 = 0x60

	GreenDrum = 0x10
	BassOne   = 0x10
	BassTwo   = 0x20
)

// Cymbals
const (
	YellowCymbal10 = 0x10
	YellowCymbal20 = 0x20
	YellowCymbal30 = 0x30
	YellowCymbal40 = 0x40
	YellowCymbal50 = 0x50
	YellowCymbal60 = 0x60
	YellowCymbal70 = 0x70

	BlueCymbal1 = 0x1
	BlueCymbal2 = 0x2
	BlueCymbal3 = 0x3
	BlueCymbal4 = 0x4
	BlueCymbal5 = 0x5
	BlueCymbal6 = 0x6
	BlueCymbal7 = 0x7

	GreenCymbal10 = 0x10
	GreenCymbal20 = 0x20
	GreenCymbal30 = 0x30
	GreenCymbal40 = 0x40
	GreenCymbal50 = 0x50
	GreenCymbal60 = 0x60
	GreenCymbal70 = 0x70
)

// Buttons
const (
	ButtonXbox    = 0x1
	ButtonMenu    = 0x4
	ButtonOptions = 0x8
)

const (
	DpadDown  = 0x1
	DpadUp    = 0x2
	DpadLeft  = 0x4
	DpadRight = 0x8
)

// Packet Pieces
const (
	PosRedDrum      = 8
	PosYellowDrum   = 10
	PosBlueDrum     = 11
	PosGreenDrum    = 8
	PosBassPedal    = 9
	PosYellowCymbal = 12
	PosBlueCymbal   = 12
	PosGreenCymbal  = 13

	PosButtons = 8
	PosDpad    = 9
)

// CreateDrumPacket returns a DrumPacket struct
// filled with the values of the given packet
//
// Note: the function assumes that the given packet
// has already had the XboxHeader removed from it
func CreateDrumPacket(packet []byte) DrumPacket {
	// fmt.Printf("(%d) %s\n", len(packet), hex.EncodeToString(packet))
	drums := getDrums(packet)
	cymbals := getCymbals(packet)
	dpad := getDpad(packet[PosDpad])
	buttons := getButtons(packet[PosButtons])
	return DrumPacket{
		Drums:   drums,
		Cymbals: cymbals,
		Dpad:    dpad,
		Buttons: buttons,
	}
}

func getDrums(packet []byte) Drums {
	return Drums{
		Red:     packet[PosRedDrum]&RedDrum != 0,
		Yellow:  getYellowDrum(packet[PosYellowDrum]),
		Blue:    getBlueDrum(packet[PosBlueDrum]),
		Green:   packet[PosGreenDrum]&GreenDrum != 0,
		BassOne: packet[PosBassPedal]&BassOne != 0,
		BassTwo: packet[PosBassPedal]&BassTwo != 0,
	}
}

func getYellowDrum(drumBitMask byte) bool {
	if drumBitMask == 0 {
		return false
	} else if drumBitMask&YellowDrum1 != 0 ||
		drumBitMask&YellowDrum2 != 0 ||
		drumBitMask&YellowDrum3 != 0 ||
		drumBitMask&YellowDrum4 != 0 ||
		drumBitMask&YellowDrum5 != 0 ||
		drumBitMask&YellowDrum6 != 0 ||
		drumBitMask&YellowDrum7 != 0 {
		return true
	} else {
		return false
	}
}

func getBlueDrum(drumBitMask byte) bool {
	if drumBitMask == 0 {
		return false
	} else if drumBitMask&BlueDrum10 != 0 ||
		drumBitMask&BlueDrum20 != 0 ||
		drumBitMask&BlueDrum30 != 0 ||
		drumBitMask&BlueDrum40 != 0 ||
		drumBitMask&BlueDrum50 != 0 ||
		drumBitMask&BlueDrum60 != 0 {
		return true
	} else {
		return false
	}
}

func getCymbals(packet []byte) Cymbals {
	return Cymbals{
		Yellow: getYellowCymbal(packet[PosYellowCymbal]),
		Blue:   getBlueCymbal(packet[PosBlueCymbal]),
		Green:  getGreenCymbal(packet[PosGreenCymbal]),
	}
}

func getYellowCymbal(drumBitMask byte) bool {
	if drumBitMask == 0 {
		return false
	} else if drumBitMask&YellowCymbal10 != 0 ||
		drumBitMask&YellowCymbal20 != 0 ||
		drumBitMask&YellowCymbal30 != 0 ||
		drumBitMask&YellowCymbal40 != 0 ||
		drumBitMask&YellowCymbal50 != 0 ||
		drumBitMask&YellowCymbal60 != 0 ||
		drumBitMask&YellowCymbal70 != 0 {
		return true
	} else {
		return false
	}
}

func getBlueCymbal(drumBitMask byte) bool {
	if drumBitMask == 0 {
		return false
	} else if drumBitMask&BlueCymbal1 != 0 ||
		drumBitMask&BlueCymbal2 != 0 ||
		drumBitMask&BlueCymbal3 != 0 ||
		drumBitMask&BlueCymbal4 != 0 ||
		drumBitMask&BlueCymbal5 != 0 ||
		drumBitMask&BlueCymbal6 != 0 ||
		drumBitMask&BlueCymbal7 != 0 {
		return true
	} else {
		return false
	}
}

func getGreenCymbal(drumBitMask byte) bool {
	if drumBitMask == 0 {
		return false
	} else if drumBitMask&GreenCymbal10 != 0 ||
		drumBitMask&GreenCymbal20 != 0 ||
		drumBitMask&GreenCymbal30 != 0 ||
		drumBitMask&GreenCymbal40 != 0 ||
		drumBitMask&GreenCymbal50 != 0 ||
		drumBitMask&GreenCymbal60 != 0 ||
		drumBitMask&GreenCymbal70 != 0 {
		return true
	} else {
		return false
	}
}

func getButtons(buttonBitMask byte) Buttons {
	return Buttons{
		Menu:    buttonBitMask&ButtonMenu != 0,
		Options: buttonBitMask&ButtonOptions != 0,
		Xbox:    buttonBitMask&ButtonXbox != 0,
	}
}

func getDpad(dpadBitMask byte) Dpad {
	return Dpad{
		Up:    dpadBitMask&DpadUp != 0,
		Down:  dpadBitMask&DpadDown != 0,
		Left:  dpadBitMask&DpadLeft != 0,
		Right: dpadBitMask&DpadRight != 0,
	}
}

type Drums struct {
	Red, Yellow, Blue, Green, BassOne, BassTwo bool
}

type Cymbals struct {
	Yellow, Blue, Green bool
}

type Buttons struct {
	Menu, Options, Xbox bool
}

type Dpad struct {
	Up, Down, Left, Right bool
}

type DrumPacket struct {
	Drums   Drums
	Cymbals Cymbals
	Dpad    Dpad
	Buttons Buttons
}
