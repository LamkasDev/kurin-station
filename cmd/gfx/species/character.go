package species

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gameplay/templates"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/item"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetKurinCharacterRect(renderer *gfx.KurinRenderer, character *gameplay.KurinCharacter) sdl.Rect {
	position := sdlutils.AddFPoints(character.PositionRender, gameplay.GetAnimationOffset(character))
	return render.WorldToScreenRect(renderer, sdl.FRect{
		X: position.X, Y: position.Y,
		W: float32(gameplay.KurinTileSize.X), H: float32(gameplay.KurinTileSize.Y),
	})
}

func RenderKurinCharacter(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, character *gameplay.KurinCharacter) *error {
	graphic := layer.Data.(KurinRendererLayerCharacterData).Species[character.Species].Types[character.Type]
	rect := GetKurinCharacterRect(renderer, character)

	switch character.Direction {
	case gameplay.KurinDirectionNorth:
		RenderKurinCharacterHand(renderer, layer, character, gameplay.KurinHandLeft, rect)
		RenderKurinCharacterHand(renderer, layer, character, gameplay.KurinHandRight, rect)
	case gameplay.KurinDirectionEast:
		RenderKurinCharacterHand(renderer, layer, character, gameplay.KurinHandLeft, rect)
	case gameplay.KurinDirectionSouth:
		break
	case gameplay.KurinDirectionWest:
		RenderKurinCharacterHand(renderer, layer, character, gameplay.KurinHandRight, rect)
	}

	for _, part := range graphic.Template.Parts {
		offset := part.Offset
		if offset == nil {
			offset = &templates.KurinSpeciesTemplateBodypartOffset{}
		}

		texture := graphic.Textures[part.Id][character.Direction]
		prect := sdlutils.AddRectAndPoint(rect, sdl.Point{X: int32(float32(offset.X) * renderer.RendererContext.CameraZoom.X), Y: int32(float32(offset.Y) * renderer.RendererContext.CameraZoom.Y)})
		if err := renderer.Renderer.Copy(texture.Texture, nil, &prect); err != nil {
			return &err
		}
	}

	switch character.Direction {
	case gameplay.KurinDirectionNorth:
		break
	case gameplay.KurinDirectionEast:
		RenderKurinCharacterHand(renderer, layer, character, gameplay.KurinHandRight, rect)
	case gameplay.KurinDirectionSouth:
		RenderKurinCharacterHand(renderer, layer, character, gameplay.KurinHandLeft, rect)
		RenderKurinCharacterHand(renderer, layer, character, gameplay.KurinHandRight, rect)
	case gameplay.KurinDirectionWest:
		RenderKurinCharacterHand(renderer, layer, character, gameplay.KurinHandLeft, rect)
	}

	return nil
}

func RenderKurinCharacterHand(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, character *gameplay.KurinCharacter, hand gameplay.KurinHand, rect sdl.Rect) *error {
	handItem := character.Inventory.Hands[hand]
	if handItem != nil {
		handDirections, ok := layer.Data.(KurinRendererLayerCharacterData).ItemLayer.Data.(item.KurinRendererLayerItemData).Items[handItem.Type].Hands[hand]
		if ok {
			if err := renderer.Renderer.Copy(handDirections[character.Direction].Texture, nil, &rect); err != nil {
				return &err
			}
		}
	}

	return nil
}
