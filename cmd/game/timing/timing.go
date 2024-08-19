package timing

type Timing struct {
	FrameTime float32
}

var TimingGlobal = NewTiming()

func NewTiming() Timing {
	return Timing{}
}
