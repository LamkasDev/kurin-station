package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/mathutils"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/arl/math32"
)

type KurinObjectLatheData struct {
	Energy    uint32
	MaxEnergy uint32
	Order     *KurinLatheOrder
}

type KurinLatheOrder struct {
	ItemType   string
	Energy     uint32
	TicksLeft  uint32
	TotalTicks uint32
}

func NewKurinObjectLathe(tile *KurinTile) *KurinObject {
	obj := NewKurinObjectRaw[*KurinObjectLatheData](tile, "lathe")
	obj.Health = 0
	obj.Process = func(object *KurinObject) {
		data := object.Data.(*KurinObjectLatheData)
		data.Energy = mathutils.MaxUint32(data.Energy+1, data.MaxEnergy)
		if data.Order == nil {
			return
		}

		data.Order.TicksLeft--
		if data.Order.TicksLeft == 0 {
			AddKurinItemToMapRaw(NewKurinItem(data.Order.ItemType, 1), &GameInstance.Map, &sdlutils.Transform{Position: common.GetPositionInDirectionFVCenter(object.Tile.Position, common.KurinDirectionEast), Rotation: 0})
			data.Order = nil
		}
	}
	obj.GetTexture = func(object *KurinObject) int {
		data := object.Data.(*KurinObjectLatheData)
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
	obj.OnInteraction = func(object *KurinObject, item *KurinItem) bool {
		OpenKurinDialog(&KurinDialogRequest{Type: "lathe", Data: &KurinDialogLatheData{Lathe: object}})
		return true
	}
	obj.Data = &KurinObjectLatheData{
		MaxEnergy: 8000,
	}

	return obj
}

func CreateNewOrderAtLathe(lathe *KurinObject, order *KurinLatheOrder) bool {
	data := lathe.Data.(*KurinObjectLatheData)
	if data.Energy < order.Energy {
		return false
	}
	data.Energy -= order.Energy
	data.Order = order

	return true
}

func NewKurinLatheOrder(itemType string) *KurinLatheOrder {
	return &KurinLatheOrder{
		ItemType:   itemType,
		Energy:     600,
		TicksLeft:  300,
		TotalTicks: 300,
	}
}
