package hud

import (
	"fmt"
	"math"

	"github.com/LamkasDev/kurin/cmd/common/mathutils"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/item"
	"github.com/veandco/go-sdl2/sdl"
)

type RendererLayerHUDData struct {
	Icons       map[string]*HUDGraphic
	ItemLayer   *gfx.RendererLayer
	HoveredItem *gameplay.Item
}

func NewRendererLayerHUD(itemLayer *gfx.RendererLayer) *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadRendererLayerHUD,
		Render: RenderRendererLayerHUD,
		Data: &RendererLayerHUDData{
			Icons:     map[string]*HUDGraphic{},
			ItemLayer: itemLayer,
		},
	}
}

func LoadRendererLayerHUD(layer *gfx.RendererLayer) error {
	textures := []string{"hand_l", "lhandactive", "hand_r", "rhandactive", "act_equip", "swap_1", "swap_2", "pda", "selector", "template", "template_active", "credit", "objective_window", "cat"}
	var err error
	for _, texture := range textures {
		if layer.Data.(*RendererLayerHUDData).Icons[texture], err = NewHUDGraphic(texture); err != nil {
			return err
		}
	}

	return nil
}

func RenderRendererLayerHUD(layer *gfx.RendererLayer) error {
	if gfx.RendererInstance.Context.CameraMode != gfx.RendererCameraModeCharacter {
		return nil
	}
	half := gfx.GetHalfWindowSize(&gfx.RendererInstance.Context)
	data := layer.Data.(*RendererLayerHUDData)
	itemData := data.ItemLayer.Data.(*item.RendererLayerItemData)

	sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["hand_l"].Texture, HUDElementHandLeft.GetPosition(gfx.RendererInstance.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
	if gameplay.GameInstance.SelectedCharacter.ActiveHand == gameplay.HandLeft {
		sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["lhandactive"].Texture, sdl.Point{X: half.X, Y: gfx.RendererInstance.Context.WindowSize.Y - 72}, sdl.FPoint{X: 2, Y: 2})
	}
	lhand := gameplay.GameInstance.SelectedCharacter.Inventory.Hands[gameplay.HandLeft]
	if lhand != nil {
		graphic := itemData.Items[lhand.Type]
		if graphic.Outline != nil && data.HoveredItem == lhand {
			sdlutils.RenderTexture(gfx.RendererInstance.Renderer, graphic.Outline, HUDElementHandLeft.GetPosition(gfx.RendererInstance.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
		}
		for _, i := range lhand.Template.GetTextures(lhand) {
			sdlutils.RenderTexture(gfx.RendererInstance.Renderer, graphic.Textures[i].Base, HUDElementHandLeft.GetPosition(gfx.RendererInstance.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
		}
	}

	sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["hand_r"].Texture, HUDElementHandRight.GetPosition(gfx.RendererInstance.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
	if gameplay.GameInstance.SelectedCharacter.ActiveHand == gameplay.HandRight {
		sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["rhandactive"].Texture, sdl.Point{X: half.X - 64, Y: gfx.RendererInstance.Context.WindowSize.Y - 72}, sdl.FPoint{X: 2, Y: 2})
	}
	rhand := gameplay.GameInstance.SelectedCharacter.Inventory.Hands[gameplay.HandRight]
	if rhand != nil {
		graphic := itemData.Items[rhand.Type]
		if graphic.Outline != nil && data.HoveredItem == rhand {
			sdlutils.RenderTexture(gfx.RendererInstance.Renderer, graphic.Outline, HUDElementHandRight.GetPosition(gfx.RendererInstance.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
		}
		for _, i := range rhand.Template.GetTextures(rhand) {
			sdlutils.RenderTexture(gfx.RendererInstance.Renderer, graphic.Textures[i].Base, HUDElementHandRight.GetPosition(gfx.RendererInstance.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
		}
	}

	sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["act_equip"].Texture, sdl.Point{X: half.X - 64, Y: gfx.RendererInstance.Context.WindowSize.Y - 136}, sdl.FPoint{X: 2, Y: 2})
	sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["swap_1"].Texture, sdl.Point{X: half.X - 64, Y: gfx.RendererInstance.Context.WindowSize.Y - 136}, sdl.FPoint{X: 2, Y: 2})
	sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["swap_2"].Texture, sdl.Point{X: half.X, Y: gfx.RendererInstance.Context.WindowSize.Y - 136}, sdl.FPoint{X: 2, Y: 2})

	sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["pda"].Texture, HUDElementPDA.GetPosition(gfx.RendererInstance.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
	if HUDElementPDA.Hovered {
		layer.Data.(*RendererLayerHUDData).Icons["selector"].Texture.Texture.SetColorMod(0, 255, 0)
		sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["selector"].Texture, HUDElementPDA.GetPosition(gfx.RendererInstance.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
	}

	goals := HUDElementGoals.GetPosition(gfx.RendererInstance.Context.WindowSize)
	sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["objective_window"].Texture, goals, sdl.FPoint{X: 1, Y: 1})
	if len(gameplay.GameInstance.Narrator.Objectives) > 0 {
		objective := gameplay.GameInstance.Narrator.Objectives[0]
		text := objective.Text[0:mathutils.MinInt(int(math.Floor(float64(objective.Ticks)/6)), len(objective.Text))]
		sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["cat"].Texture, sdl.Point{X: goals.X + 6, Y: goals.Y}, sdl.FPoint{X: 2, Y: 2})
		sdlutils.RenderLabel(gfx.RendererInstance.Renderer, "hud.goals", gfx.RendererInstance.Fonts.Default, sdlutils.White, text, sdl.Point{X: goals.X + 72, Y: goals.Y + 16}, sdl.FPoint{X: 1, Y: 1})
		for i, requirement := range objective.Requirements {
			pos := sdl.Point{X: goals.X + 72, Y: goals.Y + 36 + (int32(i) * 16)}
			text := "??"
			switch data := requirement.Data.(type) {
			case *gameplay.ObjectiveRequirementDataCredits:
				text = fmt.Sprintf("Earn %d credits (%d/%d)", data.Count, gameplay.GameInstance.Credits, data.Count)
			case *gameplay.ObjectiveRequirementDataCreate:
				text = fmt.Sprintf("Create %d %s (%d/%d)", data.Count, data.ObjectType, data.Progress, data.Count)
			case *gameplay.ObjectiveRequirementDataDestroy:
				text = fmt.Sprintf("Destroy %d %s (%d/%d)", data.Count, data.ObjectType, data.Progress, data.Count)
			}

			sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.White)
			if requirement.Template.IsDone(requirement) {
				gfx.RendererInstance.Renderer.FillRect(&sdl.Rect{X: pos.X, Y: pos.Y + 2, W: 10, H: 10})
			} else {
				gfx.RendererInstance.Renderer.DrawRect(&sdl.Rect{X: pos.X, Y: pos.Y + 2, W: 10, H: 10})
			}
			sdlutils.RenderLabel(gfx.RendererInstance.Renderer, fmt.Sprintf("hud.goals.%d", i), gfx.RendererInstance.Fonts.DefaultSmall, sdlutils.White, text, sdl.Point{X: pos.X + 16, Y: pos.Y}, sdl.FPoint{X: 1, Y: 1})
		}
	}

	credit := HUDElementCredit.GetPosition(gfx.RendererInstance.Context.WindowSize)
	sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["credit"].Texture, credit, sdl.FPoint{X: 2, Y: 2})
	sdlutils.RenderLabel(gfx.RendererInstance.Renderer, "hud.credits", gfx.RendererInstance.Fonts.Default, sdlutils.White, fmt.Sprint(gameplay.GameInstance.Credits), sdl.Point{X: credit.X + 58, Y: credit.Y + 28}, sdl.FPoint{X: 1, Y: 1})

	return nil
}
