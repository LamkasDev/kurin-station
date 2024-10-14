package gameplay

type ThinktreeNode struct {
	Type string

	Data     interface{}
	Template *ThinktreeNodeTemplate
}
