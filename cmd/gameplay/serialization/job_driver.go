package serialization

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
)

type JobDriverData struct {
	Type      string
	Position  *sdlutils.Vector3
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
		data.Position = &jobDriver.Tile.Position
	}

	return data
}

func DecodeJobDriver(data JobDriverData) *gameplay.JobDriver {
	jobDriver := gameplay.NewJobDriver(data.Type, nil)
	if data.Position != nil {
		jobDriver.Tile = gameplay.GameInstance.Map.Tiles[data.Position.Base.X][data.Position.Base.Y][data.Position.Z]
	}
	jobDriver.ToilIndex = data.ToilIndex
	jobDriver.Template.DecodeData(jobDriver, data.Data)

	return jobDriver
}
