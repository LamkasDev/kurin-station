package gameplay

type ReservationController struct{}

func NewReservationController() ReservationController {
	return ReservationController{}
}

func ReserveItem(item *Item) {
	item.Reserved = true
}

func UnreserveItem(item *Item) {
	item.Reserved = false
}
