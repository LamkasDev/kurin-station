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
	}
	if jobDriver.Tile != nil {
		data.Position = &jobDriver.Tile.Position
	}
	if jobDriver.Data != nil {
		data.Data = jobDriver.Template.EncodeData(jobDriver)
	}

	return data
}

func DecodeJobDriver(kmap *gameplay.Map, data JobDriverData) *gameplay.JobDriver {
	jobDriver := gameplay.NewJobDriver(data.Type, nil)
	jobDriver.ToilIndex = data.ToilIndex
	if data.Position != nil {
		jobDriver.Tile = kmap.Tiles[data.Position.Base.X][data.Position.Base.Y][data.Position.Z]
	}
	if data.Data != nil {
		jobDriver.Template.DecodeData(jobDriver, data.Data)
	}
	jobDriver.Template.Initialize(jobDriver)

	return jobDriver
}
