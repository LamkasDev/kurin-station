package gameplay

type JobDriver struct {
	Type         string
	Tile         *Tile
	Character    *Character
	Toils        []*JobToil
	ToilIndex    uint8
	TimeoutTicks uint64

	Template *JobDriverTemplate
	Data     interface{}
}
