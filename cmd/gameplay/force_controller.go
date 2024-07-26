package gameplay

type KurinForceController struct {
	Forces map[*KurinItem]*KurinForce
}

func NewKurinForceController() KurinForceController {
	return KurinForceController{
		Forces: map[*KurinItem]*KurinForce{},
	}
}
