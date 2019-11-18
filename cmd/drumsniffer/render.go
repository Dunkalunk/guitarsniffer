package main

import (
	"dunkalunk/drumpacket"
	"image"
	"image/color"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/golang/freetype/truetype"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dgl"
	"github.com/llgcode/draw2d/draw2dkit"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/goitalic"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/goregular"
)

var (
	colourWhite  = rgb(0xff, 0xff, 0xff)
	colourGrey   = rgb(81, 81, 81)
	colourGreen  = rgb(51, 135, 2)
	colourRed    = rgb(211, 10, 10)
	colourYellow = rgb(204, 193, 0)
	colourBlue   = rgb(0, 43, 173)
	colourOrange = rgb(206, 102, 16)
)

type glBool struct {
	Condition *bool
	Off, On   color.RGBA
}

func rgb(r, g, b uint8) color.RGBA {
	return color.RGBA{R: r, G: g, B: b, A: 0xff}
}

func setupFontCache() {
	TTFs := map[string]([]byte){
		"goregular": goregular.TTF,
		"gobold":    gobold.TTF,
		"goitalic":  goitalic.TTF,
		"gomono":    gomono.TTF,
	}
	for fontName, TTF := range TTFs {
		font, err := truetype.Parse(TTF)
		if err != nil {
			panic(err)
		}
		draw2d.GetGlobalFontCache().Store(draw2d.FontData{Name: fontName}, font)
	}
}

func initDisplay(w, h int) {
	gl.ClearColor(0, 0, 0, 0)
	/* Establish viewing area to cover entire window. */
	gl.Viewport(0, 0, int32(w), int32(h))
	/* PROJECTION Matrix mode. */
	gl.MatrixMode(gl.PROJECTION)
	/* Reset project matrix. */
	gl.LoadIdentity()
	/* Map abstract coords directly to window coords. */
	gl.Ortho(0, float64(w), 0, float64(h), -1, 1)
	/* Invert Y axis so increasing Y goes down. */
	gl.Scalef(1, -1, 1)
	/* Shift origin up to upper-left corner. */
	gl.Translatef(0, float32(-h), 0)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Disable(gl.DEPTH_TEST)
}

func display(width, height int) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.LineWidth(1)
	gc := draw2dgl.NewGraphicContext(width, height)
	drawDrumGroup(gc, &currentPacket.Drums, 10, 10, "Drums")
	drawCymbalGroup(gc, &currentPacket.Cymbals, 10, 50, "Cymbals")
	drawDpad(gc, 130, 10, &currentPacket.Dpad)
	drawButtonGroup(gc, 130, 50, &currentPacket.Buttons)
	drawPacket(gc, 5, 90, &currentPacketHex)

	gl.Flush() /* Single buffered, so needs a flush. */
}

var ychange = float64(15)

func drawDrumGroup(gc *draw2dgl.GraphicContext, drums *drumpacket.Drums, x, y float64, str string) {
	drumXPos := func(baseX float64, i int) float64 {
		return baseX + float64(i*15)
	}
	var i = 0
	gc.SetFontData(draw2d.FontData{Name: "goregular"})
	gc.SetFillColor(image.White)
	gc.SetFontSize(10)
	gc.FillStringAt(str, drumXPos(x, i), y)

	drawDrum(gc, drumXPos(x, i), y+ychange, &glBool{
		Condition: &drums.Red,
		Off:       colourGrey,
		On:        colourRed,
	})
	i++
	drawDrum(gc, drumXPos(x, i), y+ychange, &glBool{
		Condition: &drums.Yellow,
		Off:       colourGrey,
		On:        colourYellow,
	})
	i++
	drawDrum(gc, drumXPos(x, i), y+ychange, &glBool{
		Condition: &drums.Blue,
		Off:       colourGrey,
		On:        colourBlue,
	})
	i++
	drawDrum(gc, drumXPos(x, i), y+ychange, &glBool{
		Condition: &drums.Green,
		Off:       colourGrey,
		On:        colourGreen,
	})
	i++
	drawDrum(gc, drumXPos(x, i), y+ychange, &glBool{
		Condition: &drums.BassOne,
		Off:       colourGrey,
		On:        colourOrange,
	})
	i++
	drawDrum(gc, drumXPos(x, i), y+ychange, &glBool{
		Condition: &drums.BassTwo,
		Off:       colourGrey,
		On:        colourOrange,
	})
}

func drawCymbalGroup(gc *draw2dgl.GraphicContext, cymbals *drumpacket.Cymbals, x, y float64, str string) {
	cymbalXPos := func(baseX float64, i int) float64 {
		return baseX + float64(i*15)
	}
	var i = 0
	gc.SetFontData(draw2d.FontData{Name: "goregular"})
	gc.SetFillColor(image.White)
	gc.SetFontSize(10)
	gc.FillStringAt(str, cymbalXPos(x, i), y)

	drawDrum(gc, cymbalXPos(x, i), y+ychange, &glBool{
		Condition: &cymbals.Yellow,
		Off:       colourGrey,
		On:        colourYellow,
	})
	i++
	drawDrum(gc, cymbalXPos(x, i), y+ychange, &glBool{
		Condition: &cymbals.Blue,
		Off:       colourGrey,
		On:        colourBlue,
	})
	i++
	drawDrum(gc, cymbalXPos(x, i), y+ychange, &glBool{
		Condition: &cymbals.Green,
		Off:       colourGrey,
		On:        colourGreen,
	})
	i++
}

