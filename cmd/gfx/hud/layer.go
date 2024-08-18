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

type KurinRendererLayerHUDData struct {
	Icons       map[string]*KurinHUDGraphic
	ItemLayer   *gfx.RendererLayer
	HoveredItem *gameplay.KurinItem
}

func NewKurinRendererLayerHUD(itemLayer *gfx.RendererLayer) *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadKurinRendererLayerHUD,
		Render: RenderKurinRendererLayerHUD,
		Data: &KurinRendererLayerHUDData{
			Icons:     map[string]*KurinHUDGraphic{},
			ItemLayer: itemLayer,
		},
	}
}

func LoadKurinRendererLayerHUD(layer *gfx.RendererLayer) error {
	textures := []string{"hand_l", "lhandactive", "hand_r", "rhandactive", "act_equip", "swap_1", "swap_2", "pda", "selector", "template", "template_active", "credit", "objective_window", "cat"}
	var err error
	for _, texture := range textures {
		if layer.Data.(*KurinRendererLayerHUDData).Icons[texture], err = NewKurinHUDGraphic(texture); err != nil {
			return err
		}
	}

	return nil
}

func RenderKurinRendererLayerHUD(layer *gfx.RendererLayer) error {
	if gfx.RendererInstance.Context.CameraMode != gfx.KurinRendererCameraModeCharacter {
		return nil
	}
	half := gfx.GetHalfWindowSize(&gfx.RendererInstance.Context)
	data := layer.Data.(*KurinRendererLayerHUDData)
	itemData := data.ItemLayer.Data.(*item.KurinRendererLayerItemData)

	sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["hand_l"].Texture, KurinHUDElementHandLeft.GetPosition(gfx.RendererInstance.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
	if gameplay.GameInstance.SelectedCharacter.ActiveHand == gameplay.KurinHandLeft {
		sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["lhandactive"].Texture, sdl.Point{X: half.X, Y: gfx.RendererInstance.Context.WindowSize.Y - 72}, sdl.FPoint{X: 2, Y: 2})
	}
	lhand := gameplay.GameInstance.SelectedCharacter.Inventory.Hands[gameplay.KurinHandLeft]
	if lhand != nil {
		graphic := itemData.Items[lhand.Type]
		if graphic.Outline != nil && data.HoveredItem == lhand {
			sdlutils.RenderTexture(gfx.RendererInstance.Renderer, graphic.Outline, KurinHUDElementHandLeft.GetPosition(gfx.RendererInstance.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
		}
		for _, i := range lhand.GetTextures(lhand) {
			sdlutils.RenderTexture(gfx.RendererInstance.Renderer, graphic.Textures[i].Base, KurinHUDElementHandLeft.GetPosition(gfx.RendererInstance.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
		}
	}

	sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["hand_r"].Texture, KurinHUDElementHandRight.GetPosition(gfx.RendererInstance.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
	if gameplay.GameInstance.SelectedCharacter.ActiveHand == gameplay.KurinHandRight {
		sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["rhandactive"].Texture, sdl.Point{X: half.X - 64, Y: gfx.RendererInstance.Context.WindowSize.Y - 72}, sdl.FPoint{X: 2, Y: 2})
	}
	rhand := gameplay.GameInstance.SelectedCharacter.Inventory.Hands[gameplay.KurinHandRight]
	if rhand != nil {
		graphic := itemData.Items[rhand.Type]
		if graphic.Outline != nil && data.HoveredItem == rhand {
			sdlutils.RenderTexture(gfx.RendererInstance.Renderer, graphic.Outline, KurinHUDElementHandRight.GetPosition(gfx.RendererInstance.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
		}
		for _, i := range rhand.GetTextures(rhand) {
			sdlutils.RenderTexture(gfx.RendererInstance.Renderer, graphic.Textures[i].Base, KurinHUDElementHandRight.GetPosition(gfx.RendererInstance.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
		}
	}

	sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["act_equip"].Texture, sdl.Point{X: half.X - 64, Y: gfx.RendererInstance.Context.WindowSize.Y - 136}, sdl.FPoint{X: 2, Y: 2})
	sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["swap_1"].Texture, sdl.Point{X: half.X - 64, Y: gfx.RendererInstance.Context.WindowSize.Y - 136}, sdl.FPoint{X: 2, Y: 2})
	sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["swap_2"].Texture, sdl.Point{X: half.X, Y: gfx.RendererInstance.Context.WindowSize.Y - 136}, sdl.FPoint{X: 2, Y: 2})

	sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["pda"].Texture, KurinHUDElementPDA.GetPosition(gfx.RendererInstance.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
	if KurinHUDElementPDA.Hovered {
		layer.Data.(*KurinRendererLayerHUDData).Icons["selector"].Texture.Texture.SetColorMod(0, 255, 0)
		sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["selector"].Texture, KurinHUDElementPDA.GetPosition(gfx.RendererInstance.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
	}

	goals := KurinHUDElementGoals.GetPosition(gfx.RendererInstance.Context.WindowSize)
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
			case gameplay.KurinNarratorObjectiveRequirementDataCredits:
				text = fmt.Sprintf("Earn %d credits (%d/%d)", data.Count, gameplay.GameInstance.Credits, data.Count)
			case gameplay.KurinNarratorObjectiveRequirementDataCreate:
				text = fmt.Sprintf("Create %d %s (%d/%d)", data.Count, data.ObjectType, data.Progress, data.Count)
			case gameplay.KurinNarratorObjectiveRequirementDataDestroy:
				text = fmt.Sprintf("Destroy %d %s (%d/%d)", data.Count, data.ObjectType, data.Progress, data.Count)
			}

			sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.White)
			if requirement.IsDone(requirement) {
				gfx.RendererInstance.Renderer.FillRect(&sdl.Rect{X: pos.X, Y: pos.Y + 2, W: 10, H: 10})
			} else {
				gfx.RendererInstance.Renderer.DrawRect(&sdl.Rect{X: pos.X, Y: pos.Y + 2, W: 10, H: 10})
			}
			sdlutils.RenderLabel(gfx.RendererInstance.Renderer, fmt.Sprintf("hud.goals.%d", i), gfx.RendererInstance.Fonts.DefaultSmall, sdlutils.White, text, sdl.Point{X: pos.X + 16, Y: pos.Y}, sdl.FPoint{X: 1, Y: 1})
		}
	}

	credit := KurinHUDElementCredit.GetPosition(gfx.RendererInstance.Context.WindowSize)
	sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["credit"].Texture, credit, sdl.FPoint{X: 2, Y: 2})
	sdlutils.RenderLabel(gfx.RendererInstance.Renderer, "hud.credits", gfx.RendererInstance.Fonts.Default, sdlutils.White, fmt.Sprint(gameplay.GameInstance.Credits), sdl.Point{X: credit.X + 58, Y: credit.Y + 28}, sdl.FPoint{X: 1, Y: 1})

	return nil
}
