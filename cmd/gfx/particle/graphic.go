package particle

import (
	"fmt"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type ParticleGraphic struct {
	Textures []*sdlutils.TextureWithSize
}

func NewParticleGraphic(particleId string, states uint8) (*ParticleGraphic, error) {
	graphic := ParticleGraphic{
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
