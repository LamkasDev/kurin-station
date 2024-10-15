package gameplay

type Bodypart struct {
	Type   string
	Points uint16

	Template *BodypartTemplate
	Data     interface{}
}

func GetBodypartState(bodypart *Bodypart) uint8 {
	percentage := (float32(bodypart.Points) / float32(bodypart.Template.MaxPoints)) * 100
	if percentage == 100 {
		return 0
	} else if percentage >= 80 {
		return 1
	} else if percentage >= 60 {
		return 2
	} else if percentage >= 40 {
		return 3
	} else if percentage >= 20 {
		return 4
	}

	return 5
}

func HitBodypart(bodypart *Bodypart, points uint16) {
	bodypart.Points = max(0, bodypart.Points-points)
}
