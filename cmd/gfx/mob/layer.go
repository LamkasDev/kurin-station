package mob

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type RendererLayerMobData struct {
	Humanoids map[string]*HumanGraphicContainer
	Mobs      map[string]*MobGraphic
	ItemLayer *gfx.RendererLayer
}

func NewRendererLayerMob(itemLayer *gfx.RendererLayer) *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadRendererLayerMob,
		Render: RenderRendererLayerMob,
		Data: &RendererLayerMobData{
			Humanoids: map[string]*HumanGraphicContainer{},
			Mobs:      map[string]*MobGraphic{},
			ItemLayer: itemLayer,
		},
	}
}

func LoadRendererLayerMob(layer *gfx.RendererLayer) error {
	var err error
	if layer.Data.(*RendererLayerMobData).Humanoids["character"], err = NewHumanGraphicContainer("human"); err != nil {
		return err
	}
	if layer.Data.(*RendererLayerMobData).Mobs["cat"], err = NewMobGraphic("cat"); err != nil {
		return err
	}

	return nil
}

func RenderRendererLayerMob(layer *gfx.RendererLayer) error {
	for _, mob := range gameplay.GameInstance.Mobs {
		switch mob.Data.(type) {
		case *gameplay.MobCharacterData:
			RenderCharacter(layer, mob)
		default:
			RenderMob(layer, mob)
		}
	}

	return nil
}
