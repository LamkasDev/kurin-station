package gameplay

func NewKurinJobDriverPanic() *KurinJobDriver {
	job := NewKurinJobDriverRaw[interface{}]("panic")
	job.Toils = []*KurinJobToil{
		NewKurinJobToilPanic(),
	}

	return job
}
