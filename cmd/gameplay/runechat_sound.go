package gameplay

type RunechatSound struct {
	Runechat *Runechat
}

func NewRunechatSound(runechat *Runechat) *RunechatSound {
	return &RunechatSound{
		Runechat: runechat,
	}
}
