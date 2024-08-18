package gameplay

func NewKurinJobToilRaw[D any](toilType string) *KurinJobToil {
	return &KurinJobToil{
		Start: func(driver *KurinJobDriver, toil *KurinJobToil) KurinJobToilStatus {
			return KurinJobToilStatusWorking
		},
		Process: func(driver *KurinJobDriver, toil *KurinJobToil) KurinJobToilStatus {
			return KurinJobToilStatusWorking
		},
		End: func(driver *KurinJobDriver, toil *KurinJobToil) {},
	}
}
