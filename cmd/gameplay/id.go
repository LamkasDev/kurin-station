package gameplay

var NextId = uint32(0)

func GetNextId() uint32 {
	NextId++

	return NextId
}
