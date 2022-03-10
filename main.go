package main

import (
	"fmt"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
)

func main() {
	//err := loadROM("pong.ch8")
	//err := loadROM("zero_demo.ch8")
	err := loadROM("ibm.ch8")
	//err := loadROM("invaders2.ch8")
	if err != nil {
		panic(err)
	}
	pixelgl.Run(run)
}

func run() {
	initDisp()
	go timerTick()
	go tick()
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