func drawDrum(gc *draw2dgl.GraphicContext, x, y float64, glbool *glBool) {
	gc.BeginPath()
	draw2dkit.Circle(gc, x, y, 6)

	var fillColour = glbool.Off
	if *glbool.Condition {
		fillColour = glbool.On
	}
	gc.SetFillColor(fillColour)

	gc.Fill()
}

type dpadDirection uint8

var (
	dpadUp    dpadDirection = 0
	dpadDown  dpadDirection = 1
	dpadLeft  dpadDirection = 2
	dpadRight dpadDirection = 3
)

func drawDpad(gc *draw2dgl.GraphicContext, x, y float64, dpad *drumpacket.Dpad) {
	gc.SetFontData(draw2d.FontData{Name: "goregular"})
	gc.SetFillColor(image.White)
	gc.SetFontSize(10)
	gc.FillStringAt("Dpad", x-12.5, y)

	drawDpadButton(gc, x, y+15, dpadUp, &glBool{
		Condition: &dpad.Up,
		Off:       colourGrey,
		On:        colourWhite,
	})
	drawDpadButton(gc, x, y+15, dpadDown, &glBool{
		Condition: &dpad.Down,
		Off:       colourGrey,
		On:        colourWhite,
	})
	drawDpadButton(gc, x, y+15, dpadLeft, &glBool{
		Condition: &dpad.Left,
		Off:       colourGrey,
		On:        colourWhite,
	})
	drawDpadButton(gc, x, y+15, dpadRight, &glBool{
		Condition: &dpad.Right,
		Off:       colourGrey,
		On:        colourWhite,
	})
}

func drawDpadButton(gc *draw2dgl.GraphicContext, x, y float64, dir dpadDirection, glbool *glBool) {
	gc.BeginPath()

	switch dir {
	case dpadUp:
		draw2dkit.Rectangle(gc, x-2.5, y+5, x+2.5, y+10)
		break
	case dpadDown:
		draw2dkit.Rectangle(gc, x-2.5, y-5, x+2.5, y-10)
		break
	case dpadLeft:
		draw2dkit.Rectangle(gc, x-5, y-2.5, x-10, y+2.5)
		break
	case dpadRight:
		draw2dkit.Rectangle(gc, x+5, y-2.5, x+10, y+2.5)
		break
	}

	var fillColour = glbool.Off
	if *glbool.Condition {
		fillColour = glbool.On
	}
	gc.SetFillColor(fillColour)

	gc.Fill()
}

func drawButtonGroup(gc *draw2dgl.GraphicContext, x, y float64, buttons *drumpacket.Buttons) {
	gc.SetFontData(draw2d.FontData{Name: "goregular"})
	gc.SetFillColor(image.White)
	gc.SetFontSize(10)
	gc.FillStringAt("Buttons", x-20, y)

	drawButton(gc, x-7.5, y+10, &glBool{
		Condition: &buttons.Options,
		Off:       colourGrey,
		On:        colourWhite,
	})
	drawButton(gc, x+7.5, y+10, &glBool{
		Condition: &buttons.Menu,
		Off:       colourGrey,
		On:        colourWhite,
	})
}

func drawButton(gc *draw2dgl.GraphicContext, x, y float64, glbool *glBool) {
	gc.BeginPath()
	draw2dkit.Circle(gc, x, y, 5)

	var fillColour = glbool.Off
	if *glbool.Condition {
		fillColour = glbool.On
	}
	gc.SetFillColor(fillColour)

	gc.Fill()
}

func drawAxis(gc *draw2dgl.GraphicContext, x, y, fraction float64, str string) {
	gc.SetFontData(draw2d.FontData{Name: "goregular"})
	gc.SetFillColor(image.White)
	gc.SetFontSize(10)

	gc.FillStringAt(str, x, y)

	gc.BeginPath()
	draw2dkit.Rectangle(gc, x, y+5, x+30, y+10)
	gc.SetFillColor(colourGrey)
	gc.Fill()

	gc.BeginPath()
	draw2dkit.Rectangle(gc, x, y+5, x+fraction*30, y+10)
	gc.SetFillColor(colourRed)
	gc.Fill()
}

func drawPacket(gc *draw2dgl.GraphicContext, x, y float64, packet *string) {
	gc.SetFontData(draw2d.FontData{Name: "gomono"})
	gc.SetFillColor(image.White)
	gc.SetFontSize(8)
	gc.FillStringAt(*packet, x, y)
}
