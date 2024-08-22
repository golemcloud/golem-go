package golemhost

import (
	"fmt"
	"time"

	"github.com/golemcloud/golem-go/binding"
)

type FilterComparator int

const (
	FilterComparatorEqual FilterComparator = iota
	FilterComparatorNotEqual
	FilterComparatorGreaterEqual
	FilterComparatorGreater
	FilterComparatorLessEqual
	FilterComparatorLess
)

func (filterComparator FilterComparator) ToBinding() binding.GolemApi0_2_0_HostFilterComparator {
	switch filterComparator {
	case FilterComparatorEqual:
		return binding.GolemApi0_2_0_HostFilterComparatorEqual()
	case FilterComparatorNotEqual:
		return binding.GolemApi0_2_0_HostFilterComparatorNotEqual()
	case FilterComparatorGreaterEqual:
		return binding.GolemApi0_2_0_HostFilterComparatorGreaterEqual()
	case FilterComparatorGreater:
		return binding.GolemApi0_2_0_HostFilterComparatorGreater()
	case FilterComparatorLessEqual:
		return binding.GolemApi0_2_0_HostFilterComparatorLessEqual()
	case FilterComparatorLess:
		return binding.GolemApi0_2_0_HostFilterComparatorLess()
	default:
		panic(fmt.Sprintf("ToBinding: unhandled filterComparator: %d", filterComparator))
	}
}

type StringFilterComparator int

const (
	StringFilterComparatorEqual StringFilterComparator = iota
	StringFilterComparatorNotEqual
	StringFilterComparatorLike
	StringFilterComparatorNotLike
)

func (stringFilterComparator StringFilterComparator) ToBinding() binding.GolemApi0_2_0_HostStringFilterComparator {
	switch stringFilterComparator {
	case StringFilterComparatorEqual:
		return binding.GolemApi0_2_0_HostStringFilterComparatorEqual()
	case StringFilterComparatorNotEqual:
		return binding.GolemApi0_2_0_HostStringFilterComparatorNotEqual()
	case StringFilterComparatorLike:
		return binding.GolemApi0_2_0_HostStringFilterComparatorLike()
	case StringFilterComparatorNotLike:
		return binding.GolemApi0_2_0_HostStringFilterComparatorNotLike()
	default:
		panic(fmt.Sprintf("ToBinding: unhandled stringFilterComparator: %d", stringFilterComparator))
	}
}

type WorkerAnyFilter struct {
	Filters []WorkerAllFilter
}

func (f WorkerAnyFilter) ToBinding() binding.GolemApi0_2_0_HostWorkerAnyFilter {
	filter := binding.GolemApi0_2_0_HostWorkerAnyFilter{
		Filters: make([]binding.GolemApi0_2_0_HostWorkerAllFilter, len(f.Filters)),
	}
	for i := range f.Filters {
		filter.Filters[i] = f.Filters[i].ToBinding()
	}
	return filter
}

type WorkerAllFilter struct {
	Filters []WorkerFilter
}

func (f WorkerAllFilter) ToBinding() binding.GolemApi0_2_0_HostWorkerAllFilter {
	filter := binding.GolemApi0_2_0_HostWorkerAllFilter{}
	for i := range f.Filters {
		filter.Filters = append(filter.Filters, f.Filters[i].ToBinding()...)
	}
	return filter
}

type WorkerEnvFilter struct {
	Name  string
	Value string
}

type WorkerFilter struct {
	Name           *string
	NameComparator StringFilterComparator

	Status           *WorkerStatus
	StatusComparator FilterComparator

	Version           *uint64
	VersionComparator FilterComparator

	CreatedAt           *time.Time
	CreatedAtComparator FilterComparator

	Env           *WorkerEnvFilter
	EnvComparator StringFilterComparator
}

func (f WorkerFilter) ToBinding() []binding.GolemApi0_2_0_HostWorkerPropertyFilter {
	var filter []binding.GolemApi0_2_0_HostWorkerPropertyFilter

	if f.Name != nil {
		filter = append(
			filter,
			binding.GolemApi0_2_0_HostWorkerPropertyFilterName(
				binding.GolemApi0_2_0_HostWorkerNameFilter{
					Comparator: f.NameComparator.ToBinding(),
					Value:      *f.Name,
				},
			),
		)
	}

	if f.Status != nil {
		filter = append(
			filter,
			binding.GolemApi0_2_0_HostWorkerPropertyFilterStatus(
				binding.GolemApi0_2_0_HostWorkerStatusFilter{
					Comparator: f.StatusComparator.ToBinding(),
					Value:      f.Status.ToBinding(),
				},
			),
		)
	}

	if f.Version != nil {
		filter = append(
			filter,
			binding.GolemApi0_2_0_HostWorkerPropertyFilterVersion(
				binding.GolemApi0_2_0_HostWorkerVersionFilter{
					Comparator: f.VersionComparator.ToBinding(),
					Value:      *f.Version,
				},
			),
		)
	}

	if f.CreatedAt != nil {
		filter = append(
			filter,
			binding.GolemApi0_2_0_HostWorkerPropertyFilterCreatedAt(
				binding.GolemApi0_2_0_HostWorkerCreatedAtFilter{
					Comparator: f.CreatedAtComparator.ToBinding(),
					Value:      uint64(f.CreatedAt.UnixNano()),
				},
			),
		)
	}

	if f.Env != nil {
		filter = append(
			filter,
			binding.GolemApi0_2_0_HostWorkerPropertyFilterEnv(
				binding.GolemApi0_2_0_HostWorkerEnvFilter{
					Comparator: f.EnvComparator.ToBinding(),
					Name:       f.Env.Name,
					Value:      f.Env.Value,
				},
			),
		)
	}

	return filter
}
