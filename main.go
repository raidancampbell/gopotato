package main

import (
	"fmt"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
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
	for !disp.window.Closed() {

		drawWindow(imd)

		frames++
		select {
		case <-second:
			disp.window.SetTitle(fmt.Sprintf("%s | FPS: %d", "gopotato", frames))
			frames = 0
		default:
		}
	}
}
