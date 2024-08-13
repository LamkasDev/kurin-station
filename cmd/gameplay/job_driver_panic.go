package gameplay

func NewKurinJobDriverPanic() *KurinJobDriver {
	return &KurinJobDriver{
		Type: "panic",
		Toils: []*KurinJobToil{
			NewKurinJobToilPanic(),
		},
	}
}
