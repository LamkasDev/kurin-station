package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
)

type KurinRunechat struct {
	Message string
	Ticks   uint32
	Data    interface{}
	Texture *sdlutils.TextureWithSize
}

func NewKurinRunechat(message string) *KurinRunechat {
	return &KurinRunechat{
		Message: message,
		Ticks:   360,
	}
}

func DestroyKurinRunechat(runechat *KurinRunechat) {
	if runechat.Texture != nil {
		runechat.Texture.Texture.Destroy()
	}
}
