package sqlike

func (session *basicSession) Explain() ExplainSelectBranchStep {
	return &ExplainSelectBranchStepImpl{
		parent: &ExplainSelectBranchStepImpl{
			parent: &SelectExplainStep{
				parent: session.NewRootStep(),
			},
		},
	}
}

func (session *basicSession) SelectOne() SelectOneBranchStep {
	return &SelectOneBranchStepImpl{
		parent: &InstantStep{
			parent:    session.NewRootStep(),
			statement: sqlDialect[session.dialect][StatementTypeSelectOne],
		},
	}
}

func (session *basicSession) Select(columns ...Column) SelectColumnBranchStep {
	return &SelectColumnBranchStepImpl{
		parent: &SelectColumnStep{
			parent:  session.NewRootStep(),
			columns: columns,
		},
	}
}
