package main

import (
	"encoding/hex"
	"fmt"
	"runtime"
	"sync"

	"dunkalunk/drumjoypad"
	"dunkalunk/drumpacket"
	drumsniffer "dunkalunk/sniffer"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

var currentPacket = drumpacket.DrumPacket{}
var currentPacketHex = ""
var drumJoypad *drumjoypad.DrumJoypad
var RunDataThread = true

var threads sync.WaitGroup

var width, height int

func main() {
	go guiThread(&RunDataThread)
	threads.Add(1)
	dataThread(&RunDataThread)

	threads.Wait()
}

func guiThread(runDataThread *bool) {
	defer func() {
		threads.Done()
		*runDataThread = false
	}()

	runtime.LockOSThread()

	setupFontCache()

	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	width, height := 230, 100

	glfw.WindowHint(glfw.Resizable, glfw.False)
	window, err := glfw.CreateWindow(width, height, "Drum Sniffer", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	window.SetKeyCallback(onKey)

	glfw.SwapInterval(1)

	err = gl.Init()
	if err != nil {
		panic(err)
	}

	initDisplay(width, height)

	for !window.ShouldClose() {
		display(width, height)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch {
	case key == glfw.KeyEscape && action == glfw.Press,
		key == glfw.KeyQ && action == glfw.Press:
		w.SetShouldClose(true)
	}
}

func dataThread(runDataThread *bool) {
	fmt.Println("Getting Joypad...")
	joypad, err := drumjoypad.GetJoypad()
	if err != nil {
		panic(err)
	}
	drumJoypad = joypad
	fmt.Println("Obtained!")
	defer drumJoypad.Relinquish()

	fmt.Println("Starting Sniffer...")
	sniffer, err := drumsniffer.Start()
	defer sniffer.Stop()
	if err != nil {
		panic(err)
	}
	for *runDataThread {
		select {
		case packet := <-sniffer.Packets:
			handlePacket(&packet)
		default:
			continue
		}
	}
}

func handlePacket(packet *drumsniffer.Packet) {
	// The packet returned when pressing the Xbox button is 31
	// bytes long, not 40, meaning that currently we're
	// ignoring that it exists but the code is there for it
	if packet.CaptureInfo.Length != 36 {
		return
	}
	currentPacket = drumpacket.CreateDrumPacket(packet.Data[drumpacket.XboxHeaderLength:])
	currentPacketHex = hex.EncodeToString(packet.Data[drumpacket.XboxHeaderLength:])
	drumJoypad.SetValues(currentPacket)
	drumJoypad.Update()
}
