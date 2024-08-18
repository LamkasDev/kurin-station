package render

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/arl/math32"
)

func GetThreeDots() string {
	return "..."[:int(math32.Floor(float32(gameplay.GameInstance.Ticks)/10))%4]
}
