package actions

import (
	"sort"
	"strings"

	"github.com/LamkasDev/kurin/cmd/gfx/structure"
	"golang.org/x/exp/maps"
	"robpike.io/filter"
)

func GetMenuStructureGraphics(data *KurinRendererLayerActionsData) []*structure.KurinStructureGraphic {
	structures := maps.Values(data.ObjectLayer.Data.(structure.KurinRendererLayerObjectData).Structures)
	sort.Slice(structures, func(i, j int) bool {
		return structures[i].Template.Id < structures[j].Template.Id
	})
	newStructures := filter.Choose(structures, func(item *structure.KurinStructureGraphic) bool {
		return item.Template.Name != nil && strings.Contains(*item.Template.Name, data.Input)
	}).([]*structure.KurinStructureGraphic)

	return newStructures
}
