package main

import (
	"time"

	"github.com/LamkasDev/kitsune/cmd/browser/handler"
	"github.com/LamkasDev/kitsune/cmd/browser/life"
	"github.com/LamkasDev/kitsune/cmd/common/constants"
)

func main() {
	if err := constants.LoadConstants(); err != nil {
		panic(err)
	}

	instance, err := life.NewKitsuneInstance()
	if err != nil {
		panic(err)
	}

	tick := time.Tick(16 * time.Millisecond)
	for {
		if run, _ := handler.HandleEvents(instance); !run {
			break
		}
		select {
		case <-tick:
			handler.HandleEventsFrame(instance)
			life.RunKitsuneInstance(instance)
		}
	}
	life.FreeKitsuneInstance(instance)
}
