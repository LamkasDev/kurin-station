package main

import (
	"os"
	"runtime/pprof"
	"time"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/game/life"
	"github.com/LamkasDev/kurin/cmd/game/timing"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	var err error
	if err = constants.LoadConstants(); err != nil {
		panic(err)
	}
	StartProfiling()
	if err = life.InitializeSystems(); err != nil {
		panic(err)
	}

	fps := 60
	tickMs := 1000 / float32(fps)
	tick := time.Tick(time.Duration(tickMs) * time.Millisecond)
	timing.TimingGlobal.FrameTime = tickMs

	for {
		<-tick
		if err = life.RunSystems(); err != nil {
			break
		}
		timing.TimingGlobal.FrameTime = tickMs
	}

	if event.EventManagerInstance.Close {
		StopProfiling()
		if err = life.FreeSystems(); err != nil {
			panic(err)
		}
		sdl.Quit()
		return
	}

	if err != nil {
		panic(err)
	}
}

func StartProfiling() {
	cpuProfileFile, err := os.Create(constants.ApplicationProfileCpu)
	if err != nil {
		panic(err)
	}

	if err := pprof.StartCPUProfile(cpuProfileFile); err != nil {
		panic(err)
	}
}

func StopProfiling() {
	pprof.StopCPUProfile()

	heapProfileFile, err := os.Create(constants.ApplicationProfileHeap)
	if err != nil {
		panic(err)
	}

	if err := pprof.WriteHeapProfile(heapProfileFile); err != nil {
		panic(err)
	}
}
