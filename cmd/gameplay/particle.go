package gameplay

import (
	"math/rand"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type Particle struct {
	Type     string
	Scale    float32
	Rotation float64
	Color    sdl.Color
	Movement sdl.FPoint

	Position sdlutils.FVector3
	Index    uint8
	Ticks    uint32
}

func NewParticleIon(position sdlutils.FVector3, color sdl.Color) *Particle {
	return &Particle{
		Type:     "ion",
		Scale:    0.75,
		Rotation: 90,
		Color:    color,
		Movement: sdl.FPoint{X: -0.03},
		Position: position,
		Index:    0,
		Ticks:    60 + uint32(rand.Float32()*20),
	}
}

func NewParticleCross(position sdlutils.FVector3, scale float32, color sdl.Color) *Particle {
	return &Particle{
		Type:     "cross",
		Scale:    scale,
		Rotation: 0,
		Color:    color,
		Movement: sdl.FPoint{X: (rand.Float32() - 0.5) * 0.05, Y: (rand.Float32() - 0.5) * 0.05},
		Position: position,
		Index:    0,
		Ticks:    20,
	}
}

func ProcessParticle(particle *Particle) {
	particle.Ticks--
	particle.Position.Base.X += particle.Movement.X
	particle.Position.Base.Y += particle.Movement.Y
	switch particle.Type {
	case "ion":
		if particle.Ticks == 60 || particle.Ticks == 40 {
			particle.Index++
		}
	}
}
