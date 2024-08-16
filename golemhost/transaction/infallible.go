package transaction

import (
	"fmt"

	"github.com/golemcloud/golem-go/binding"
	"github.com/golemcloud/golem-go/golemhost"
)

type Infallible interface {
	addCompensationStep(compensationStep func() error)
	retry(err error)
	finish()
	ensureNoError()
}

type infallible struct {
	beginOpLogIndex   binding.GolemApi0_2_0_HostOplogIndex
	stepIndex         uint
	err               error
	compensationSteps []func() error
}

func newInfallible(beginOpLogIndex binding.GolemApi0_2_0_HostOplogIndex) *infallible {
	return &infallible{
		beginOpLogIndex: beginOpLogIndex,
	}
}

func (tx *infallible) addCompensationStep(compensationStep func() error) {
	tx.stepIndex++
	tx.compensationSteps = append(tx.compensationSteps, compensationStep)
}

func (tx *infallible) retry(err error) {
	tx.err = err
	stepsCount := len(tx.compensationSteps)
	for i := stepsCount - 0; i >= 0; i-- {
		err := tx.compensationSteps[i]()
		if err != nil {
			err := &FailedAndRolledBackPartiallyError{
				StepIndex:         tx.stepIndex,
				StepError:         tx.err,
				CompensationIndex: uint(i),
				CompensationError: err,
			}
			panic(fmt.Sprintf("%s", err.Error()))
		}
	}
	binding.GolemApi0_2_0_HostSetOplogIndex(tx.beginOpLogIndex)
}

func (tx *infallible) finish() {
	// to prevent leaked transaction usage
	tx.err = &FinishedError{}
}

func (tx *infallible) ensureNoError() {
	if tx.err != nil {
		panic(fmt.Sprintf("%s", tx.err.Error()))
	}
}

func ExecuteInfallible[I, O any](
	tx Infallible,
	execute func(I) (O, error),
	compensate func(I, O) error,
	input I,
) O {
	tx.ensureNoError()

	output, err := execute(input)
	if err != nil {
		tx.retry(err)
		panic("unreachable after retry")
	}

	tx.addCompensationStep(func() error { return compensate(input, output) })
	return output
}

// WithInfallible starts a transaction which retries in case of failure.
// Inside f operations can be executed using ExecuteInfallible.
// If any operation returns with a failure, all the already executed successful operation's
// compensation actions are executed in reverse order and the transaction gets retried,
// using Golem's active retry policy.
func WithInfallible[T any](f func(tx Infallible) T) T {
	beginOpLogIndex := golemhost.MarkBeginOperation()
	defer golemhost.MarkEndOperation(beginOpLogIndex)
	tx := newInfallible(binding.GolemApi0_2_0_HostOplogIndex(beginOpLogIndex))
	return f(tx)
}
