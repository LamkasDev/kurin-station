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

	profileFile, perr := os.Create(constants.ApplicationProfile)
	if perr != nil {
		panic(perr)
	}
	if err := pprof.StartCPUProfile(profileFile); err != nil {
		panic(err)
	}

	instance, err := life.NewKurinInstance()
	if err != nil {
		panic(*err)
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

	if ferr := life.FreeKurinInstance(&instance); ferr != nil {
		panic(*ferr)
	}
	pprof.StopCPUProfile()
	sdl.Quit()
	if err != nil {
		panic(*err)
	}
}
