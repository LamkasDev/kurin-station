package gameplay

type KurinJobToil struct {
	Start   KurinJobToilStart
	Process KurinJobToilProcess
	Data    interface{}
	Ticks   int32
}

type KurinJobToilStart func(driver *KurinJobDriver, toil *KurinJobToil)

type KurinJobToilProcess func(driver *KurinJobDriver, toil *KurinJobToil) bool
