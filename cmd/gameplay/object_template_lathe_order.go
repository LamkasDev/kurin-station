package gameplay

type LatheOrder struct {
	ItemType   string
	Energy     uint32
	TicksLeft  uint32
	TotalTicks uint32
}

func NewLatheOrder(itemType string) *LatheOrder {
	return &LatheOrder{
		ItemType:   itemType,
		Energy:     600,
		TicksLeft:  300,
		TotalTicks: 300,
	}
}
