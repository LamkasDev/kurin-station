package gameplay

type JobToil struct {
	Type    string
	Started bool
	Ticks   int32

	Template *JobToilTemplate
	Data     interface{}
}

type JobToilStatus uint8

const (
	JobToilStatusFailed   = JobToilStatus(0)
	JobToilStatusWorking  = JobToilStatus(1)
	JobToilStatusComplete = JobToilStatus(2)
)
