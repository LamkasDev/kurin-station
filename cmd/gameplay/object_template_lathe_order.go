package gameplay

type LatheOrder struct {
	ItemType   string
	ItemCount  uint16
	Energy     uint32
	TicksLeft  uint32
	TotalTicks uint32
}

func NewLatheOrder(itemType string, itemCount uint16) *LatheOrder {
	order := &LatheOrder{
		ItemType:   itemType,
		ItemCount:  itemCount,
		Energy:     600,
		TicksLeft:  300,
		TotalTicks: 300,
	}
	order.Energy *= uint32(itemCount)
	order.TicksLeft *= uint32(itemCount)
	order.TotalTicks *= uint32(itemCount)

	return order
}
