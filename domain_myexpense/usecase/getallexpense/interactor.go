package getallexpense

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type getAllExpenseInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &getAllExpenseInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *getAllExpenseInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	res := &InportResponse{}

	// code your usecase definition here ...

	expenseObjs, err := r.outport.FindAllExpense(ctx, "expenseID")
	if err != nil {
		return nil, err
	}

	for _, obj := range expenseObjs {
		fmt.Printf("%v\n", obj)
	}

	res.Items = expenseObjs

	return res, nil
}
