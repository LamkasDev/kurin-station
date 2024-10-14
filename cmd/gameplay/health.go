package gameplay

type Health struct {
	Dead   bool
	Points uint16

	LastDamageTicks  uint64
	LastDamageSource interface{}
}

func NewHealth() Health {
	return Health{
		Dead:   false,
		Points: 3,
	}
}
