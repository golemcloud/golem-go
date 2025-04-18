package transaction

import (
	"fmt"

	"github.com/golemcloud/golem-go/binding/golem/api/host"
)

type InfallibleTx interface {
	addCompensationStep(compensationStep func() error)
	retry(err error)
	finish()
	ensureNoError()
}

type infallible struct {
	beginOpLogIndex   host.OplogIndex
	stepIndex         uint
	err               error
	compensationSteps []func() error
}

func newInfallible(beginOpLogIndex host.OplogIndex) *infallible {
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
	for i := stepsCount - 1; i >= 0; i-- {
		err := tx.compensationSteps[i]()
		if err != nil {
			err := &FailedAndRolledBackPartiallyError{
				ExecuteIndex:      tx.stepIndex,
				ExecuteError:      tx.err,
				CompensationIndex: uint(i),
				CompensationError: err,
			}
			panic(fmt.Sprintf("%s", err.Error()))
		}
	}
	host.SetOplogIndex(tx.beginOpLogIndex)
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

func ExecuteInfallible[I, O any](tx InfallibleTx, op Operation[I, O], input I) O {
	tx.ensureNoError()

	output, err := op.Execute(input)
	if err != nil {
		tx.retry(err)
		panic("unreachable after retry")
	}

	tx.addCompensationStep(func() error { return op.Compensate(input, output) })
	return output
}

// Infallible starts a transaction which retries in case of failure.
// Inside f operations can be executed using ExecuteInfallible.
// If any operation returns with a failure, all the already executed successful operation's
// compensation actions are executed in reverse order and the transaction gets retried,
// using Golem's active retry policy.
func Infallible[T any](f func(tx InfallibleTx) T) T {
	beginOpLogIndex := host.MarkBeginOperation()
	defer host.MarkEndOperation(beginOpLogIndex)
	tx := newInfallible(beginOpLogIndex)
	return f(tx)
}
