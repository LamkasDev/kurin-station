package mob

import (
	"fmt"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type MobGraphic struct {
	Textures []*sdlutils.TextureWithSize
	Dead     *sdlutils.TextureWithSize
}

func NewMobGraphic(speciesId string) (*MobGraphic, error) {
	graphic := MobGraphic{
		Textures: make([]*sdlutils.TextureWithSize, 4),
	}

	var err error
	for i := range 4 {
		graphicPath := path.Join(constants.TexturesPath, "mob", speciesId, fmt.Sprintf("%s_%d.png", speciesId, i))
		if graphic.Textures[i], err = sdlutils.LoadTexture(gfx.RendererInstance.Renderer, graphicPath); err != nil {
			return &graphic, err
		}
	}

	deadGraphicPath := path.Join(constants.TexturesPath, "mob", speciesId, fmt.Sprintf("%s_dead.png", speciesId))
	graphic.Dead, _ = sdlutils.LoadTexture(gfx.RendererInstance.Renderer, deadGraphicPath)

	return &graphic, nil
}
