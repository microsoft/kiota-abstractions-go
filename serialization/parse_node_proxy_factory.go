package serialization

type ParseNodeProxyFactory struct {
	factory        ParseNodeFactory
	onBeforeAction ParsableAction
	onAfterAction  ParsableAction
}

// NewParseNodeProxyFactory constructs a new instance of ParseNodeProxyFactory
func NewParseNodeProxyFactory(
	factory ParseNodeFactory,
	onBeforeAction ParsableAction,
	onAfterAction ParsableAction,
) *ParseNodeProxyFactory {
	return &ParseNodeProxyFactory{
		factory:        factory,
		onBeforeAction: onBeforeAction,
		onAfterAction:  onAfterAction,
	}
}

func (p *ParseNodeProxyFactory) GetValidContentType() (string, error) {
	return p.factory.GetValidContentType()
}

func (p *ParseNodeProxyFactory) GetRootParseNode(contentType string, content []byte) (ParseNode, error) {
	node, err := p.factory.GetRootParseNode(contentType, content)
	if err != nil {
		return nil, err
	}

	originalBefore := node.GetOnBeforeAssignFieldValues()
	err = node.SetOnBeforeAssignFieldValues(func(parsable Parsable) {
		if parsable != nil {
			p.onBeforeAction(parsable)
		}
		if originalBefore != nil {
			originalBefore(parsable)
		}
	})
	if err != nil {
		return nil, err
	}

	originalAfter := node.GetOnAfterAssignFieldValues()
	err = node.SetOnAfterAssignFieldValues(func(parsable Parsable) {
		if p != nil {
			p.onBeforeAction(parsable)
		}
		if originalAfter != nil {
			originalAfter(parsable)
		}
	})
	if err != nil {
		return nil, err
	}

	return node, nil
}
