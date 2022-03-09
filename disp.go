package gopotato

const (
	XRES = 64
	YRES = 32
)

type display [XRES][YRES]bool

var disp display

func initDisp() {
	disp = display{}
}

// draws the given sprite on the display, with the top left corner at the given origin
// returns whether any pixels were erased by the draw
func drawSprite(sprite []byte, originX, originY byte) bool {
	didErase := false
	for y, spriteByte := range sprite {
		for bitIdx := byte(0); bitIdx < 8; bitIdx++ {
			drawPixel := (0x01<<bitIdx)&spriteByte > 0

			isLit := disp[(originX+bitIdx)%XRES][(originY+byte(y))%YRES]
			if isLit && !drawPixel {
				didErase = true
			}
			disp[(originX+bitIdx)%XRES][(originY+byte(y))%YRES] = drawPixel
		}
	}
	return didErase
}
