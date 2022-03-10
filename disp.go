package gopotato

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	XRES = 64
	YRES = 32
)

type display struct {
	fb     framebuffer
	window *pixelgl.Window
}
type framebuffer [XRES][YRES]bool

var disp display

func initDisp() {
	disp = display{
		fb: framebuffer{},
	}
	cfg := pixelgl.WindowConfig{
		Title:  "gopotato",
		Bounds: pixel.R(0, 0, XRES, YRES),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.Clear(colornames.Black)
	disp.window = win
}

// draws the given sprite on the display, with the top left corner at the given origin
// returns whether any pixels were erased by the draw
func drawSprite(sprite []byte, originX, originY byte) bool {
	didErase := false
	for y, spriteByte := range sprite {
		for bitIdx := byte(0); bitIdx < 8; bitIdx++ {
			drawPixel := (0x01<<bitIdx)&spriteByte > 0
			isLit := disp.fb[(originX+bitIdx)%XRES][(originY+byte(y))%YRES]
			if isLit && !drawPixel {
				didErase = true
			}
			disp.fb[(originX+bitIdx)%XRES][(originY+byte(y))%YRES] = drawPixel
		}
	}
	return didErase
}
