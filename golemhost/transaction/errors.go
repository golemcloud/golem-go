package transaction

import "fmt"

type FailedAndRolledBackPartiallyError struct {
	ExecuteIndex      uint
	ExecuteError      error
	CompensationIndex uint
	CompensationError error
}

func (e *FailedAndRolledBackPartiallyError) Error() string {
	return fmt.Sprintf(
		"fallible transaction failed and rolled back partially, execute (%d) error: %s, compensation (%d) error: %s",
		e.ExecuteIndex,
		e.ExecuteError.Error(),
		e.CompensationIndex,
		e.CompensationError.Error(),
	)
}

func (e *FailedAndRolledBackPartiallyError) Unwrap() []error {
	return []error{e.ExecuteError, e.CompensationError}
}

type FailedAndRolledBackCompletelyError struct {
	ExecuteIndex uint
	ExecuteError error
}

func (e *FailedAndRolledBackCompletelyError) Error() string {
	return fmt.Sprintf(
		"fallible transaction failed and rolled back completely, execute (%d) error: %s",
		e.ExecuteIndex,
		e.ExecuteError.Error(),
	)
}

func (e *FailedAndRolledBackCompletelyError) Unwrap() error {
	return e.ExecuteError
}

type CannotExecuteInFailedTransactionError struct {
	OriginalError error
}

func (e *CannotExecuteInFailedTransactionError) Error() string {
	return fmt.Sprintf(
		"cannot execute in failed transaction, original error: %s",
		e.OriginalError.Error(),
	)
}

func (e *CannotExecuteInFailedTransactionError) Unwrap() error {
	return e.OriginalError
}

type FinishedError struct{}

func (e *FinishedError) Error() string {
	return "transaction finished"
}
