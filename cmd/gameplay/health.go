package gameplay

import "robpike.io/filter"

type Health struct {
	Bodyparts []*Bodypart
	Dead      bool

	LastDamageTicks  uint64
	LastDamageSource interface{}
}

func NewHealth() *Health {
	return &Health{
		Bodyparts: []*Bodypart{
			NewBodypart("head"),
			NewBodypart("l_arm"),
			NewBodypart("r_arm"),
			NewBodypart("chest"),
			NewBodypart("l_leg"),
			NewBodypart("r_leg"),
		},
		Dead: false,
	}
}

func GetMaxHealthPoints(health *Health) uint16 {
	maxPoints := uint16(0)
	for _, bodypart := range health.Bodyparts {
		maxPoints += bodypart.Template.MaxPoints
	}

	return maxPoints
}

func GetHealthPoints(health *Health) uint16 {
	points := uint16(0)
	for _, bodypart := range health.Bodyparts {
		points += bodypart.Points
	}

	return points
}

func GetRandomUndamagedBodypart(health *Health) *Bodypart {
	bodyparts := filter.Choose(health.Bodyparts, func(bodypart *Bodypart) bool {
		return bodypart.Points != 0
	}).([]*Bodypart)
	if len(bodyparts) == 0 {
		return nil
	}

	return bodyparts[GameInstance.Map.Random.Intn(len(bodyparts))]
}

func GetRandomBodypart(health *Health) *Bodypart {
	return health.Bodyparts[GameInstance.Map.Random.Intn(len(health.Bodyparts))]
}
