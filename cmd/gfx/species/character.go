package species

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/LamkasDev/kurin/cmd/gameplay/templates"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/item"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetCharacterRect(character *gameplay.Character) sdl.Rect {
	position := sdlutils.AddFPoints(character.PositionRender, gameplay.GetAnimationOffset(character))
	return render.WorldToScreenRect(sdl.FRect{
		X: position.X, Y: position.Y,
		W: gameplay.TileSizeF.X, H: gameplay.TileSizeF.Y,
	})
}

func RenderCharacter(layer *gfx.RendererLayer, character *gameplay.Character) error {
	graphic := layer.Data.(*RendererLayerCharacterData).Species[character.Species].Types[character.Gender]
	rect := GetCharacterRect(character)

	switch character.Direction {
	case common.DirectionNorth:
		RenderCharacterHand(layer, character, gameplay.HandLeft, rect)
		RenderCharacterHand(layer, character, gameplay.HandRight, rect)
	case common.DirectionEast:
		RenderCharacterHand(layer, character, gameplay.HandLeft, rect)
	case common.DirectionSouth:
		break
	case common.DirectionWest:
		RenderCharacterHand(layer, character, gameplay.HandRight, rect)
	}

	for _, part := range graphic.Template.Parts {
		offset := part.Offset
		if offset == nil {
			offset = &templates.SpeciesTemplateBodypartOffset{}
		}

		texture := graphic.Textures[part.Id][character.Direction]
		prect := sdlutils.AddRectAndPoint(rect, sdl.Point{X: int32(float32(offset.X) * gfx.RendererInstance.Context.CameraZoom.X), Y: int32(float32(offset.Y) * gfx.RendererInstance.Context.CameraZoom.Y)})
		if err := gfx.RendererInstance.Renderer.Copy(texture.Texture, nil, &prect); err != nil {
			return err
		}
	}

	switch character.Direction {
	case common.DirectionNorth:
		break
	case common.DirectionEast:
		RenderCharacterHand(layer, character, gameplay.HandRight, rect)
	case common.DirectionSouth:
		RenderCharacterHand(layer, character, gameplay.HandLeft, rect)
		RenderCharacterHand(layer, character, gameplay.HandRight, rect)
	case common.DirectionWest:
		RenderCharacterHand(layer, character, gameplay.HandLeft, rect)
	}

	return nil
}

func RenderCharacterHand(layer *gfx.RendererLayer, character *gameplay.Character, hand gameplay.Hand, rect sdl.Rect) error {
	handItem := character.Inventory.Hands[hand]
	if handItem != nil {
		graphic := layer.Data.(*RendererLayerCharacterData).ItemLayer.Data.(*item.RendererLayerItemData).Items[handItem.Type]
		if graphic.Template.Hand == nil || !*graphic.Template.Hand {
			return nil
		}

		graphicDirections := graphic.Hands[hand][handItem.Template.GetTextureHand(handItem)]
		if err := gfx.RendererInstance.Renderer.Copy(graphicDirections[character.Direction].Texture, nil, &rect); err != nil {
			return err
		}
	}

	return nil
}
