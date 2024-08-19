package gameplay

import "golang.org/x/exp/slices"

type RunechatController struct {
	Messages []*Runechat
	Sounds   []*RunechatSound
}

func NewRunechatController() RunechatController {
	return RunechatController{
		Messages: []*Runechat{},
		Sounds:   []*RunechatSound{},
	}
}

func CreateRunechatMessage(controller *RunechatController, runechat *Runechat) {
	controller.Messages = append(controller.Messages, runechat)
	controller.Sounds = append(controller.Sounds, NewRunechatSound(runechat))
}

func ProcessRunechat(controller *RunechatController, runechat *Runechat) {
	runechat.Ticks--
	if runechat.Ticks == 0 {
		DestroyRunechat(runechat)
		i := slices.Index(controller.Messages, runechat)
		controller.Messages = slices.Delete(controller.Messages, i, i+1)
	}
}
