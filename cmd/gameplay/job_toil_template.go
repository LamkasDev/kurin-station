package gameplay

type JobToilTemplate struct {
	Type    string
	Start   JobToilStart
	Process JobToilProcess
	End     JobToilEnd
}

type (
	JobToilStart   func(driver *JobDriver, toil *JobToil) JobToilStatus
	JobToilProcess func(driver *JobDriver, toil *JobToil) JobToilStatus
	JobToilEnd     func(driver *JobDriver, toil *JobToil)
)

func NewJobToilTemplate[D any](toilType string) *JobToilTemplate {
	return &JobToilTemplate{
		Start: func(driver *JobDriver, toil *JobToil) JobToilStatus {
			return JobToilStatusWorking
		},
		Process: func(driver *JobDriver, toil *JobToil) JobToilStatus {
			return JobToilStatusWorking
		},
		End: func(driver *JobDriver, toil *JobToil) {},
	}
}
