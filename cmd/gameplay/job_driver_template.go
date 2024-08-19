package gameplay

import "github.com/kelindar/binary"

type JobDriverTemplate struct {
	Type       string
	Initialize JobDriverInitialize
	EncodeData JobDriverEncodeData
	DecodeData JobDriverDecodeData
}

type (
	JobDriverInitialize func(job *JobDriver, data interface{})
	JobDriverEncodeData func(job *JobDriver) []byte
	JobDriverDecodeData func(job *JobDriver, data []byte)
)

func NewJobDriverTemplate[D any](jobType string) *JobDriverTemplate {
	return &JobDriverTemplate{
		Type:       jobType,
		Initialize: func(job *JobDriver, data interface{}) {},
		EncodeData: func(job *JobDriver) []byte {
			if job.Data == nil {
				return []byte{}
			}

			jobData := job.Data.(D)
			data, _ := binary.Marshal(&jobData)
			return data
		},
		DecodeData: func(job *JobDriver, data []byte) {
			if len(data) == 0 {
				return
			}

			var jobData D
			binary.Unmarshal(data, &jobData)
			job.Template.Initialize(job, jobData)
		},
	}
}
