package hud

import (
	"fmt"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinHUDGraphic struct {
	Texture sdlutils.TextureWithSize
}

func NewKurinHUDGraphic(renderer *gfx.KurinRenderer, hudId string) (*KurinHUDGraphic, error) {
	graphic := KurinHUDGraphic{}
	graphicDirectory := path.Join(constants.TexturesPath, "icons")

	var err error
	partPath := path.Join(graphicDirectory, fmt.Sprintf("%s_0.png", hudId))
	if graphic.Texture, err = sdlutils.LoadTexture(renderer.Renderer, partPath); err != nil {
		return &graphic, err
	}

	return &graphic, nil
}
