package gameplay

import (
	"math/rand"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinParticle struct {
	Type     string
	Position sdlutils.FVector3
	Scale float32
	Movement sdl.FPoint
	Color    sdl.Color
	Ticks    uint32
}

func NewKurinParticleCross(position sdlutils.FVector3, scale float32, color sdl.Color) *KurinParticle {
	return &KurinParticle{
		Type:     "cross",
		Position: position,
		Scale: scale,
		Movement: sdl.FPoint{X: (rand.Float32() - 0.5) * 0.05, Y: (rand.Float32() - 0.5) * 0.05},
		Color:    color,
		Ticks:    20,
	}
}

func ProcessKurinParticle(particle *KurinParticle) {
	particle.Ticks--
	particle.Position.Base.X += particle.Movement.X
	particle.Position.Base.Y += particle.Movement.Y
}
