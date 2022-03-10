package main

import (
	"fmt"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
	"time"
)

func main() {

	imd := imdraw.New(nil)
	frames := 0
	second := time.Tick(time.Second)
	frameNum := 0.
	for !disp.window.Closed() {

		disp.window.Clear(colornames.Black)
		imd.Clear()

		drawWindow(imd)

		imd.Draw(disp.window)
		disp.window.Update()
		frames++
		frameNum++
		select {
		case <-second:
			disp.window.SetTitle(fmt.Sprintf("%s | FPS: %d", "gopotato", frames))
			frames = 0
		default:
		}
	}
}
