package serialization

import "github.com/LamkasDev/kurin/cmd/gameplay"

type HealthData struct {
	Dead bool
}

func EncodeHealthData(mob *gameplay.Mob) HealthData {
	return HealthData{
		Dead: mob.Health.Dead,
	}
}

func DecodeHealthData(data HealthData, mob *gameplay.Mob) {
	mob.Health = gameplay.NewHealth()
	mob.Health.Dead = data.Dead
}
