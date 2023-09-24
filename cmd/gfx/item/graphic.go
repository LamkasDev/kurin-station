package item

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gameplay/templates"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinItemGraphic struct {
	Template templates.KurinItemTemplate
	Texture  sdlutils.TextureWithSizeAndSurface
	Outline  *sdlutils.TextureWithSize
	Hands    map[gameplay.KurinHand][]sdlutils.TextureWithSize
}

func NewKurinItemGraphic(renderer *gfx.KurinRenderer, itemId string) (*KurinItemGraphic, *error) {
	graphic := KurinItemGraphic{
		Hands: map[gameplay.KurinHand][]sdlutils.TextureWithSize{},
	}

	templateBytes, templateErr := os.ReadFile(path.Join(constants.DataPath, "templates", "items", fmt.Sprintf("%s.json", itemId)))
	if templateErr != nil {
		return &graphic, &templateErr
	}
	if err := json.Unmarshal(templateBytes, &graphic.Template); err != nil {
		return &graphic, &err
	}

	graphicDirectory := constants.TexturesPath
	if graphic.Template.Path != nil {
		graphicDirectory = path.Join(graphicDirectory, *graphic.Template.Path)
	} else {
		graphicDirectory = path.Join(graphicDirectory, "items")
	}

	var err *error
	graphicPath := path.Join(graphicDirectory, fmt.Sprintf("%s_0.png", itemId))
	if graphic.Texture, err = sdlutils.LoadTextureWithSurface(renderer.Renderer, graphicPath); err != nil {
		return &graphic, err
	}

	graphicOutlinePath := path.Join(graphicDirectory, fmt.Sprintf("%s_0_outline.png", itemId))
	if outline, err := sdlutils.LoadTexture(renderer.Renderer, graphicOutlinePath); err == nil {
		graphic.Outline = &outline
	}

	if graphic.Template.Hand == nil || *graphic.Template.Hand {
		graphic.Hands[gameplay.KurinHandLeft] = make([]sdlutils.TextureWithSize, 4)
		for i := 0; i < 4; i++ {
			handPath := path.Join(constants.TexturesPath, "melee_lefthand", fmt.Sprintf("%s_%d.png", "nullrod", i))
			if handTexture, err := sdlutils.LoadTexture(renderer.Renderer, handPath); err == nil {
				graphic.Hands[gameplay.KurinHandLeft][i] = handTexture
			}
		}

		graphic.Hands[gameplay.KurinHandRight] = make([]sdlutils.TextureWithSize, 4)
		for i := 0; i < 4; i++ {
			handPath := path.Join(constants.TexturesPath, "melee_righthand", fmt.Sprintf("%s_%d.png", "nullrod", i))
			if handTexture, err := sdlutils.LoadTexture(renderer.Renderer, handPath); err == nil {
				graphic.Hands[gameplay.KurinHandRight][i] = handTexture
			}
		}
	}

	return &graphic, nil
}
