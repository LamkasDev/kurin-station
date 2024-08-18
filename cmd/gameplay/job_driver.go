package gameplay

type KurinJobDriver struct {
	Type         string
	Tile         *KurinTile
	Character    *KurinCharacter
	Toils        []*KurinJobToil
	ToilIndex    uint8
	TimeoutTicks uint64

	Initialize KurinJobDriverInitialize
	EncodeData KurinJobDriverEncodeData
	DecodeData KurinJobDriverDecodeData
	Data       interface{}
}

type (
	KurinJobDriverInitialize func(job *KurinJobDriver, data interface{})
	KurinJobDriverEncodeData func(job *KurinJobDriver) []byte
	KurinJobDriverDecodeData func(job *KurinJobDriver, data []byte)
)
