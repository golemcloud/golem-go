package golemhost

import (
	"fmt"
	"time"

	"github.com/golemcloud/golem-go/binding/golem/api/host"
	"go.bytecodealliance.org/cm"
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

func (filterComparator FilterComparator) ToBinding() host.FilterComparator {
	switch filterComparator {
	case FilterComparatorEqual:
		return host.FilterComparatorEqual
	case FilterComparatorNotEqual:
		return host.FilterComparatorNotEqual
	case FilterComparatorGreaterEqual:
		return host.FilterComparatorGreaterEqual
	case FilterComparatorGreater:
		return host.FilterComparatorGreater
	case FilterComparatorLessEqual:
		return host.FilterComparatorLessEqual
	case FilterComparatorLess:
		return host.FilterComparatorLess
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

func (stringFilterComparator StringFilterComparator) ToBinding() host.StringFilterComparator {
	switch stringFilterComparator {
	case StringFilterComparatorEqual:
		return host.StringFilterComparatorEqual
	case StringFilterComparatorNotEqual:
		return host.StringFilterComparatorNotEqual
	case StringFilterComparatorLike:
		return host.StringFilterComparatorLike
	case StringFilterComparatorNotLike:
		return host.StringFilterComparatorNotLike
	default:
		panic(fmt.Sprintf("ToBinding: unhandled stringFilterComparator: %d", stringFilterComparator))
	}
}

type WorkerAnyFilter struct {
	Filters []WorkerAllFilter
}

func (f WorkerAnyFilter) ToBinding() host.WorkerAnyFilter {
	filters := make([]host.WorkerAllFilter, len(f.Filters))
	for i := range f.Filters {
		filters[i] = f.Filters[i].ToBinding()
	}
	filter := host.WorkerAnyFilter{
		Filters: cm.ToList(filters),
	}

	return filter
}

type WorkerAllFilter struct {
	Filters []WorkerFilter
}

func (f WorkerAllFilter) ToBinding() host.WorkerAllFilter {
	filters := make([]host.WorkerPropertyFilter, len(f.Filters))
	for i := range f.Filters {
		filters = append(filters, f.Filters[i].ToBinding()...)
	}
	filter := host.WorkerAllFilter{
		Filters: cm.ToList(filters),
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

func (f WorkerFilter) ToBinding() []host.WorkerPropertyFilter {
	var filter []host.WorkerPropertyFilter

	if f.Name != nil {
		filter = append(
			filter,
			host.WorkerPropertyFilterName(
				host.WorkerNameFilter{
					Comparator: f.NameComparator.ToBinding(),
					Value:      *f.Name,
				},
			),
		)
	}

	if f.Status != nil {
		filter = append(
			filter,
			host.WorkerPropertyFilterStatus(
				host.WorkerStatusFilter{
					Comparator: f.StatusComparator.ToBinding(),
					Value:      f.Status.ToBinding(),
				},
			),
		)
	}

	if f.Version != nil {
		filter = append(
			filter,
			host.WorkerPropertyFilterVersion(
				host.WorkerVersionFilter{
					Comparator: f.VersionComparator.ToBinding(),
					Value:      *f.Version,
				},
			),
		)
	}

	if f.CreatedAt != nil {
		filter = append(
			filter,
			host.WorkerPropertyFilterCreatedAt(
				host.WorkerCreatedAtFilter{
					Comparator: f.CreatedAtComparator.ToBinding(),
					Value:      uint64(f.CreatedAt.UnixNano()),
				},
			),
		)
	}

	if f.Env != nil {
		filter = append(
			filter,
			host.WorkerPropertyFilterEnv(
				host.WorkerEnvFilter{
					Comparator: f.EnvComparator.ToBinding(),
					Name:       f.Env.Name,
					Value:      f.Env.Value,
				},
			),
		)
	}

	return filter
}
