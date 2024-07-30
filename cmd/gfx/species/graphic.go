package species

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

type KurinSpeciesGraphicContainer struct {
	Types map[string]*KurinSpeciesGraphic
}

type KurinSpeciesGraphic struct {
	Template templates.KurinSpeciesTemplate
	Textures map[string][]sdlutils.TextureWithSize
}

func NewKurinSpeciesGraphicContainer(renderer *gfx.KurinRenderer, speciesId string) (*KurinSpeciesGraphicContainer, *error) {
	container := KurinSpeciesGraphicContainer{
		Types: map[string]*KurinSpeciesGraphic{},
	}

	var err *error
	if container.Types[gameplay.KurinDefaultType], err = NewKurinSpeciesGraphic(renderer, speciesId, gameplay.KurinDefaultType); err != nil {
		return &container, err
	}

	return &container, nil
}

func NewKurinSpeciesGraphic(renderer *gfx.KurinRenderer, speciesId string, speciesType string) (*KurinSpeciesGraphic, *error) {
	graphic := KurinSpeciesGraphic{
		Textures: map[string][]sdlutils.TextureWithSize{},
	}

	templateBytes, templateErr := os.ReadFile(path.Join(constants.DataPath, "templates", "species", fmt.Sprintf("%s.json", speciesId)))
	if templateErr != nil {
		return &graphic, &templateErr
	}
	if err := json.Unmarshal(templateBytes, &graphic.Template); err != nil {
		return &graphic, &err
	}

	var err *error
	for _, part := range graphic.Template.Parts {
		partDirectory := constants.TexturesPath
		if part.Path != nil {
			partDirectory = path.Join(partDirectory, *part.Path, part.Id)
		} else {
			partDirectory = path.Join(partDirectory, "bodyparts_greyscale", part.Id)
		}

		partFile := part.Id
		if part.Type != nil && *part.Type {
			partFile = fmt.Sprintf("%s_%s", partFile, speciesType)
		}

		graphic.Textures[part.Id] = make([]sdlutils.TextureWithSize, 4)
		for i := 0; i < 4; i++ {
			partPath := path.Join(partDirectory, fmt.Sprintf("%s_%d.png", partFile, i))
			if graphic.Textures[part.Id][i], err = sdlutils.LoadTexture(renderer.Renderer, partPath); err != nil {
				return &graphic, err
			}
		}
	}

	return &graphic, nil
}
