package sdlutils

import (
	"fmt"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/veandco/go-sdl2/sdl"
)

type TextureContainer struct {
	Path     string
	Textures map[string]*TextureWithSize
}

func NewTextureContainer(rootPath string) *TextureContainer {
	return &TextureContainer{
		Path:     path.Join(constants.TexturesPath, rootPath),
		Textures: map[string]*TextureWithSize{},
	}
}

func GetTextureFromContainer(container *TextureContainer, renderer *sdl.Renderer, id string) *TextureWithSize {
	texture, ok := container.Textures[id]
	if !ok {
		var err error
		texturePath := path.Join(container.Path, fmt.Sprintf("%s.png", id))
		if container.Textures[id], err = LoadTexture(renderer, texturePath); err != nil {
			panic(err)
		}
		texture = container.Textures[id]
	}

	return texture
}
