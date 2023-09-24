package gameplay

import (
	"math/rand"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinParticle struct {
	Type     string
	Position sdlutils.FVector3
	Movement sdl.FPoint
	Color    sdl.Color
	Ticks    uint32
}

func NewKurinParticleCross(game *KurinGame, position sdlutils.FVector3) *KurinParticle {
	return &KurinParticle{
		Type:     "cross",
		Position: position,
		Movement: sdl.FPoint{X: (rand.Float32() - 0.5) * 0.05, Y: (rand.Float32() - 0.5) * 0.05},
		Color:    sdl.Color{R: 210, G: 210, B: 210},
		Ticks:    20,
	}
}

func ProcessKurinParticle(particle *KurinParticle) {
	particle.Ticks--
	particle.Position.Base.X += particle.Movement.X
	particle.Position.Base.Y += particle.Movement.Y
}
