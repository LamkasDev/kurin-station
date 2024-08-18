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

func GetKurinCharacterRect(character *gameplay.KurinCharacter) sdl.Rect {
	position := sdlutils.AddFPoints(character.PositionRender, gameplay.GetAnimationOffset(character))
	return render.WorldToScreenRect(sdl.FRect{
		X: position.X, Y: position.Y,
		W: gameplay.KurinTileSizeF.X, H: gameplay.KurinTileSizeF.Y,
	})
}

func RenderKurinCharacter(layer *gfx.RendererLayer, character *gameplay.KurinCharacter) error {
	graphic := layer.Data.(*KurinRendererLayerCharacterData).Species[character.Species].Types[character.Type]
	rect := GetKurinCharacterRect(character)

	switch character.Direction {
	case common.KurinDirectionNorth:
		RenderKurinCharacterHand(layer, character, gameplay.KurinHandLeft, rect)
		RenderKurinCharacterHand(layer, character, gameplay.KurinHandRight, rect)
	case common.KurinDirectionEast:
		RenderKurinCharacterHand(layer, character, gameplay.KurinHandLeft, rect)
	case common.KurinDirectionSouth:
		break
	case common.KurinDirectionWest:
		RenderKurinCharacterHand(layer, character, gameplay.KurinHandRight, rect)
	}

	for _, part := range graphic.Template.Parts {
		offset := part.Offset
		if offset == nil {
			offset = &templates.KurinSpeciesTemplateBodypartOffset{}
		}

		texture := graphic.Textures[part.Id][character.Direction]
		prect := sdlutils.AddRectAndPoint(rect, sdl.Point{X: int32(float32(offset.X) * gfx.RendererInstance.Context.CameraZoom.X), Y: int32(float32(offset.Y) * gfx.RendererInstance.Context.CameraZoom.Y)})
		if err := gfx.RendererInstance.Renderer.Copy(texture.Texture, nil, &prect); err != nil {
			return err
		}
	}

	switch character.Direction {
	case common.KurinDirectionNorth:
		break
	case common.KurinDirectionEast:
		RenderKurinCharacterHand(layer, character, gameplay.KurinHandRight, rect)
	case common.KurinDirectionSouth:
		RenderKurinCharacterHand(layer, character, gameplay.KurinHandLeft, rect)
		RenderKurinCharacterHand(layer, character, gameplay.KurinHandRight, rect)
	case common.KurinDirectionWest:
		RenderKurinCharacterHand(layer, character, gameplay.KurinHandLeft, rect)
	}

	return nil
}

func RenderKurinCharacterHand(layer *gfx.RendererLayer, character *gameplay.KurinCharacter, hand gameplay.KurinHand, rect sdl.Rect) error {
	handItem := character.Inventory.Hands[hand]
	if handItem != nil {
		graphic := layer.Data.(*KurinRendererLayerCharacterData).ItemLayer.Data.(*item.KurinRendererLayerItemData).Items[handItem.Type]
		if graphic.Template.Hand == nil || !*graphic.Template.Hand {
			return nil
		}

		graphicDirections := graphic.Hands[hand][handItem.GetTextureHand(handItem)]
		if err := gfx.RendererInstance.Renderer.Copy(graphicDirections[character.Direction].Texture, nil, &rect); err != nil {
			return err
		}
	}

	return nil
}
