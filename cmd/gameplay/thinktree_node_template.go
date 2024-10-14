package gameplay

type ThinktreeNodeTemplate struct {
	Type       string
	Initialize ThinktreeNodeInitialize
	Process    ThinktreeNodeProcess
}

type (
	ThinktreeNodeInitialize func(node *ThinktreeNode)
	ThinktreeNodeProcess    func(mob *Mob, node *ThinktreeNode) bool
)

func NewThinktreeNodeTemplate[D any](nodeType string) *ThinktreeNodeTemplate {
	return &ThinktreeNodeTemplate{
		Type:       nodeType,
		Initialize: func(node *ThinktreeNode) {},
		Process: func(mob *Mob, node *ThinktreeNode) bool {
			return false
		},
	}
}
