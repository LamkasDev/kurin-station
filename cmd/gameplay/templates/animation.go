package templates

import (
	"github.com/veandco/go-sdl2/sdl"
)

type KurinAnimationTemplate struct {
	Id    string                       `json:"id"`
	Steps []KurinAnimationTemplateStep `json:"steps"`
}

type KurinAnimationTemplateStep struct {
	Ticks     int32      `json:"ticks"`
	Offset    sdl.FPoint `json:"offset"`
	Direction *bool      `json:"direction"`
}
