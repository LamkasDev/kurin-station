package life

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/event"
	eventActions "github.com/LamkasDev/kurin/cmd/event/actions"
	eventBase "github.com/LamkasDev/kurin/cmd/event/base"
	eventCamera "github.com/LamkasDev/kurin/cmd/event/camera"
	eventContext "github.com/LamkasDev/kurin/cmd/event/context"
	eventDebug "github.com/LamkasDev/kurin/cmd/event/debug"
	eventDialog "github.com/LamkasDev/kurin/cmd/event/dialog"
	eventForce "github.com/LamkasDev/kurin/cmd/event/force"
	eventHud "github.com/LamkasDev/kurin/cmd/event/hud"
	eventInteraction "github.com/LamkasDev/kurin/cmd/event/interaction"
	eventKeybinds "github.com/LamkasDev/kurin/cmd/event/keybinds"
	eventMovement "github.com/LamkasDev/kurin/cmd/event/movement"
	eventTool "github.com/LamkasDev/kurin/cmd/event/tool"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/actions"
	"github.com/LamkasDev/kurin/cmd/gfx/animation"
	"github.com/LamkasDev/kurin/cmd/gfx/background"
	"github.com/LamkasDev/kurin/cmd/gfx/context"
	"github.com/LamkasDev/kurin/cmd/gfx/debug"
	"github.com/LamkasDev/kurin/cmd/gfx/dialog"
	"github.com/LamkasDev/kurin/cmd/gfx/hud"
	"github.com/LamkasDev/kurin/cmd/gfx/item"
	"github.com/LamkasDev/kurin/cmd/gfx/job"
	"github.com/LamkasDev/kurin/cmd/gfx/mob"
	"github.com/LamkasDev/kurin/cmd/gfx/particle"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/LamkasDev/kurin/cmd/gfx/runechat"
	"github.com/LamkasDev/kurin/cmd/gfx/structure"
	"github.com/LamkasDev/kurin/cmd/gfx/tool"
	"github.com/LamkasDev/kurin/cmd/gfx/tooltip"
	"github.com/LamkasDev/kurin/cmd/gfx/turf"
	"github.com/LamkasDev/kurin/cmd/sound"
	"github.com/LamkasDev/kurin/cmd/sound/ambient"
	"github.com/LamkasDev/kurin/cmd/sound/music"
	"github.com/LamkasDev/kurin/cmd/sound/voice"
)

func InitializeSystems() error {
	gameplay.InitializeGame()

	var err error
	if err = gfx.InitializeRenderer(); err != nil {
		return err
	}

	turfLayer := turf.NewRendererLayerTile()
	objectLayer := structure.NewRendererLayerObject()
	itemLayer := item.NewRendererLayerItem()
	toolLayer := tool.NewRendererLayerTool(turfLayer, objectLayer)
	actionsLayer := actions.NewRendererLayerActions(turfLayer, objectLayer, itemLayer)
	hudLayer := hud.NewRendererLayerHUD(itemLayer)
	contextLayer := context.NewRendererLayerContext()
	debugLayer := debug.NewRendererLayerDebug()
	dialogLayer := dialog.NewRendererLayerDialog(itemLayer)

	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, background.NewRendererLayerBackground())
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, turfLayer)
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, objectLayer)
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, itemLayer)
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, job.NewRendererLayerJob(turfLayer, objectLayer))
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, mob.NewRendererLayerMob(itemLayer))
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, particle.NewRendererLayerParticle())
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, runechat.NewRendererLayerRunechat())
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, toolLayer)
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, actionsLayer)
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, tooltip.NewRendererLayerTooltip())
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, hudLayer)
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, contextLayer)
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, debugLayer)
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, animation.NewRendererLayerAnimation())
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, dialogLayer)
	if err := gfx.LoadRenderer(); err != nil {
		return err
	}

	if err = event.InitializeEventManager(); err != nil {
		return err
	}
	event.EventManagerInstance.Layers = append(event.EventManagerInstance.Layers, eventBase.NewEventLayerBase())
	event.EventManagerInstance.Layers = append(event.EventManagerInstance.Layers, eventDialog.NewEventLayerDialog(dialogLayer))
	event.EventManagerInstance.Layers = append(event.EventManagerInstance.Layers, eventContext.NewEventLayerContext(contextLayer))
	event.EventManagerInstance.Layers = append(event.EventManagerInstance.Layers, eventTool.NewEventLayerTool(toolLayer))
	event.EventManagerInstance.Layers = append(event.EventManagerInstance.Layers, eventActions.NewEventLayerActions(actionsLayer, toolLayer))
	event.EventManagerInstance.Layers = append(event.EventManagerInstance.Layers, eventKeybinds.NewEventLayerKeybinds())
	event.EventManagerInstance.Layers = append(event.EventManagerInstance.Layers, eventMovement.NewEventLayerMovement())
	event.EventManagerInstance.Layers = append(event.EventManagerInstance.Layers, eventForce.NewEventLayerForce())
	event.EventManagerInstance.Layers = append(event.EventManagerInstance.Layers, eventCamera.NewEventLayerCamera())
	event.EventManagerInstance.Layers = append(event.EventManagerInstance.Layers, eventHud.NewEventLayerHUD(hudLayer, itemLayer))
	event.EventManagerInstance.Layers = append(event.EventManagerInstance.Layers, eventInteraction.NewEventLayerInteraction(itemLayer))
	event.EventManagerInstance.Layers = append(event.EventManagerInstance.Layers, eventDebug.NewEventLayerDebug(debugLayer))
	if err := event.LoadEventManager(); err != nil {
		return err
	}

	if err = sound.InitializeSoundManager(); err != nil {
		return err
	}
	sound.SoundManagerInstance.Layers = append(sound.SoundManagerInstance.Layers, ambient.NewSoundLayerAmbient())
	sound.SoundManagerInstance.Layers = append(sound.SoundManagerInstance.Layers, voice.NewSoundLayerVoice())
	sound.SoundManagerInstance.Layers = append(sound.SoundManagerInstance.Layers, music.NewSoundLayerMusic())
	if err := sound.LoadSoundManager(); err != nil {
		return err
	}

	return nil
}

func RunSystems() error {
	if err := event.ProcessEventManager(); err != nil {
		return err
	}
	if err := sound.ProcessSoundManager(); err != nil {
		return err
	}

	gfx.RendererInstance.Context.CameraTileSizeF = render.GetCameraTileSize()
	gfx.RendererInstance.Context.CameraTileSize = sdlutils.FPointToPoint(gfx.RendererInstance.Context.CameraTileSizeF)
	gfx.RendererInstance.Context.CameraOffsetF = render.GetCameraOffset()
	gfx.RendererInstance.Context.CameraOffset = sdlutils.FPointToPoint(gfx.RendererInstance.Context.CameraOffsetF)
	if err := gfx.ClearRenderer(); err != nil {
		return err
	}
	if err := gfx.RenderRenderer(); err != nil {
		return err
	}
	gfx.PresentRenderer()

	return nil
}

func FreeSystems() error {
	if err := gfx.FreeRenderer(); err != nil {
		return err
	}
	if err := event.FreeEventManager(); err != nil {
		return err
	}

	return sound.FreeSoundManager()
}
