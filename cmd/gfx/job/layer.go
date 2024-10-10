package job

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/structure"
	"github.com/LamkasDev/kurin/cmd/gfx/turf"
)

type RendererLayerJobData struct {
	TurfLayer   *gfx.RendererLayer
	ObjectLayer *gfx.RendererLayer
}

func NewRendererLayerJob(turfLayer *gfx.RendererLayer, objectLayer *gfx.RendererLayer) *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadRendererLayerJob,
		Render: RenderRendererLayerJob,
		Data: &RendererLayerJobData{
			TurfLayer:   turfLayer,
			ObjectLayer: objectLayer,
		},
	}
}

func LoadRendererLayerJob(layer *gfx.RendererLayer) error {
	return nil
}

func RenderRendererLayerJob(layer *gfx.RendererLayer) error {
	for _, job := range gameplay.GameInstance.JobController[gameplay.FactionPlayer].Jobs {
		if job.Tile != nil && job.Tile.Position.Z != gameplay.GameInstance.SelectedZ {
			continue
		}
		if err := RenderJob(gfx.RendererInstance, layer, job); err != nil {
			return err
		}
	}

	for _, mob := range gameplay.GameInstance.Map.Mobs {
		if mob.Faction != gameplay.FactionPlayer || mob.JobTracker.Job == nil {
			continue
		}
		if mob.JobTracker.Job.Tile != nil && mob.JobTracker.Job.Tile.Position.Z != gameplay.GameInstance.SelectedZ {
			continue
		}
		if err := RenderJob(gfx.RendererInstance, layer, mob.JobTracker.Job); err != nil {
			return err
		}
	}

	return nil
}

func RenderJob(renderer *gfx.Renderer, layer *gfx.RendererLayer, job *gameplay.JobDriver) error {
	color := sdlutils.White
	if job.TimeoutTicks > gameplay.GameInstance.Ticks {
		color = sdlutils.Red
	}

	jobData := layer.Data.(*RendererLayerJobData)
	switch data := job.Data.(type) {
	case *gameplay.JobDriverBuildData:
		if err := structure.RenderObjectBlueprint(jobData.ObjectLayer, &gameplay.Object{
			Tile: job.Tile,
			Type: data.ObjectType,
		}, color); err != nil {
			return err
		}
	case *gameplay.JobDriverBuildFloorData:
		if err := turf.RenderTileBlueprint(jobData.TurfLayer, &gameplay.Tile{
			Position: data.Position,
			Type:     data.TileType,
		}, color); err != nil {
			return err
		}
	}
	switch job.Type {
	case "destroy":
		object := gameplay.GetObjectAtTile(job.Tile)
		if object == nil {
			return nil
		}
		rect := sdlutils.ScaleRectCentered(structure.GetObjectRect(jobData.ObjectLayer, object), 0.8)
		sdlutils.RenderTextureRect(gfx.RendererInstance.Renderer, sdlutils.GetTextureFromContainer(gfx.RendererInstance.IconTextures, gfx.RendererInstance.Renderer, "delete"), rect)
	case "destroy_floor":
		rect := sdlutils.ScaleRectCentered(turf.GetTileRect(job.Tile), 0.8)
		sdlutils.RenderTextureRect(gfx.RendererInstance.Renderer, sdlutils.GetTextureFromContainer(gfx.RendererInstance.IconTextures, gfx.RendererInstance.Renderer, "delete_floor"), rect)
	}

	return nil
}
