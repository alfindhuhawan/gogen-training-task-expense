package getallexpense

import "your/path/project/domain_myexpense/model/repository"

// Outport of usecase
type Outport interface {
	repository.FindAllExpenseRepo
}
