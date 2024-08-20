package transaction

type FallibleTx interface {
	addCompensationStep(compensationStep func() error)
	fail(err error) error
	isFailed() bool
	error() error
	finish()
}

type fallible struct {
	stepIndex     uint
	err           error
	compensations []func() error
}

func (tx *fallible) addCompensationStep(compensationStep func() error) {
	tx.stepIndex++
	tx.compensations = append(tx.compensations, compensationStep)
}

func (tx *fallible) fail(err error) error {
	tx.err = err
	stepsCount := len(tx.compensations)
	for i := stepsCount - 1; i >= 0; i-- {
		err := tx.compensations[i]()
		if err != nil {
			return &FailedAndRolledBackPartiallyError{
				ExecuteIndex:      tx.stepIndex,
				ExecuteError:      tx.err,
				CompensationIndex: uint(i),
				CompensationError: err,
			}
		}
	}
	return &FailedAndRolledBackCompletelyError{
		ExecuteIndex: tx.stepIndex,
		ExecuteError: tx.error(),
	}
}

func (tx *fallible) isFailed() bool {
	return tx.err != nil
}

func (tx *fallible) error() error {
	return tx.err
}

func (tx *fallible) finish() {
	// to prevent leaked transaction usage
	tx.err = &FinishedError{}
}

func ExecuteFallible[I, O any](tx FallibleTx, op Operation[I, O], input I) (O, error) {
	if tx.isFailed() {
		return *new(O), &CannotExecuteInFailedTransactionError{OriginalError: tx.error()}
	}

	output, err := op.Execute(input)
	if err != nil {
		return *new(O), tx.fail(err)
	}

	tx.addCompensationStep(func() error { return op.Compensate(input, output) })
	return output, nil
}

// Fallible starts fallible transaction execution.
// Inside f operations can be executed using ExecuteFallible.
// If any operation fails, all the already executed successful operation's compensation actions
// are executed in reverse order and the transaction returns with a failure.
func Fallible[T any](f func(tx FallibleTx) (T, error)) (T, error) {
	tx := &fallible{}
	return f(tx)
}
