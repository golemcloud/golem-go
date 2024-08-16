package transaction

type Operation[I any, O any] interface {
	Execute(I) (O, error)
	Compensate(I, O) error
}

type operation[I any, O any] struct {
	execute    func(I) (O, error)
	compensate func(I, O) error
}

func (o *operation[I, O]) Execute(input I) (O, error) {
	return o.execute(input)
}

func (o *operation[I, O]) Compensate(input I, output O) error {
	return o.compensate(input, output)
}

func NewOperation[I any, O any](execute func(I) (O, error), compensate func(I, O) error) Operation[I, O] {
	return &operation[I, O]{
		execute:    execute,
		compensate: compensate,
	}
}
