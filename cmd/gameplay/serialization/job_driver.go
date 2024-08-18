package serialization

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinJobDriverData struct {
	Type      string
	Position  *sdl.Point
	ToilIndex uint8
	Data      []byte
}

func EncodeKurinJobDriver(jobDriver *gameplay.KurinJobDriver) KurinJobDriverData {
	return KurinJobDriverData{
		Type:      jobDriver.Type,
		Position:  &jobDriver.Tile.Position.Base,
		ToilIndex: jobDriver.ToilIndex,
		Data:      jobDriver.EncodeData(jobDriver),
	}
}

func DecodeKurinJobDriver(data KurinJobDriverData) *gameplay.KurinJobDriver {
	jobDriver := gameplay.NewKurinJobDriver(data.Type)
	if data.Position != nil {
		jobDriver.Tile = gameplay.GameInstance.Map.Tiles[data.Position.X][data.Position.Y][0]
	}
	jobDriver.ToilIndex = data.ToilIndex
	jobDriver.DecodeData(jobDriver, data.Data)

	return jobDriver
}
