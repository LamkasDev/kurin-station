package gameplay

import "github.com/kelindar/binary"

func NewKurinJobDriverRaw[D any](jobType string) *KurinJobDriver {
	return &KurinJobDriver{
		Type:       jobType,
		Initialize: func(job *KurinJobDriver, data interface{}) {},
		EncodeData: func(job *KurinJobDriver) []byte {
			if job.Data == nil {
				return []byte{}
			}

			jobData := job.Data.(D)
			data, _ := binary.Marshal(&jobData)
			return data
		},
		DecodeData: func(job *KurinJobDriver, data []byte) {
			if len(data) == 0 {
				return
			}

			var jobData D
			binary.Unmarshal(data, &jobData)
			job.Initialize(job, jobData)
		},
	}
}

func NewKurinJobDriver(jobType string) *KurinJobDriver {
	switch jobType {
	case "build":
		return NewKurinJobDriverBuild()
	case "panic":
		return NewKurinJobDriverPanic()
	}

	return NewKurinJobDriverRaw[interface{}](jobType)
}
