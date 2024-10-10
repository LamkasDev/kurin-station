package gameplay

type Health struct {
	Dead   bool
	Points uint16
}

func NewHealth() Health {
	return Health{
		Dead:   false,
		Points: 3,
	}
}
