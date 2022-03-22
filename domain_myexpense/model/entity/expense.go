package entity

import (
	"time"
	"your/path/project/domain_myexpense/model/errorenum"
)

type Expense struct {
	ID    uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Value int       `json:"value"`
	Desc  string    `json:"desc"`
	Date  time.Time `json:"date"`
}

type ExpenseRequest struct {
	Value int
	Desc  string
	Date  string
}

func NewExpense(req ExpenseRequest) (*Expense, error) {

	// value must greater than zero
	if req.Value <= 0 {
		return nil, errorenum.ValueMustGreaterThanZero
	}

	// desc cannot be empty
	if req.Desc == "" {
		return nil, errorenum.DescMustNotEmpty
	}

	// date cannot be empty
	if req.Date == "" {
		return nil, errorenum.DateMustNotEmpty
	}

	/* 	check if format date
	change format date if not error	*/
	dateFormat := "2006-01-02"
	dateInput, err := time.Parse(dateFormat, req.Date)
	if err != nil {
		return nil, errorenum.DateMustFormatMatch
	}

	var obj Expense

	// assign value here
	obj.Value = req.Value
	obj.Desc = req.Desc
	obj.Date = dateInput

	return &obj, nil
}

func (r *Expense) Validate() error {
	return nil
}
