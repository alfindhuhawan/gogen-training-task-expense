package errorenum

import "your/path/project/shared/model/apperror"

const (
	SomethingError           apperror.ErrorType = "ER0000 something error"
	ValueMustGreaterThanZero apperror.ErrorType = "ER0001 value must greater than zero"
	DescMustNotEmpty         apperror.ErrorType = "ER0002 desc must not empty"
	DateMustNotEmpty         apperror.ErrorType = "ER0003 date must not empty"
	DateMustFormatMatch      apperror.ErrorType = "ER0004 format date must be YYYY-MM-DD"
)
