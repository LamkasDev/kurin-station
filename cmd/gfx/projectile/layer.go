package projectile

import (
	"io/fs"
	"path"
	"path/filepath"
	"strings"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type RendererLayerProjectileData struct {
	Projectiles map[string]*sdlutils.TextureWithSize
}

func NewRendererLayerProjectile() *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadRendererLayerProjectile,
		Render: RenderRendererLayerProjectile,
		Data: &RendererLayerProjectileData{
			Projectiles: map[string]*sdlutils.TextureWithSize{},
		},
	}
}

func LoadRendererLayerProjectile(layer *gfx.RendererLayer) error {
	return filepath.WalkDir(path.Join(constants.TexturesPath, "projectiles"), func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if texture, err := sdlutils.LoadTexture(gfx.RendererInstance.Renderer, path); err == nil {
			layer.Data.(*RendererLayerProjectileData).Projectiles[strings.ReplaceAll(d.Name(), ".png", "")] = texture
		}

		return nil
	})
}

func RenderRendererLayerProjectile(layer *gfx.RendererLayer) error {
	for _, projectile := range gameplay.GameInstance.Map.Projectiles {
		if projectile.Position.Z != gameplay.GameInstance.SelectedZ {
			continue
		}
		RenderProjectile(layer, projectile)
	}

	return nil
}
