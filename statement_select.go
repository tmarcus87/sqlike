package sqlike

import "github.com/tmarcus87/sqlike/model"

func (session *basicSession) Explain() ExplainSelectBranchStep {
	return &ExplainSelectBranchStepImpl{
		parent: &SelectExplainStep{
			parent: session.NewRootStep(),
		},
	}
}

func (session *basicSession) SelectOne() SelectOneBranchStep {
	return &SelectOneBranchStepImpl{
		parent: &SelectOneStep{
			parent: session.NewRootStep(),
		},
	}
}

func (session *basicSession) Select(columns ...model.Column) SelectColumnBranchStep {
	return &SelectColumnBranchStepImpl{
		parent: &SelectColumnStep{
			parent:  session.NewRootStep(),
			columns: columns,
		},
	}
}
