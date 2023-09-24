package gameplay

import (
	"math/rand"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinItem struct {
	Type     string
	Position *sdlutils.FVector3
}

func NewKurinItem(itype string, position *sdlutils.FVector3) *KurinItem {
	return &KurinItem{
		Type:     itype,
		Position: position,
	}
}

func NewKurinItemRandom(itype string, kmap *KurinMap) *KurinItem {
	item := &KurinItem{
		Type: itype,
	}
	for {
		position := sdlutils.Vector3{Base: sdl.Point{X: int32(rand.Float32() * float32(kmap.Size.Base.X)), Y: int32(rand.Float32() * float32(kmap.Size.Base.Y))}, Z: 0}
		if CanEnterPosition(kmap, position) {
			item.Position = &sdlutils.FVector3{
				Base: sdl.FPoint{
					X: float32(position.Base.X) + 0.5,
					Y: float32(position.Base.Y) + 0.5,
				},
				Z: position.Z,
			}
			break
		}
	}

	return item
}
