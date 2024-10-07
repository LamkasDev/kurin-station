package mob

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/LamkasDev/kurin/cmd/gameplay/templates"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/item"
	"github.com/veandco/go-sdl2/sdl"
)

func RenderCharacter(layer *gfx.RendererLayer, mob *gameplay.Mob) error {
	graphic := layer.Data.(*RendererLayerMobData).Humanoids[mob.Type].Genders[mob.Gender]
	rect := GetMobRect(mob)

	switch mob.Direction {
	case common.DirectionNorth:
		RenderCharacterHand(layer, mob, gameplay.HandLeft, rect)
		RenderCharacterHand(layer, mob, gameplay.HandRight, rect)
	case common.DirectionEast:
		RenderCharacterHand(layer, mob, gameplay.HandLeft, rect)
	case common.DirectionSouth:
		break
	case common.DirectionWest:
		RenderCharacterHand(layer, mob, gameplay.HandRight, rect)
	}

	for _, part := range graphic.Template.Parts {
		offset := part.Offset
		if offset == nil {
			offset = &templates.SpeciesTemplateBodypartOffset{}
		}

		texture := graphic.Textures[part.Id][mob.Direction]
		prect := sdlutils.AddRectAndPoint(rect, sdl.Point{X: int32(float32(offset.X) * gfx.RendererInstance.Context.CameraZoom.X), Y: int32(float32(offset.Y) * gfx.RendererInstance.Context.CameraZoom.Y)})
		if err := gfx.RendererInstance.Renderer.Copy(texture.Texture, nil, &prect); err != nil {
			return err
		}
	}

	switch mob.Direction {
	case common.DirectionNorth:
		break
	case common.DirectionEast:
		RenderCharacterHand(layer, mob, gameplay.HandRight, rect)
	case common.DirectionSouth:
		RenderCharacterHand(layer, mob, gameplay.HandLeft, rect)
		RenderCharacterHand(layer, mob, gameplay.HandRight, rect)
	case common.DirectionWest:
		RenderCharacterHand(layer, mob, gameplay.HandLeft, rect)
	}

	return nil
}

func RenderCharacterHand(layer *gfx.RendererLayer, mob *gameplay.Mob, hand gameplay.Hand, rect sdl.Rect) error {
	handItem := mob.Data.(*gameplay.MobCharacterData).Inventory.Hands[hand]
	if handItem != nil {
		graphic := layer.Data.(*RendererLayerMobData).ItemLayer.Data.(*item.RendererLayerItemData).Items[handItem.Type]
		if graphic.Template.Hand == nil || !*graphic.Template.Hand {
			return nil
		}

		graphicDirections := graphic.Hands[hand][handItem.Template.GetTextureHand(handItem)]
		if err := gfx.RendererInstance.Renderer.Copy(graphicDirections[mob.Direction].Texture, nil, &rect); err != nil {
			return err
		}
	}

	return nil
}
