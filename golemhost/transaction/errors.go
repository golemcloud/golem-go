package transaction

import "fmt"

type FailedAndRolledBackPartiallyError struct {
	StepIndex         uint
	StepError         error
	CompensationIndex uint
	CompensationError error
}

func (e *FailedAndRolledBackPartiallyError) Error() string {
	return fmt.Sprintf(
		"fallible transaction failed and rolled back partially, step (%d) error: %s, compensation (%d) error: %s",
		e.StepIndex,
		e.StepError.Error(),
		e.CompensationIndex,
		e.CompensationError.Error(),
	)
}

func (e *FailedAndRolledBackPartiallyError) Unwrap() []error {
	return []error{e.StepError, e.CompensationError}
}

type FailedAndRolledBackCompletelyError struct {
	StepIndex uint
	StepError error
}

func (e *FailedAndRolledBackCompletelyError) Error() string {
	return fmt.Sprintf(
		"fallible transaction failed and rolled back completely, step (%d) error: %s",
		e.StepIndex,
		e.StepError.Error(),
	)
}

func (e *FailedAndRolledBackCompletelyError) Unwrap() error {
	return e.StepError
}

type CannotExecuteStepInFailedTransactionError struct {
	OriginalError error
}

func (e *CannotExecuteStepInFailedTransactionError) Error() string {
	return fmt.Sprintf(
		"cannot execute step in failed transaction, original error: %s",
		e.OriginalError.Error(),
	)
}

func (e *CannotExecuteStepInFailedTransactionError) Unwrap() error {
	return e.OriginalError
}

type FinishedError struct{}

func (e *FinishedError) Error() string {
	return "transaction finished"
}
