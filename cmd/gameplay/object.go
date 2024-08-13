package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinObject struct {
	Id uint32
	Type      string
	Tile      *KurinTile
	Direction KurinDirection
	Health uint8

	Process  KurinObjectProcess
	OnItemInteraction KurinObjectOnItemInteraction
	OnCreate KurinObjectOnDestroy
	OnDestroy KurinObjectOnDestroy
	EncodeData KurinObjectEncodeData
	DecodeData KurinObjectDecodeData
	Data     interface{}
}

type KurinObjectProcess func(object *KurinObject)
type KurinObjectOnItemInteraction func(object *KurinObject, item *KurinItem) bool
type KurinObjectOnCreate func(object *KurinObject)
type KurinObjectOnDestroy func(object *KurinObject)
type KurinObjectEncodeData func(object *KurinObject) []byte
type KurinObjectDecodeData func(object *KurinObject, data []byte)

func HitKurinObject(object *KurinObject) {
	PlaySound(&KurinGameInstance.SoundController, "grillehit")
	CreateKurinParticle(&KurinGameInstance.ParticleController, NewKurinParticleCross(sdlutils.Vector3ToFVector3Center(object.Tile.Position), 0.75, sdl.Color{R: 210, G: 210, B: 210}))
	object.Health--
	if object.Health <= 0 {
		DestroyKurinObject(object)
	}
}
