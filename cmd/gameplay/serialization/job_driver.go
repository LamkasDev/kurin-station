package serialization

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/veandco/go-sdl2/sdl"
)

type JobDriverData struct {
	Type      string
	Position  *sdl.Point
	ToilIndex uint8
	Data      []byte
}

func EncodeJobDriver(jobDriver *gameplay.JobDriver) JobDriverData {
	data := JobDriverData{
		Type:      jobDriver.Type,
		ToilIndex: jobDriver.ToilIndex,
		Data:      jobDriver.Template.EncodeData(jobDriver),
	}
	if jobDriver.Tile != nil {
		data.Position = &jobDriver.Tile.Position.Base
	}

	return data
}

func DecodeJobDriver(data JobDriverData) *gameplay.JobDriver {
	jobDriver := gameplay.NewJobDriver(data.Type, nil)
	if data.Position != nil {
		jobDriver.Tile = gameplay.GameInstance.Map.Tiles[data.Position.X][data.Position.Y][0]
	}
	jobDriver.ToilIndex = data.ToilIndex
	jobDriver.Template.DecodeData(jobDriver, data.Data)

	return jobDriver
}
