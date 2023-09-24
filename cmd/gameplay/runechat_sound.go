package gameplay

type KurinRunechatSound struct {
	Runechat *KurinRunechat
}

func NewKurinRunechatSound(runechat *KurinRunechat) *KurinRunechatSound {
	return &KurinRunechatSound{
		Runechat: runechat,
	}
}
