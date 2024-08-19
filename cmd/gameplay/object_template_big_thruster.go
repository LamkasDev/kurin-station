package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

func NewObjectTemplateThruster(thrusterType string, particleColor sdl.Color) *ObjectTemplate {
	template := NewObjectTemplate[interface{}](thrusterType, false)
	template.Process = func(object *Object) {
		if GameInstance.Ticks%10 == 0 {
			pos := sdlutils.FVector3{Base: sdl.FPoint{X: float32(object.Tile.Position.Base.X) - 0.3, Y: float32(object.Tile.Position.Base.Y) + 0.5}, Z: 0}
			CreateParticle(&GameInstance.ParticleController, NewParticleIon(pos, particleColor))
		}
	}
	template.MaxHealth = 30

	return template
}
