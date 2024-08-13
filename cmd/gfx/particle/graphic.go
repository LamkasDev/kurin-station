package particle

import (
	"fmt"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinParticleGraphic struct {
	Texture sdlutils.TextureWithSize
}

func NewKurinParticleGraphic(renderer *gfx.KurinRenderer, particleId string) (*KurinParticleGraphic, error) {
	graphic := KurinParticleGraphic{}
	graphicDirectory := path.Join(constants.TexturesPath, "particles")

	var err error
	graphicPath := path.Join(graphicDirectory, fmt.Sprintf("%s_0.png", particleId))
	if graphic.Texture, err = sdlutils.LoadTexture(renderer.Renderer, graphicPath); err != nil {
		return &graphic, err
	}

	return &graphic, nil
}
