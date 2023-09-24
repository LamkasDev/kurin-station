package job

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/structure"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinRendererLayerJobData struct {
	ObjectLayer *gfx.KurinRendererLayer
}

func NewKurinRendererLayerJob(objectLayer *gfx.KurinRendererLayer) *gfx.KurinRendererLayer {
	return &gfx.KurinRendererLayer{
		Load:   LoadKurinRendererLayerJob,
		Render: RenderKurinRendererLayerJob,
		Data: KurinRendererLayerJobData{
			ObjectLayer: objectLayer,
		},
	}
}

func LoadKurinRendererLayerJob(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) *error {
	return nil
}

func RenderKurinRendererLayerJob(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, game *gameplay.KurinGame) *error {
	data := layer.Data.(KurinRendererLayerJobData)
	for _, ijob := range game.JobController.Jobs.Rep {
		if ijob == nil {
			continue
		}

		job := ijob.(*gameplay.KurinJobDriver)
		switch val := job.Data.(type) {
		case gameplay.KurinJobDriverBuildData:
			if err := structure.RenderKurinObjectBlueprint(renderer, data.ObjectLayer, job.Tile, val.Prefab, sdl.Color{R: 255, G: 255, B: 255}); err != nil {
				return err
			}
		}
	}

	for _, character := range game.Characters {
		if character.JobTracker.Job == nil {
			continue
		}

		switch val := character.JobTracker.Job.Data.(type) {
		case gameplay.KurinJobDriverBuildData:
			if err := structure.RenderKurinObjectBlueprint(renderer, data.ObjectLayer, character.JobTracker.Job.Tile, val.Prefab, sdl.Color{R: 255, G: 255, B: 255}); err != nil {
				return err
			}
		}
	}

	return nil
}
