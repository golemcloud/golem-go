package transaction

import (
	"fmt"

	"github.com/golemcloud/golem-go/binding"
	"github.com/golemcloud/golem-go/golemhost"
)

type Infallible interface {
	addCompensationStep(compensationStep func() error)
	retry(err error)
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

func ExecuteInfallibleStep[I, O any](
	tx Infallible,
	transactionStep func(I) (O, error),
	compensationStep func(O, I) error,
	input I,
) O {
	output, err := transactionStep(input)
	if err != nil {
		tx.retry(err)
		panic("unreachable after retry")
	}

	tx.addCompensationStep(func() error { return compensationStep(output, input) })
	return output
}

func WithInfallible[T any](f func(tx Infallible) (T, error)) (T, error) {
	return golemhost.Atomically(func() (T, error) {
		beginOpLogIndex := binding.GolemApi0_2_0_HostGetOplogIndex()
		tx := newInfallible(beginOpLogIndex)
		return f(tx)
	})
}
