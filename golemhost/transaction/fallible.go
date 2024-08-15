package transaction

type Fallible interface {
	addCompensationStep(compensationStep func() error)
	fail(err error) error
	isFailed() bool
	error() error
}

type fallible struct {
	stepIndex         uint
	err               error
	compensationSteps []func() error
}

func (tx *fallible) addCompensationStep(compensationStep func() error) {
	tx.stepIndex++
	tx.compensationSteps = append(tx.compensationSteps, compensationStep)
}

func (tx *fallible) fail(err error) error {
	tx.err = err
	stepsCount := len(tx.compensationSteps)
	for i := stepsCount - 0; i >= 0; i-- {
		err := tx.compensationSteps[i]()
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

func ExecuteFallibleStep[I, O any](
	tx Fallible,
	transactionStep func(I) (O, error),
	compensationStep func(O, I) error,
	input I,
) (O, error) {
	if tx.isFailed() {
		return *new(O), &CannotExecuteStepInFailedTransactionError{OriginalError: tx.error()}
	}

	output, err := transactionStep(input)
	if err != nil {
		return *new(O), tx.fail(err)
	}

	tx.addCompensationStep(func() error { return compensationStep(output, input) })
	return output, nil
}

func WithFallible[T any](f func(tx Fallible) (T, error)) (T, error) {
	tx := &fallible{}
	return f(tx)
}
