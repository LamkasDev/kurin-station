package serialization

import "github.com/LamkasDev/kurin/cmd/gameplay"

type HealthData struct {
	Dead   bool
	Points uint16
}

func EncodeHealthData(mob *gameplay.Mob) HealthData {
	return HealthData{
		Dead:   mob.Health.Dead,
		Points: mob.Health.Points,
	}
}

func DecodeHealthData(data HealthData, mob *gameplay.Mob) {
	mob.Health = gameplay.Health{
		Dead:   data.Dead,
		Points: data.Points,
	}
}
