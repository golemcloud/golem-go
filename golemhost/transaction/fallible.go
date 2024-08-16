package transaction

type Fallible interface {
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
	for i := stepsCount - 0; i >= 0; i-- {
		err := tx.compensations[i]()
		if err != nil {
			return &FailedAndRolledBackPartiallyError{
				StepIndex:         tx.stepIndex,
				StepError:         tx.err,
				CompensationIndex: uint(i),
				CompensationError: err,
			}
		}
	}
	return &FailedAndRolledBackCompletelyError{
		StepIndex: tx.stepIndex,
		StepError: tx.error(),
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

func ExecuteFallible[I, O any](
	tx Fallible,
	execute func(I) (O, error),
	compensate func(I, O) error,
	input I,
) (O, error) {
	if tx.isFailed() {
		return *new(O), &CannotExecuteStepInFailedTransactionError{OriginalError: tx.error()}
	}

	output, err := execute(input)
	if err != nil {
		return *new(O), tx.fail(err)
	}

	tx.addCompensationStep(func() error { return compensate(input, output) })
	return output, nil
}

// WithFallible starts fallible transaction execution.
// Inside f operations can be executed using ExecuteFallible.
// If any operation fails, all the already executed successful operation's compensation actions
// are executed in reverse order and the transaction returns with a failure.
func WithFallible[T any](f func(tx Fallible) (T, error)) (T, error) {
	tx := &fallible{}
	return f(tx)
}
