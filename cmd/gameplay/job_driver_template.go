package gameplay

import "github.com/kelindar/binary"

type JobDriverTemplate struct {
	Type          string
	ReturnsOnFail bool
	Initialize    JobDriverInitialize
	EncodeData    JobDriverEncodeData
	DecodeData    JobDriverDecodeData
}

type (
	JobDriverInitialize func(job *JobDriver)
	JobDriverEncodeData func(job *JobDriver) []byte
	JobDriverDecodeData func(job *JobDriver, data []byte)
)

func NewJobDriverTemplate[D any](jobType string) *JobDriverTemplate {
	return &JobDriverTemplate{
		Type:          jobType,
		ReturnsOnFail: true,
		Initialize:    func(job *JobDriver) {},
		EncodeData: func(job *JobDriver) []byte {
			jobData := job.Data.(D)
			data, _ := binary.Marshal(&jobData)
			return data
		},
		DecodeData: func(job *JobDriver, data []byte) {
			var jobData D
			binary.Unmarshal(data, &jobData)
			job.Data = jobData
		},
	}
}
