package templates

import (
	"github.com/veandco/go-sdl2/sdl"
)

type AnimationTemplate struct {
	Id    string                  `json:"id"`
	Steps []AnimationTemplateStep `json:"steps"`
}

type AnimationTemplateStep struct {
	Ticks     int32      `json:"ticks"`
	Offset    sdl.FPoint `json:"offset"`
	Direction *bool      `json:"direction"`
}
