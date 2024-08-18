package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

func NewKurinObjectBigThruster(tile *KurinTile) *KurinObject {
	obj := NewKurinObjectRaw[interface{}](tile, "big_thruster")
	obj.Health = 0
	obj.Process = func(object *KurinObject) {
		if GameInstance.Ticks%10 == 0 {
			pos := sdlutils.FVector3{Base: sdl.FPoint{X: float32(object.Tile.Position.Base.X) - 0.3, Y: float32(object.Tile.Position.Base.Y) + 0.5}, Z: 0}
			CreateKurinParticle(&GameInstance.ParticleController, NewKurinParticleIon(pos))
		}
	}

	return obj
}
