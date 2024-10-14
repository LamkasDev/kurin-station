package gameplay

var ThinktreeNodeContainer = map[string]*ThinktreeNodeTemplate{}

func RegisterThinktreeNodes() {
	ThinktreeNodeContainer["attack"] = NewThinktreeNodeTemplateAttack()
	ThinktreeNodeContainer["revenge"] = NewThinktreeNodeTemplateRevenge()
	ThinktreeNodeContainer["panic"] = NewThinktreeNodeTemplatePanic()
	ThinktreeNodeContainer["wander"] = NewThinktreeNodeTemplateCreate()
}

func NewThinktreeNode(nodeType string) *ThinktreeNode {
	node := &ThinktreeNode{
		Type:     nodeType,
		Template: ThinktreeNodeContainer[nodeType],
	}
	node.Template.Initialize(node)

	return node
}
