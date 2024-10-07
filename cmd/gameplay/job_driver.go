package gameplay

type JobDriver struct {
	Type         string
	Tile         *Tile
	Mob          *Mob
	Toils        []*JobToil
	ToilIndex    uint8
	TimeoutTicks uint64

	Template *JobDriverTemplate
	Data     interface{}
}
