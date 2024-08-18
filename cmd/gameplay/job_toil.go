package gameplay

type KurinJobToil struct {
	Type    string
	Started bool
	Ticks   int32

	Start   KurinJobToilStart
	Process KurinJobToilProcess
	End     KurinJobToilEnd
	Data    interface{}
}

type (
	KurinJobToilStart   func(driver *KurinJobDriver, toil *KurinJobToil) KurinJobToilStatus
	KurinJobToilProcess func(driver *KurinJobDriver, toil *KurinJobToil) KurinJobToilStatus
	KurinJobToilEnd     func(driver *KurinJobDriver, toil *KurinJobToil)
)

type KurinJobToilStatus uint8

const (
	KurinJobToilStatusFailed   = KurinJobToilStatus(0)
	KurinJobToilStatusWorking  = KurinJobToilStatus(1)
	KurinJobToilStatusComplete = KurinJobToilStatus(2)
)
