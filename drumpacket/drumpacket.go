package drumpacket

const XboxHeaderLength = 22

// Drums
const (
	RedDrum = 0x14

	YellowDrum1 = 0x1
	YellowDrum2 = 0x2
	YellowDrum3 = 0x3
	YellowDrum4 = 0x4
	YellowDrum5 = 0x5
	YellowDrum6 = 0x6
	YellowDrum7 = 0x7

	BlueDrum10 = 0xA
	BlueDrum20 = 0x14
	BlueDrum30 = 0x1E
	BlueDrum40 = 0x28
	BlueDrum50 = 0x32
	BlueDrum60 = 0x3C

	GreenDrum = 0xA
	BassOne   = 0xA
	BassTwo   = 0x14
)

// Cymbals
const (
	YellowCymbal10 = 0xA
	YellowCymbal20 = 0x14
	YellowCymbal30 = 0x1E
	YellowCymbal40 = 0x28
	YellowCymbal50 = 0x32
	YellowCymbal60 = 0x3C
	YellowCymbal70 = 0x46

	BlueCymbal1 = 0x1
	BlueCymbal2 = 0x2
	BlueCymbal3 = 0x3
	BlueCymbal4 = 0x4
	BlueCymbal5 = 0x5
	BlueCymbal6 = 0x6
	BlueCymbal7 = 0x7

	GreenCymbal10 = 0xA
	GreenCymbal20 = 0x14
	GreenCymbal30 = 0x1E
	GreenCymbal40 = 0x28
	GreenCymbal50 = 0x32
	GreenCymbal60 = 0x3C
	GreenCymbal70 = 0x46
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
		Red:     getRedDrum(packet[PosRedDrum]),
		Yellow:  getYellowDrum(packet[PosYellowDrum]),
		Blue:    getBlueDrum(packet[PosBlueDrum]),
		Green:   getGreenDrum(packet[PosGreenDrum]),
		BassOne: packet[PosBassPedal]&BassOne != 0,
		BassTwo: packet[PosBassPedal]&BassTwo != 0,
	}
}

func getRedDrum(drumBitMask byte) bool {
	if drumBitMask&RedDrum != 0 {
		return true
	} else {
		return false
	}
}

func getYellowDrum(drumBitMask byte) bool {
	if drumBitMask&YellowDrum1 != 0 {
		return true
	} else if drumBitMask&YellowDrum2 != 0 {
		return true
	} else if drumBitMask&YellowDrum3 != 0 {
		return true
	} else if drumBitMask&YellowDrum4 != 0 {
		return true
	} else if drumBitMask&YellowDrum5 != 0 {
		return true
	} else if drumBitMask&YellowDrum6 != 0 {
		return true
	} else if drumBitMask&YellowDrum7 != 0 {
		return true
	} else {
		return false
	}
}

func getBlueDrum(drumBitMask byte) bool {
	if drumBitMask&BlueDrum10 != 0 {
		return true
	} else if drumBitMask&BlueDrum20 != 0 {
		return true
	} else if drumBitMask&BlueDrum30 != 0 {
		return true
	} else if drumBitMask&BlueDrum40 != 0 {
		return true
	} else if drumBitMask&BlueDrum50 != 0 {
		return true
	} else if drumBitMask&BlueDrum60 != 0 {
		return true
	} else {
		return false
	}
}

func getGreenDrum(drumBitMask byte) bool {
	if drumBitMask&GreenDrum != 0 {
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
	if drumBitMask&YellowCymbal10 != 0 {
		return true
	} else if drumBitMask&YellowCymbal20 != 0 {
		return true
	} else if drumBitMask&YellowCymbal30 != 0 {
		return true
	} else if drumBitMask&YellowCymbal40 != 0 {
		return true
	} else if drumBitMask&YellowCymbal50 != 0 {
		return true
	} else if drumBitMask&YellowCymbal60 != 0 {
		return true
	} else if drumBitMask&YellowCymbal70 != 0 {
		return true
	} else {
		return false
	}
}

func getBlueCymbal(drumBitMask byte) bool {
	if drumBitMask&BlueCymbal1 != 0 {
		return true
	} else if drumBitMask&BlueCymbal2 != 0 {
		return true
	} else if drumBitMask&BlueCymbal3 != 0 {
		return true
	} else if drumBitMask&BlueCymbal4 != 0 {
		return true
	} else if drumBitMask&BlueCymbal5 != 0 {
		return true
	} else if drumBitMask&BlueCymbal6 != 0 {
		return true
	} else if drumBitMask&BlueCymbal7 != 0 {
		return true
	} else {
		return false
	}
}

func getGreenCymbal(drumBitMask byte) bool {
	if drumBitMask&GreenCymbal10 != 0 {
		return true
	} else if drumBitMask&GreenCymbal20 != 0 {
		return true
	} else if drumBitMask&GreenCymbal30 != 0 {
		return true
	} else if drumBitMask&GreenCymbal40 != 0 {
		return true
	} else if drumBitMask&GreenCymbal50 != 0 {
		return true
	} else if drumBitMask&GreenCymbal60 != 0 {
		return true
	} else if drumBitMask&GreenCymbal70 != 0 {
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