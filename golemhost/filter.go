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

func (filterComparator FilterComparator) ToBinding() binding.GolemApi1_1_6_HostFilterComparator {
	switch filterComparator {
	case FilterComparatorEqual:
		return binding.GolemApi1_1_6_HostFilterComparatorEqual()
	case FilterComparatorNotEqual:
		return binding.GolemApi1_1_6_HostFilterComparatorNotEqual()
	case FilterComparatorGreaterEqual:
		return binding.GolemApi1_1_6_HostFilterComparatorGreaterEqual()
	case FilterComparatorGreater:
		return binding.GolemApi1_1_6_HostFilterComparatorGreater()
	case FilterComparatorLessEqual:
		return binding.GolemApi1_1_6_HostFilterComparatorLessEqual()
	case FilterComparatorLess:
		return binding.GolemApi1_1_6_HostFilterComparatorLess()
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

func (stringFilterComparator StringFilterComparator) ToBinding() binding.GolemApi1_1_6_HostStringFilterComparator {
	switch stringFilterComparator {
	case StringFilterComparatorEqual:
		return binding.GolemApi1_1_6_HostStringFilterComparatorEqual()
	case StringFilterComparatorNotEqual:
		return binding.GolemApi1_1_6_HostStringFilterComparatorNotEqual()
	case StringFilterComparatorLike:
		return binding.GolemApi1_1_6_HostStringFilterComparatorLike()
	case StringFilterComparatorNotLike:
		return binding.GolemApi1_1_6_HostStringFilterComparatorNotLike()
	default:
		panic(fmt.Sprintf("ToBinding: unhandled stringFilterComparator: %d", stringFilterComparator))
	}
}

type WorkerAnyFilter struct {
	Filters []WorkerAllFilter
}

func (f WorkerAnyFilter) ToBinding() binding.GolemApi1_1_6_HostWorkerAnyFilter {
	filter := binding.GolemApi1_1_6_HostWorkerAnyFilter{
		Filters: make([]binding.GolemApi1_1_6_HostWorkerAllFilter, len(f.Filters)),
	}
	for i := range f.Filters {
		filter.Filters[i] = f.Filters[i].ToBinding()
	}
	return filter
}

type WorkerAllFilter struct {
	Filters []WorkerFilter
}

func (f WorkerAllFilter) ToBinding() binding.GolemApi1_1_6_HostWorkerAllFilter {
	filter := binding.GolemApi1_1_6_HostWorkerAllFilter{}
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

func (f WorkerFilter) ToBinding() []binding.GolemApi1_1_6_HostWorkerPropertyFilter {
	var filter []binding.GolemApi1_1_6_HostWorkerPropertyFilter

	if f.Name != nil {
		filter = append(
			filter,
			binding.GolemApi1_1_6_HostWorkerPropertyFilterName(
				binding.GolemApi1_1_6_HostWorkerNameFilter{
					Comparator: f.NameComparator.ToBinding(),
					Value:      *f.Name,
				},
			),
		)
	}

	if f.Status != nil {
		filter = append(
			filter,
			binding.GolemApi1_1_6_HostWorkerPropertyFilterStatus(
				binding.GolemApi1_1_6_HostWorkerStatusFilter{
					Comparator: f.StatusComparator.ToBinding(),
					Value:      f.Status.ToBinding(),
				},
			),
		)
	}

	if f.Version != nil {
		filter = append(
			filter,
			binding.GolemApi1_1_6_HostWorkerPropertyFilterVersion(
				binding.GolemApi1_1_6_HostWorkerVersionFilter{
					Comparator: f.VersionComparator.ToBinding(),
					Value:      *f.Version,
				},
			),
		)
	}

	if f.CreatedAt != nil {
		filter = append(
			filter,
			binding.GolemApi1_1_6_HostWorkerPropertyFilterCreatedAt(
				binding.GolemApi1_1_6_HostWorkerCreatedAtFilter{
					Comparator: f.CreatedAtComparator.ToBinding(),
					Value:      uint64(f.CreatedAt.UnixNano()),
				},
			),
		)
	}

	if f.Env != nil {
		filter = append(
			filter,
			binding.GolemApi1_1_6_HostWorkerPropertyFilterEnv(
				binding.GolemApi1_1_6_HostWorkerEnvFilter{
					Comparator: f.EnvComparator.ToBinding(),
					Name:       f.Env.Name,
					Value:      f.Env.Value,
				},
			),
		)
	}

	return filter
}
