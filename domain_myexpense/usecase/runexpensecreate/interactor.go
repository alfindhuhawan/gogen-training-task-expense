package runexpensecreate

import (
	"context"
	"your/path/project/domain_myexpense/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type runExpenseCreateInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &runExpenseCreateInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *runExpenseCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	res := &InportResponse{}

	// code your usecase definition here ...

	expenseObj, err := entity.NewExpense(entity.ExpenseRequest{
		Value: req.Value,
		Desc:  req.Desc,
		Date:  req.Date,
	})
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveExpense(ctx, expenseObj)
	if err != nil {
		return nil, err
	}

	//!

	return res, nil
}
