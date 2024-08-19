package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
)

type Runechat struct {
	Message string
	Ticks   uint32
	Data    interface{}
	Texture *sdlutils.TextureWithSize
}

func NewRunechat(message string) *Runechat {
	return &Runechat{
		Message: message,
		Ticks:   360,
	}
}

func DestroyRunechat(runechat *Runechat) {
	if runechat.Texture != nil {
		runechat.Texture.Texture.Destroy()
	}
}
