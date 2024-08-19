package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/mathutils"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/arl/math32"
)

type ObjectLatheData struct {
	Energy    uint32
	MaxEnergy uint32
	Order     *LatheOrder
}

func NewObjectTemplateLathe() *ObjectTemplate {
	template := NewObjectTemplate[*ObjectLatheData]("lathe", false)
	template.Process = func(object *Object) {
		data := object.Data.(*ObjectLatheData)
		data.Energy = mathutils.MaxUint32(data.Energy+1, data.MaxEnergy)
	}
	template.GetTexture = func(object *Object) int {
		data := object.Data.(*ObjectLatheData)
		if data.Order == nil {
			return 0
		}
		if data.Order.TicksLeft > 90 {
			if int(math32.Floor(float32(data.Order.TicksLeft)/20))%2 == 0 {
				return 0
			}
			return 1
		}

		return int(math32.Floor((90 - float32(data.Order.TicksLeft)) / 10))
	}
	template.OnInteraction = func(object *Object, item *Item) bool {
		OpenDialog(&DialogRequest{Type: "lathe", Data: &DialogLatheData{Lathe: object}})
		return true
	}
	template.GetDefaultData = func() interface{} {
		return &ObjectLatheData{
			MaxEnergy: 8000,
		}
	}
	template.MaxHealth = 15

	return template
}

func CreateNewOrderAtLathe(lathe *Object, order *LatheOrder) bool {
	data := lathe.Data.(*ObjectLatheData)
	if data.Order != nil {
		return false
	}
	if data.Energy < order.Energy {
		return false
	}
	data.Energy -= order.Energy
	data.Order = order

	job := NewJobDriver("manufacture", lathe.Tile)
	job.Template.Initialize(job, nil)
	PushJobToController(GameInstance.JobController[FactionPlayer], job)

	return true
}

func ProgressOrderAtLathe(lathe *Object) bool {
	data := lathe.Data.(*ObjectLatheData)
	if data.Order == nil {
		return true
	}
	data.Order.TicksLeft--
	if data.Order.TicksLeft == 0 {
		AddItemToMapRaw(NewItem(data.Order.ItemType, 1), &GameInstance.Map, &sdlutils.Transform{Position: common.GetPositionInDirectionFVCenter(lathe.Tile.Position, common.DirectionEast), Rotation: 0})
		data.Order = nil
		return true
	}

	return false
}
