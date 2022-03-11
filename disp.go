package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"sync"
)

const (
	XRES  = 64
	YRES  = 32
	SCALE = 10
)

type display struct {
	*sync.Mutex
	fb      framebuffer
	updated bool
	window  *pixelgl.Window
}
type framebuffer [XRES][YRES]bool

var disp display

func initDisp() {
	cfg := pixelgl.WindowConfig{
		Title:  "gopotato",
		Bounds: pixel.R(0, 0, XRES*SCALE, YRES*SCALE),
		VSync:  true,
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
	disp.Lock()
	defer disp.Unlock()
	disp.updated = true
	didErase := false
	for y, spriteByte := range sprite {
		for bitIdx := byte(0); bitIdx < 8; bitIdx++ {
			drawPixel := (0x80>>bitIdx)&spriteByte > 0
			// cheap short circuit around the xor before evaluating hundreds of modulo ops
			if !drawPixel {
				continue
			}

			isLit := disp.fb[(originX+bitIdx)%XRES][(originY+byte(y))%YRES]
			if isLit && drawPixel {
				didErase = true
			}
			// pixels are drawn via xor. with no xor logical operator, we must expand it
			disp.fb[(originX+bitIdx)%XRES][(originY+byte(y))%YRES] = (drawPixel || isLit) && !(drawPixel && isLit)
		}
	}
	return didErase
}

func drawWindow(imd *imdraw.IMDraw) {
	disp.Lock()
	defer disp.Unlock()
	defer disp.window.Update()
	if !disp.updated {
		return
	}

	disp.window.Clear(colornames.Black)
	imd.Clear()

	for rownum, row := range disp.fb {
		for colnum, pix := range row {
			if !pix { // only draw anything if the pixel is lit.
				continue
			}
			// origin according to Pixel is the lower left corner
			// the CHIP-8 and our framebuffer use the upper left corner
			imd.Color = colornames.White
			imd.Push(pixel.V(float64(rownum*SCALE), float64(YRES*SCALE-colnum*SCALE)),
				pixel.V(float64(rownum*SCALE+1*(SCALE-1)), float64(YRES*SCALE-colnum*SCALE+1*(SCALE-1))))
			imd.Rectangle(0.)
		}
	}

	imd.Draw(disp.window)

}
