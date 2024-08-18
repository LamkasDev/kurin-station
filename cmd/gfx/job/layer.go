package job

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/structure"
)

type KurinRendererLayerJobData struct {
	ObjectLayer *gfx.RendererLayer
}

func NewKurinRendererLayerJob(objectLayer *gfx.RendererLayer) *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadKurinRendererLayerJob,
		Render: RenderKurinRendererLayerJob,
		Data: &KurinRendererLayerJobData{
			ObjectLayer: objectLayer,
		},
	}
}

func LoadKurinRendererLayerJob(layer *gfx.RendererLayer) error {
	return nil
}

func RenderKurinRendererLayerJob(layer *gfx.RendererLayer) error {
	for _, job := range gameplay.GameInstance.JobController.Jobs {
		if err := RenderKurinJob(gfx.RendererInstance, layer, job); err != nil {
			return err
		}
	}

	for _, character := range gameplay.GameInstance.Characters {
		if character.JobTracker.Job == nil {
			continue
		}
		if err := RenderKurinJob(gfx.RendererInstance, layer, character.JobTracker.Job); err != nil {
			return err
		}
	}

	return nil
}

func RenderKurinJob(renderer *gfx.KurinRenderer, layer *gfx.RendererLayer, job *gameplay.KurinJobDriver) error {
	data := layer.Data.(*KurinRendererLayerJobData)
	switch val := job.Data.(type) {
	case *gameplay.KurinJobDriverBuildData:
		color := sdlutils.White
		if job.TimeoutTicks > gameplay.GameInstance.Ticks {
			color = sdlutils.Red
		}
		if err := structure.RenderKurinObjectBlueprint(data.ObjectLayer, &gameplay.KurinObject{
			Tile: job.Tile,
			Type: val.Prefab,
		}, color); err != nil {
			return err
		}
	}

	return nil
}
