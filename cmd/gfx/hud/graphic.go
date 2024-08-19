package hud

import (
	"fmt"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type HUDGraphic struct {
	Texture *sdlutils.TextureWithSize
}

func NewHUDGraphic(graphicId string) (*HUDGraphic, error) {
	graphic := HUDGraphic{}
	graphicDirectory := path.Join(constants.TexturesPath, "icons")

	var err error
	partPath := path.Join(graphicDirectory, fmt.Sprintf("%s.png", graphicId))
	if graphic.Texture, err = sdlutils.LoadTexture(gfx.RendererInstance.Renderer, partPath); err != nil {
		return &graphic, err
	}

	return &graphic, nil
}
