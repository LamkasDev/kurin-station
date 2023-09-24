package timing

type KurinTiming struct {
	FrameTime float32
}

var KurinTimingGlobal = NewKurinTiming()

func NewKurinTiming() KurinTiming {
	return KurinTiming{}
}
