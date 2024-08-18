package particle

import (
	"fmt"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinParticleGraphic struct {
	Textures []*sdlutils.TextureWithSize
}

func NewKurinParticleGraphic(particleId string, states uint8) (*KurinParticleGraphic, error) {
	graphic := KurinParticleGraphic{
		Textures: make([]*sdlutils.TextureWithSize, states),
	}
	graphicDirectory := path.Join(constants.TexturesPath, "particles")

	var err error
	for i := range states {
		graphicPath := path.Join(graphicDirectory, fmt.Sprintf("%s_%d.png", particleId, i))
		if graphic.Textures[i], err = sdlutils.LoadTexture(gfx.RendererInstance.Renderer, graphicPath); err != nil {
			return &graphic, err
		}
	}

	return &graphic, nil
}
