package gameplay

import "golang.org/x/exp/slices"

type KurinRunechatController struct {
	Messages []*KurinRunechat
	Sounds   []*KurinRunechatSound
}

func NewKurinRunechatController() KurinRunechatController {
	return KurinRunechatController{
		Messages: []*KurinRunechat{},
		Sounds:   []*KurinRunechatSound{},
	}
}

func CreateKurinRunechatMessage(controller *KurinRunechatController, runechat *KurinRunechat) {
	controller.Messages = append(controller.Messages, runechat)
	controller.Sounds = append(controller.Sounds, NewKurinRunechatSound(runechat))
}

func ProcessKurinRunechat(controller *KurinRunechatController, runechat *KurinRunechat) {
	runechat.Ticks--
	if runechat.Ticks == 0 {
		DestroyKurinRunechat(runechat)
		i := slices.Index(controller.Messages, runechat)
		controller.Messages = slices.Delete(controller.Messages, i, i+1)
	}
}
