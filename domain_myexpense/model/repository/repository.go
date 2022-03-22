package repository

import (
	"context"
	"your/path/project/domain_myexpense/model/entity"
)

type SaveExpenseRepo interface {
	SaveExpense(ctx context.Context, obj *entity.Expense) error
}

type FindAllExpenseRepo interface {
	FindAllExpense(ctx context.Context, someID string) ([]*entity.Expense, error)
}
