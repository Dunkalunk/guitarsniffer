package drumjoypad

import (
	"dunkalunk/drumpacket"
	"github.com/artman41/vjoy"
)

const (
	minJoyID = 1
	maxJoyID = 16
)

// Button IDs
const (
	redDrum uint = iota
	yellowDrum
	blueDrum
	greenDrum
	yellowCymbal
	blueCymbal
	greenCymbal
	bassOne
	bassTwo
	dpadUp
	dpadDown
	dpadLeft
	dpadRight
	buttonMenu
	buttonOptions
)

// DrumJoypad is a Container for the JoypadDevice
// with utility functions baked in to retrieve the
// specific Buttons for drums and cymbals
type DrumJoypad struct {
	joypad *vjoy.Device
	rID    uint
}

// The following methods retrieve the Drums
func (drumJoypad DrumJoypad) RedDrum() *vjoy.Button {
	return drumJoypad.joypad.Button(redDrum)
}

func (drumJoypad DrumJoypad) YellowDrum() *vjoy.Button {
	return drumJoypad.joypad.Button(yellowDrum)
}

func (drumJoypad DrumJoypad) BlueDrum() *vjoy.Button {
	return drumJoypad.joypad.Button(blueDrum)
}

func (drumJoypad DrumJoypad) GreenDrum() *vjoy.Button {
	return drumJoypad.joypad.Button(greenDrum)
}

// The following methods retrieve the Cymbals
func (drumJoypad DrumJoypad) YellowCymbal() *vjoy.Button {
	return drumJoypad.joypad.Button(yellowCymbal)
}

func (drumJoypad DrumJoypad) BlueCymbal() *vjoy.Button {
	return drumJoypad.joypad.Button(blueCymbal)
}

func (drumJoypad DrumJoypad) GreenCymbal() *vjoy.Button {
	return drumJoypad.joypad.Button(greenCymbal)
}

// The following methods retrieve the Bass Pedals
func (drumJoypad DrumJoypad) BassOne() *vjoy.Button {
	return drumJoypad.joypad.Button(bassOne)
}

func (drumJoypad DrumJoypad) BassTwo() *vjoy.Button {
	return drumJoypad.joypad.Button(bassTwo)
}

// The following methods retrieve the Dpad Buttons
func (drumJoypad DrumJoypad) DpadUp() *vjoy.Button {
	return drumJoypad.joypad.Button(dpadUp)
}

func (drumJoypad DrumJoypad) DpadDown() *vjoy.Button {
	return drumJoypad.joypad.Button(dpadDown)
}

func (drumJoypad DrumJoypad) DpadLeft() *vjoy.Button {
	return drumJoypad.joypad.Button(dpadLeft)
}

func (drumJoypad DrumJoypad) DpadRight() *vjoy.Button {
	return drumJoypad.joypad.Button(dpadRight)
}

// ButtonMenu retrieves the Menu button
func (drumJoypad DrumJoypad) ButtonMenu() *vjoy.Button {
	return drumJoypad.joypad.Button(buttonMenu)
}

// ButtonOptions retrieves the Options button
func (drumJoypad DrumJoypad) ButtonOptions() *vjoy.Button {
	return drumJoypad.joypad.Button(buttonOptions)
}

func (drumJoypad DrumJoypad) SetDrumValues(drums drumpacket.Drums) {
	drumJoypad.GreenDrum().Set(drums.Green)
	drumJoypad.RedDrum().Set(drums.Red)
	drumJoypad.YellowDrum().Set(drums.Yellow)
	drumJoypad.BlueDrum().Set(drums.Blue)
	drumJoypad.BassOne().Set(drums.BassOne)
	drumJoypad.BassTwo().Set(drums.BassTwo)
}

func (drumJoypad DrumJoypad) SetCymbalValues(cymbals drumpacket.Cymbals) {
	drumJoypad.YellowCymbal().Set(cymbals.Yellow)
	drumJoypad.BlueCymbal().Set(cymbals.Blue)
	drumJoypad.GreenCymbal().Set(cymbals.Green)
}

func (drumJoypad DrumJoypad) SetDpadValues(dpad drumpacket.Dpad) {
	drumJoypad.DpadUp().Set(dpad.Up)
	drumJoypad.DpadDown().Set(dpad.Down)
	drumJoypad.DpadLeft().Set(dpad.Left)
	drumJoypad.DpadRight().Set(dpad.Right)
}

const maxFloat int = 0x7fff

func convertByte(b byte) int {
	fraction := float32(b) / float32(0xFF)
	return int(fraction * float32(maxFloat))
}

func (drumJoypad DrumJoypad) SetButtonValues(buttons drumpacket.Buttons) {
	drumJoypad.ButtonMenu().Set(buttons.Menu)
	drumJoypad.ButtonOptions().Set(buttons.Options)
}

func (drumJoypad DrumJoypad) SetValues(drumPacket drumpacket.DrumPacket) {
	drumJoypad.SetDrumValues(drumPacket.Drums)
	drumJoypad.SetCymbalValues(drumPacket.Cymbals)
	drumJoypad.SetDpadValues(drumPacket.Dpad)
	drumJoypad.SetButtonValues(drumPacket.Buttons)
}

// Update the vJoyDevice with the set
// Button & Axis values
func (drumJoypad DrumJoypad) Update() error {
	return drumJoypad.joypad.Update()
}

// Reset centers all Axes & resets all Buttons
func (drumJoypad DrumJoypad) Reset() {
	drumJoypad.joypad.Reset()
}

// Relinquish closes the joypad device
func (drumJoypad DrumJoypad) Relinquish() {
	drumJoypad.joypad.Relinquish()
}

// GetVirtualID returns the rID assigned by vJoy
func (drumJoypad DrumJoypad) GetVirtualID() uint {
	return drumJoypad.rID
}

// GetJoypad attempts to obtain a free Joypad
// with a Virtual Device ID between 1 and 16
func GetJoypad() (*DrumJoypad, error) {
	if !vjoy.Available() {
		return nil, ErrUnavailable
	}
	dev, rID, err := accquireJoypad()
	if err != nil {
		return nil, err
	}
	return &DrumJoypad{
		joypad: dev,
		rID:    rID,
	}, nil
}

func accquireJoypad() (dev *vjoy.Device, rID uint, err error) {
	var currentID uint = minJoyID
	dev, err = vjoy.Acquire(currentID)
	if err != nil {
		currentID++
	}
	for err == vjoy.ErrDeviceAlreadyOwned && currentID <= maxJoyID {
		dev, err = vjoy.Acquire(currentID)
		currentID++
	}
	if err != nil {
		return nil, 0, err
	}

	return dev, currentID, nil
}
