package main

import (
	"fmt"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

const (
	CPU_PROFILE  = false
	MEM_PROFILE  = false
	DEBUG_OUTPUT = false
)

func main() {
	err := loadROM("chip8-roms/programs/IBM Logo.ch8")
	if err != nil {
		panic(err)
	}
	pixelgl.Run(run)
}

func run() {
	if CPU_PROFILE {
		f, err := os.Create("cpu.pprof")
		if err != nil {
			panic(err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()
	}

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

	if MEM_PROFILE {
		f, err := os.Create("mem.pprof")
		if err != nil {
			panic(err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			panic(err)
		}
	}
	initDisp()
}
