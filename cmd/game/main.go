package main

import (
	"os"
	"runtime/pprof"
	"time"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/game/life"
	"github.com/LamkasDev/kurin/cmd/game/timing"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := constants.LoadConstants(); err != nil {
		panic(err)
	}

	StartProfiling()

	instance, err := life.NewKurinInstance()
	if err != nil {
		panic(err)
	}

	fps := 60
	tickMs := 1000 / float32(fps)
	tick := time.Tick(time.Duration(tickMs) * time.Millisecond)
	timing.KurinTimingGlobal.FrameTime = tickMs

	for {
		<-tick

		if err = life.RunKurinInstance(&instance); err != nil {
			break
		}

		timing.KurinTimingGlobal.FrameTime = tickMs
	}

	if instance.EventManager.Close {
		StopProfiling()

		if err := life.FreeKurinInstance(&instance); err != nil {
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
