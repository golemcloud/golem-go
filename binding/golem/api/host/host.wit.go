// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package host represents the imported interface "golem:api/host@1.1.7".
//
// The Golem host API provides low level access to Golem specific features such as
// promises and control over
// the durability and transactional guarantees the executor provides.
package host

import (
	"github.com/golemcloud/golem-go/binding/golem/rpc/types"
	monotonicclock "github.com/golemcloud/golem-go/binding/wasi/clocks/monotonic-clock"
	"go.bytecodealliance.org/cm"
)

// Duration represents the type alias "golem:api/host@1.1.7#duration".
//
// See [monotonicclock.Duration] for more information.
type Duration = monotonicclock.Duration

// ComponentID represents the type alias "golem:api/host@1.1.7#component-id".
//
// See [types.ComponentID] for more information.
type ComponentID = types.ComponentID

// UUID represents the type alias "golem:api/host@1.1.7#uuid".
//
// See [types.UUID] for more information.
type UUID = types.UUID

// WorkerID represents the type alias "golem:api/host@1.1.7#worker-id".
//
// See [types.WorkerID] for more information.
type WorkerID = types.WorkerID

// OplogIndex represents the u64 "golem:api/host@1.1.7#oplog-index".
//
// An index into the persistent log storing all performed operations of a worker
//
//	type oplog-index = u64
type OplogIndex uint64

// PromiseID represents the record "golem:api/host@1.1.7#promise-id".
//
// A promise ID is a value that can be passed to an external Golem API to complete
// that promise
// from an arbitrary external source, while Golem workers can await for this completion.
//
//	record promise-id {
//		worker-id: worker-id,
//		oplog-idx: oplog-index,
//	}
type PromiseID struct {
	_        cm.HostLayout `json:"-"`
	WorkerID WorkerID      `json:"worker-id"`
	OplogIdx OplogIndex    `json:"oplog-idx"`
}

// ComponentVersion represents the u64 "golem:api/host@1.1.7#component-version".
//
// Represents a Golem component's version
//
//	type component-version = u64
type ComponentVersion uint64

// AccountID represents the record "golem:api/host@1.1.7#account-id".
//
// Represents a Golem Cloud account
//
//	record account-id {
//		value: string,
//	}
type AccountID struct {
	_     cm.HostLayout `json:"-"`
	Value string        `json:"value"`
}

// RetryPolicy represents the record "golem:api/host@1.1.7#retry-policy".
//
// Configures how the executor retries failures
//
//	record retry-policy {
//		max-attempts: u32,
//		min-delay: duration,
//		max-delay: duration,
//		multiplier: f64,
//		max-jitter-factor: option<f64>,
//	}
type RetryPolicy struct {
	_ cm.HostLayout `json:"-"`
	// The maximum number of retries before the worker becomes permanently failed
	MaxAttempts uint32 `json:"max-attempts"`

	// The minimum delay between retries (applied to the first retry)
	MinDelay Duration `json:"min-delay"`

	// The maximum delay between retries
	MaxDelay Duration `json:"max-delay"`

	// Multiplier applied to the delay on each retry to implement exponential backoff
	Multiplier float64 `json:"multiplier"`

	// The maximum amount of jitter to add to the delay
	MaxJitterFactor cm.Option[float64] `json:"max-jitter-factor"`
}

// PersistenceLevel represents the variant "golem:api/host@1.1.7#persistence-level".
//
// Configurable persistence level for workers
//
//	variant persistence-level {
//		persist-nothing,
//		persist-remote-side-effects,
//		smart,
//	}
type PersistenceLevel uint8

const (
	PersistenceLevelPersistNothing PersistenceLevel = iota
	PersistenceLevelPersistRemoteSideEffects
	PersistenceLevelSmart
)

var _PersistenceLevelStrings = [3]string{
	"persist-nothing",
	"persist-remote-side-effects",
	"smart",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e PersistenceLevel) String() string {
	return _PersistenceLevelStrings[e]
}

// MarshalText implements [encoding.TextMarshaler].
func (e PersistenceLevel) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

// UnmarshalText implements [encoding.TextUnmarshaler], unmarshaling into an enum
// case. Returns an error if the supplied text is not one of the enum cases.
func (e *PersistenceLevel) UnmarshalText(text []byte) error {
	return _PersistenceLevelUnmarshalCase(e, text)
}

var _PersistenceLevelUnmarshalCase = cm.CaseUnmarshaler[PersistenceLevel](_PersistenceLevelStrings[:])

// UpdateMode represents the enum "golem:api/host@1.1.7#update-mode".
//
// Describes how to update a worker to a different component version
//
//	enum update-mode {
//		automatic,
//		snapshot-based
//	}
type UpdateMode uint8

const (
	// Automatic update tries to recover the worker using the new component version
	// and may fail if there is a divergence.
	UpdateModeAutomatic UpdateMode = iota

	// Manual, snapshot-based update uses a user-defined implementation of the `save-snapshot`
	// interface
	// to store the worker's state, and a user-defined implementation of the `load-snapshot`
	// interface to
	// load it into the new version.
	UpdateModeSnapshotBased
)

var _UpdateModeStrings = [2]string{
	"automatic",
	"snapshot-based",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e UpdateMode) String() string {
	return _UpdateModeStrings[e]
}

// MarshalText implements [encoding.TextMarshaler].
func (e UpdateMode) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

// UnmarshalText implements [encoding.TextUnmarshaler], unmarshaling into an enum
// case. Returns an error if the supplied text is not one of the enum cases.
func (e *UpdateMode) UnmarshalText(text []byte) error {
	return _UpdateModeUnmarshalCase(e, text)
}

var _UpdateModeUnmarshalCase = cm.CaseUnmarshaler[UpdateMode](_UpdateModeStrings[:])

// FilterComparator represents the enum "golem:api/host@1.1.7#filter-comparator".
//
//	enum filter-comparator {
//		equal,
//		not-equal,
//		greater-equal,
//		greater,
//		less-equal,
//		less
//	}
type FilterComparator uint8

const (
	FilterComparatorEqual FilterComparator = iota
	FilterComparatorNotEqual
	FilterComparatorGreaterEqual
	FilterComparatorGreater
	FilterComparatorLessEqual
	FilterComparatorLess
)

var _FilterComparatorStrings = [6]string{
	"equal",
	"not-equal",
	"greater-equal",
	"greater",
	"less-equal",
	"less",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e FilterComparator) String() string {
	return _FilterComparatorStrings[e]
}

// MarshalText implements [encoding.TextMarshaler].
func (e FilterComparator) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

// UnmarshalText implements [encoding.TextUnmarshaler], unmarshaling into an enum
// case. Returns an error if the supplied text is not one of the enum cases.
func (e *FilterComparator) UnmarshalText(text []byte) error {
	return _FilterComparatorUnmarshalCase(e, text)
}

var _FilterComparatorUnmarshalCase = cm.CaseUnmarshaler[FilterComparator](_FilterComparatorStrings[:])

// StringFilterComparator represents the enum "golem:api/host@1.1.7#string-filter-comparator".
//
//	enum string-filter-comparator {
//		equal,
//		not-equal,
//		like,
//		not-like
//	}
type StringFilterComparator uint8

const (
	StringFilterComparatorEqual StringFilterComparator = iota
	StringFilterComparatorNotEqual
	StringFilterComparatorLike
	StringFilterComparatorNotLike
)

var _StringFilterComparatorStrings = [4]string{
	"equal",
	"not-equal",
	"like",
	"not-like",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e StringFilterComparator) String() string {
	return _StringFilterComparatorStrings[e]
}

// MarshalText implements [encoding.TextMarshaler].
func (e StringFilterComparator) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

// UnmarshalText implements [encoding.TextUnmarshaler], unmarshaling into an enum
// case. Returns an error if the supplied text is not one of the enum cases.
func (e *StringFilterComparator) UnmarshalText(text []byte) error {
	return _StringFilterComparatorUnmarshalCase(e, text)
}

var _StringFilterComparatorUnmarshalCase = cm.CaseUnmarshaler[StringFilterComparator](_StringFilterComparatorStrings[:])

// WorkerStatus represents the enum "golem:api/host@1.1.7#worker-status".
//
//	enum worker-status {
//		running,
//		idle,
//		suspended,
//		interrupted,
//		retrying,
//		failed,
//		exited
//	}
type WorkerStatus uint8

const (
	// The worker is running an invoked function
	WorkerStatusRunning WorkerStatus = iota

	// The worker is ready to run an invoked function
	WorkerStatusIdle

	// An invocation is active but waiting for something (sleeping, waiting for a promise)
	WorkerStatusSuspended

	// The last invocation was interrupted but will be resumed
	WorkerStatusInterrupted

	// The last invocation failed and a retry was scheduled
	WorkerStatusRetrying

	// The last invocation failed and the worker can no longer be used
	WorkerStatusFailed

	// The worker exited after a successful invocation and can no longer be invoked
	WorkerStatusExited
)

var _WorkerStatusStrings = [7]string{
	"running",
	"idle",
	"suspended",
	"interrupted",
	"retrying",
	"failed",
	"exited",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e WorkerStatus) String() string {
	return _WorkerStatusStrings[e]
}

// MarshalText implements [encoding.TextMarshaler].
func (e WorkerStatus) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

// UnmarshalText implements [encoding.TextUnmarshaler], unmarshaling into an enum
// case. Returns an error if the supplied text is not one of the enum cases.
func (e *WorkerStatus) UnmarshalText(text []byte) error {
	return _WorkerStatusUnmarshalCase(e, text)
}

var _WorkerStatusUnmarshalCase = cm.CaseUnmarshaler[WorkerStatus](_WorkerStatusStrings[:])

// WorkerNameFilter represents the record "golem:api/host@1.1.7#worker-name-filter".
//
//	record worker-name-filter {
//		comparator: string-filter-comparator,
//		value: string,
//	}
type WorkerNameFilter struct {
	_          cm.HostLayout          `json:"-"`
	Comparator StringFilterComparator `json:"comparator"`
	Value      string                 `json:"value"`
}

// WorkerStatusFilter represents the record "golem:api/host@1.1.7#worker-status-filter".
//
//	record worker-status-filter {
//		comparator: filter-comparator,
//		value: worker-status,
//	}
type WorkerStatusFilter struct {
	_          cm.HostLayout    `json:"-"`
	Comparator FilterComparator `json:"comparator"`
	Value      WorkerStatus     `json:"value"`
}

// WorkerVersionFilter represents the record "golem:api/host@1.1.7#worker-version-filter".
//
//	record worker-version-filter {
//		comparator: filter-comparator,
//		value: u64,
//	}
type WorkerVersionFilter struct {
	_          cm.HostLayout    `json:"-"`
	Comparator FilterComparator `json:"comparator"`
	Value      uint64           `json:"value"`
}

// WorkerCreatedAtFilter represents the record "golem:api/host@1.1.7#worker-created-at-filter".
//
//	record worker-created-at-filter {
//		comparator: filter-comparator,
//		value: u64,
//	}
type WorkerCreatedAtFilter struct {
	_          cm.HostLayout    `json:"-"`
	Comparator FilterComparator `json:"comparator"`
	Value      uint64           `json:"value"`
}

// WorkerEnvFilter represents the record "golem:api/host@1.1.7#worker-env-filter".
//
//	record worker-env-filter {
//		name: string,
//		comparator: string-filter-comparator,
//		value: string,
//	}
type WorkerEnvFilter struct {
	_          cm.HostLayout          `json:"-"`
	Name       string                 `json:"name"`
	Comparator StringFilterComparator `json:"comparator"`
	Value      string                 `json:"value"`
}

// WorkerPropertyFilter represents the variant "golem:api/host@1.1.7#worker-property-filter".
//
//	variant worker-property-filter {
//		name(worker-name-filter),
//		status(worker-status-filter),
//		version(worker-version-filter),
//		created-at(worker-created-at-filter),
//		env(worker-env-filter),
//	}
type WorkerPropertyFilter cm.Variant[uint8, WorkerEnvFilterShape, WorkerVersionFilter]

// WorkerPropertyFilterName returns a [WorkerPropertyFilter] of case "name".
func WorkerPropertyFilterName(data WorkerNameFilter) WorkerPropertyFilter {
	return cm.New[WorkerPropertyFilter](0, data)
}

// Name returns a non-nil *[WorkerNameFilter] if [WorkerPropertyFilter] represents the variant case "name".
func (self *WorkerPropertyFilter) Name() *WorkerNameFilter {
	return cm.Case[WorkerNameFilter](self, 0)
}

// WorkerPropertyFilterStatus returns a [WorkerPropertyFilter] of case "status".
func WorkerPropertyFilterStatus(data WorkerStatusFilter) WorkerPropertyFilter {
	return cm.New[WorkerPropertyFilter](1, data)
}

// Status returns a non-nil *[WorkerStatusFilter] if [WorkerPropertyFilter] represents the variant case "status".
func (self *WorkerPropertyFilter) Status() *WorkerStatusFilter {
	return cm.Case[WorkerStatusFilter](self, 1)
}

// WorkerPropertyFilterVersion returns a [WorkerPropertyFilter] of case "version".
func WorkerPropertyFilterVersion(data WorkerVersionFilter) WorkerPropertyFilter {
	return cm.New[WorkerPropertyFilter](2, data)
}

// Version returns a non-nil *[WorkerVersionFilter] if [WorkerPropertyFilter] represents the variant case "version".
func (self *WorkerPropertyFilter) Version() *WorkerVersionFilter {
	return cm.Case[WorkerVersionFilter](self, 2)
}

// WorkerPropertyFilterCreatedAt returns a [WorkerPropertyFilter] of case "created-at".
func WorkerPropertyFilterCreatedAt(data WorkerCreatedAtFilter) WorkerPropertyFilter {
	return cm.New[WorkerPropertyFilter](3, data)
}

// CreatedAt returns a non-nil *[WorkerCreatedAtFilter] if [WorkerPropertyFilter] represents the variant case "created-at".
func (self *WorkerPropertyFilter) CreatedAt() *WorkerCreatedAtFilter {
	return cm.Case[WorkerCreatedAtFilter](self, 3)
}

// WorkerPropertyFilterEnv returns a [WorkerPropertyFilter] of case "env".
func WorkerPropertyFilterEnv(data WorkerEnvFilter) WorkerPropertyFilter {
	return cm.New[WorkerPropertyFilter](4, data)
}

// Env returns a non-nil *[WorkerEnvFilter] if [WorkerPropertyFilter] represents the variant case "env".
func (self *WorkerPropertyFilter) Env() *WorkerEnvFilter {
	return cm.Case[WorkerEnvFilter](self, 4)
}

var _WorkerPropertyFilterStrings = [5]string{
	"name",
	"status",
	"version",
	"created-at",
	"env",
}

// String implements [fmt.Stringer], returning the variant case name of v.
func (v WorkerPropertyFilter) String() string {
	return _WorkerPropertyFilterStrings[v.Tag()]
}

// WorkerAllFilter represents the record "golem:api/host@1.1.7#worker-all-filter".
//
//	record worker-all-filter {
//		filters: list<worker-property-filter>,
//	}
type WorkerAllFilter struct {
	_       cm.HostLayout                 `json:"-"`
	Filters cm.List[WorkerPropertyFilter] `json:"filters"`
}

// WorkerAnyFilter represents the record "golem:api/host@1.1.7#worker-any-filter".
//
//	record worker-any-filter {
//		filters: list<worker-all-filter>,
//	}
type WorkerAnyFilter struct {
	_       cm.HostLayout            `json:"-"`
	Filters cm.List[WorkerAllFilter] `json:"filters"`
}

// WorkerMetadata represents the record "golem:api/host@1.1.7#worker-metadata".
//
//	record worker-metadata {
//		worker-id: worker-id,
//		args: list<string>,
//		env: list<tuple<string, string>>,
//		status: worker-status,
//		component-version: u64,
//		retry-count: u64,
//	}
type WorkerMetadata struct {
	_                cm.HostLayout      `json:"-"`
	WorkerID         WorkerID           `json:"worker-id"`
	Args             cm.List[string]    `json:"args"`
	Env              cm.List[[2]string] `json:"env"`
	Status           WorkerStatus       `json:"status"`
	ComponentVersion uint64             `json:"component-version"`
	RetryCount       uint64             `json:"retry-count"`
}

// GetWorkers represents the imported resource "golem:api/host@1.1.7#get-workers".
//
//	resource get-workers
type GetWorkers cm.Resource

// ResourceDrop represents the imported resource-drop for resource "get-workers".
//
// Drops a resource handle.
//
//go:nosplit
func (self GetWorkers) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_GetWorkersResourceDrop((uint32)(self0))
	return
}

// NewGetWorkers represents the imported constructor for resource "get-workers".
//
//	constructor(component-id: component-id, filter: option<worker-any-filter>, precise:
//	bool)
//
//go:nosplit
func NewGetWorkers(componentID ComponentID, filter cm.Option[WorkerAnyFilter], precise bool) (result GetWorkers) {
	componentId0, componentId1 := lower_ComponentID(componentID)
	filter0, filter1, filter2 := lower_OptionWorkerAnyFilter(filter)
	precise0 := (uint32)(cm.BoolToU32(precise))
	result0 := wasmimport_NewGetWorkers((uint64)(componentId0), (uint64)(componentId1), (uint32)(filter0), (*WorkerAllFilter)(filter1), (uint32)(filter2), (uint32)(precise0))
	result = cm.Reinterpret[GetWorkers]((uint32)(result0))
	return
}

// GetNext represents the imported method "get-next".
//
//	get-next: func() -> option<list<worker-metadata>>
//
//go:nosplit
func (self GetWorkers) GetNext() (result cm.Option[cm.List[WorkerMetadata]]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_GetWorkersGetNext((uint32)(self0), &result)
	return
}

// RevertWorkerTarget represents the variant "golem:api/host@1.1.7#revert-worker-target".
//
// Target parameter for the `revert-worker` operation
//
//	variant revert-worker-target {
//		revert-to-oplog-index(oplog-index),
//		revert-last-invocations(u64),
//	}
type RevertWorkerTarget cm.Variant[uint8, uint64, OplogIndex]

// RevertWorkerTargetRevertToOplogIndex returns a [RevertWorkerTarget] of case "revert-to-oplog-index".
//
// Revert to a specific oplog index. The given index will be the last one to be kept.
func RevertWorkerTargetRevertToOplogIndex(data OplogIndex) RevertWorkerTarget {
	return cm.New[RevertWorkerTarget](0, data)
}

// RevertToOplogIndex returns a non-nil *[OplogIndex] if [RevertWorkerTarget] represents the variant case "revert-to-oplog-index".
func (self *RevertWorkerTarget) RevertToOplogIndex() *OplogIndex {
	return cm.Case[OplogIndex](self, 0)
}

// RevertWorkerTargetRevertLastInvocations returns a [RevertWorkerTarget] of case "revert-last-invocations".
//
// Revert the last N invocations.
func RevertWorkerTargetRevertLastInvocations(data uint64) RevertWorkerTarget {
	return cm.New[RevertWorkerTarget](1, data)
}

// RevertLastInvocations returns a non-nil *[uint64] if [RevertWorkerTarget] represents the variant case "revert-last-invocations".
func (self *RevertWorkerTarget) RevertLastInvocations() *uint64 {
	return cm.Case[uint64](self, 1)
}

var _RevertWorkerTargetStrings = [2]string{
	"revert-to-oplog-index",
	"revert-last-invocations",
}

// String implements [fmt.Stringer], returning the variant case name of v.
func (v RevertWorkerTarget) String() string {
	return _RevertWorkerTargetStrings[v.Tag()]
}

// ForkResult represents the enum "golem:api/host@1.1.7#fork-result".
//
// Indicates which worker the code is running on after `fork`
//
//	enum fork-result {
//		original,
//		forked
//	}
type ForkResult uint8

const (
	// The original worker that called `fork`
	ForkResultOriginal ForkResult = iota

	// The new worker
	ForkResultForked
)

var _ForkResultStrings = [2]string{
	"original",
	"forked",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e ForkResult) String() string {
	return _ForkResultStrings[e]
}

// MarshalText implements [encoding.TextMarshaler].
func (e ForkResult) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

// UnmarshalText implements [encoding.TextUnmarshaler], unmarshaling into an enum
// case. Returns an error if the supplied text is not one of the enum cases.
func (e *ForkResult) UnmarshalText(text []byte) error {
	return _ForkResultUnmarshalCase(e, text)
}

var _ForkResultUnmarshalCase = cm.CaseUnmarshaler[ForkResult](_ForkResultStrings[:])

// CreatePromise represents the imported function "create-promise".
//
// Create a new promise
//
//	create-promise: func() -> promise-id
//
//go:nosplit
func CreatePromise() (result PromiseID) {
	wasmimport_CreatePromise(&result)
	return
}

// AwaitPromise represents the imported function "await-promise".
//
// Suspends execution until the given promise gets completed, and returns the payload
// passed to
// the promise completion.
//
//	await-promise: func(promise-id: promise-id) -> list<u8>
//
//go:nosplit
func AwaitPromise(promiseID PromiseID) (result cm.List[uint8]) {
	promiseId0, promiseId1, promiseId2, promiseId3, promiseId4 := lower_PromiseID(promiseID)
	wasmimport_AwaitPromise((uint64)(promiseId0), (uint64)(promiseId1), (*uint8)(promiseId2), (uint32)(promiseId3), (uint64)(promiseId4), &result)
	return
}

// PollPromise represents the imported function "poll-promise".
//
// Checks whether the given promise is completed. If not, it returns None. If the
// promise is completed,
// it returns the payload passed to the promise completion.
//
//	poll-promise: func(promise-id: promise-id) -> option<list<u8>>
//
//go:nosplit
func PollPromise(promiseID PromiseID) (result cm.Option[cm.List[uint8]]) {
	promiseId0, promiseId1, promiseId2, promiseId3, promiseId4 := lower_PromiseID(promiseID)
	wasmimport_PollPromise((uint64)(promiseId0), (uint64)(promiseId1), (*uint8)(promiseId2), (uint32)(promiseId3), (uint64)(promiseId4), &result)
	return
}

// CompletePromise represents the imported function "complete-promise".
//
// Completes the given promise with the given payload. Returns true if the promise
// was completed, false
// if the promise was already completed. The payload is passed to the worker that
// is awaiting the promise.
//
//	complete-promise: func(promise-id: promise-id, data: list<u8>) -> bool
//
//go:nosplit
func CompletePromise(promiseID PromiseID, data cm.List[uint8]) (result bool) {
	promiseId0, promiseId1, promiseId2, promiseId3, promiseId4 := lower_PromiseID(promiseID)
	data0, data1 := cm.LowerList(data)
	result0 := wasmimport_CompletePromise((uint64)(promiseId0), (uint64)(promiseId1), (*uint8)(promiseId2), (uint32)(promiseId3), (uint64)(promiseId4), (*uint8)(data0), (uint32)(data1))
	result = (bool)(cm.U32ToBool((uint32)(result0)))
	return
}

// DeletePromise represents the imported function "delete-promise".
//
// Deletes the given promise
//
//	delete-promise: func(promise-id: promise-id)
//
//go:nosplit
func DeletePromise(promiseID PromiseID) {
	promiseId0, promiseId1, promiseId2, promiseId3, promiseId4 := lower_PromiseID(promiseID)
	wasmimport_DeletePromise((uint64)(promiseId0), (uint64)(promiseId1), (*uint8)(promiseId2), (uint32)(promiseId3), (uint64)(promiseId4))
	return
}

// GetOplogIndex represents the imported function "get-oplog-index".
//
// Returns the current position in the persistent op log
//
//	get-oplog-index: func() -> oplog-index
//
//go:nosplit
func GetOplogIndex() (result OplogIndex) {
	result0 := wasmimport_GetOplogIndex()
	result = (OplogIndex)((uint64)(result0))
	return
}

// SetOplogIndex represents the imported function "set-oplog-index".
//
// Makes the current worker travel back in time and continue execution from the given
// position in the persistent
// op log.
//
//	set-oplog-index: func(oplog-idx: oplog-index)
//
//go:nosplit
func SetOplogIndex(oplogIdx OplogIndex) {
	oplogIdx0 := (uint64)(oplogIdx)
	wasmimport_SetOplogIndex((uint64)(oplogIdx0))
	return
}

// OplogCommit represents the imported function "oplog-commit".
//
// Blocks the execution until the oplog has been written to at least the specified
// number of replicas,
// or the maximum number of replicas if the requested number is higher.
//
//	oplog-commit: func(replicas: u8)
//
//go:nosplit
func OplogCommit(replicas uint8) {
	replicas0 := (uint32)(replicas)
	wasmimport_OplogCommit((uint32)(replicas0))
	return
}

// MarkBeginOperation represents the imported function "mark-begin-operation".
//
// Marks the beginning of an atomic operation.
// In case of a failure within the region selected by `mark-begin-operation` and `mark-end-operation`
// the whole region will be reexecuted on retry.
// The end of the region is when `mark-end-operation` is called with the returned
// oplog-index.
//
//	mark-begin-operation: func() -> oplog-index
//
//go:nosplit
func MarkBeginOperation() (result OplogIndex) {
	result0 := wasmimport_MarkBeginOperation()
	result = (OplogIndex)((uint64)(result0))
	return
}

// MarkEndOperation represents the imported function "mark-end-operation".
//
// Commits this atomic operation. After `mark-end-operation` is called for a given
// index, further calls
// with the same parameter will do nothing.
//
//	mark-end-operation: func(begin: oplog-index)
//
//go:nosplit
func MarkEndOperation(begin OplogIndex) {
	begin0 := (uint64)(begin)
	wasmimport_MarkEndOperation((uint64)(begin0))
	return
}

// GetRetryPolicy represents the imported function "get-retry-policy".
//
// Gets the current retry policy associated with the worker
//
//	get-retry-policy: func() -> retry-policy
//
//go:nosplit
func GetRetryPolicy() (result RetryPolicy) {
	wasmimport_GetRetryPolicy(&result)
	return
}

// SetRetryPolicy represents the imported function "set-retry-policy".
//
// Overrides the current retry policy associated with the worker. Following this call,
// `get-retry-policy` will return the
// new retry policy.
//
//	set-retry-policy: func(new-retry-policy: retry-policy)
//
//go:nosplit
func SetRetryPolicy(newRetryPolicy RetryPolicy) {
	newRetryPolicy0, newRetryPolicy1, newRetryPolicy2, newRetryPolicy3, newRetryPolicy4, newRetryPolicy5 := lower_RetryPolicy(newRetryPolicy)
	wasmimport_SetRetryPolicy((uint32)(newRetryPolicy0), (uint64)(newRetryPolicy1), (uint64)(newRetryPolicy2), (float64)(newRetryPolicy3), (uint32)(newRetryPolicy4), (float64)(newRetryPolicy5))
	return
}

// GetOplogPersistenceLevel represents the imported function "get-oplog-persistence-level".
//
// Gets the worker's current persistence level.
//
//	get-oplog-persistence-level: func() -> persistence-level
//
//go:nosplit
func GetOplogPersistenceLevel() (result PersistenceLevel) {
	result0 := wasmimport_GetOplogPersistenceLevel()
	result = (PersistenceLevel)((uint32)(result0))
	return
}

// SetOplogPersistenceLevel represents the imported function "set-oplog-persistence-level".
//
// Sets the worker's current persistence level. This can increase the performance
// of execution in cases where durable
// execution is not required.
//
//	set-oplog-persistence-level: func(new-persistence-level: persistence-level)
//
//go:nosplit
func SetOplogPersistenceLevel(newPersistenceLevel PersistenceLevel) {
	newPersistenceLevel0 := (uint32)(newPersistenceLevel)
	wasmimport_SetOplogPersistenceLevel((uint32)(newPersistenceLevel0))
	return
}

// GetIdempotenceMode represents the imported function "get-idempotence-mode".
//
// Gets the current idempotence mode. See `set-idempotence-mode` for details.
//
//	get-idempotence-mode: func() -> bool
//
//go:nosplit
func GetIdempotenceMode() (result bool) {
	result0 := wasmimport_GetIdempotenceMode()
	result = (bool)(cm.U32ToBool((uint32)(result0)))
	return
}

// SetIdempotenceMode represents the imported function "set-idempotence-mode".
//
// Sets the current idempotence mode. The default is true.
// True means side-effects are treated idempotent and Golem guarantees at-least-once
// semantics.
// In case of false the executor provides at-most-once semantics, failing the worker
// in case it is
// not known if the side effect was already executed.
//
//	set-idempotence-mode: func(idempotent: bool)
//
//go:nosplit
func SetIdempotenceMode(idempotent bool) {
	idempotent0 := (uint32)(cm.BoolToU32(idempotent))
	wasmimport_SetIdempotenceMode((uint32)(idempotent0))
	return
}

// GenerateIdempotencyKey represents the imported function "generate-idempotency-key".
//
// Generates an idempotency key. This operation will never be replayed —
// i.e. not only is this key generated, but it is persisted and committed, such that
// the key can be used in third-party systems (e.g. payment processing)
// to introduce idempotence.
//
//	generate-idempotency-key: func() -> uuid
//
//go:nosplit
func GenerateIdempotencyKey() (result UUID) {
	wasmimport_GenerateIdempotencyKey(&result)
	return
}

// UpdateWorker represents the imported function "update-worker".
//
// Initiates an update attempt for the given worker. The function returns immediately
// once the request has been processed,
// not waiting for the worker to get updated.
//
//	update-worker: func(worker-id: worker-id, target-version: component-version, mode:
//	update-mode)
//
//go:nosplit
func UpdateWorker(workerID WorkerID, targetVersion ComponentVersion, mode UpdateMode) {
	workerId0, workerId1, workerId2, workerId3 := lower_WorkerID(workerID)
	targetVersion0 := (uint64)(targetVersion)
	mode0 := (uint32)(mode)
	wasmimport_UpdateWorker((uint64)(workerId0), (uint64)(workerId1), (*uint8)(workerId2), (uint32)(workerId3), (uint64)(targetVersion0), (uint32)(mode0))
	return
}

// GetSelfMetadata represents the imported function "get-self-metadata".
//
// Get current worker metadata
//
//	get-self-metadata: func() -> worker-metadata
//
//go:nosplit
func GetSelfMetadata() (result WorkerMetadata) {
	wasmimport_GetSelfMetadata(&result)
	return
}

// GetWorkerMetadata represents the imported function "get-worker-metadata".
//
// Get worker metadata
//
//	get-worker-metadata: func(worker-id: worker-id) -> option<worker-metadata>
//
//go:nosplit
func GetWorkerMetadata(workerID WorkerID) (result cm.Option[WorkerMetadata]) {
	workerId0, workerId1, workerId2, workerId3 := lower_WorkerID(workerID)
	wasmimport_GetWorkerMetadata((uint64)(workerId0), (uint64)(workerId1), (*uint8)(workerId2), (uint32)(workerId3), &result)
	return
}

// ForkWorker represents the imported function "fork-worker".
//
// Fork a worker to another worker at a given oplog index
//
//	fork-worker: func(source-worker-id: worker-id, target-worker-id: worker-id, oplog-idx-cut-off:
//	oplog-index)
//
//go:nosplit
func ForkWorker(sourceWorkerID WorkerID, targetWorkerID WorkerID, oplogIdxCutOff OplogIndex) {
	sourceWorkerId0, sourceWorkerId1, sourceWorkerId2, sourceWorkerId3 := lower_WorkerID(sourceWorkerID)
	targetWorkerId0, targetWorkerId1, targetWorkerId2, targetWorkerId3 := lower_WorkerID(targetWorkerID)
	oplogIdxCutOff0 := (uint64)(oplogIdxCutOff)
	wasmimport_ForkWorker((uint64)(sourceWorkerId0), (uint64)(sourceWorkerId1), (*uint8)(sourceWorkerId2), (uint32)(sourceWorkerId3), (uint64)(targetWorkerId0), (uint64)(targetWorkerId1), (*uint8)(targetWorkerId2), (uint32)(targetWorkerId3), (uint64)(oplogIdxCutOff0))
	return
}

// RevertWorker represents the imported function "revert-worker".
//
// Revert a worker to a previous state
//
//	revert-worker: func(worker-id: worker-id, revert-target: revert-worker-target)
//
//go:nosplit
func RevertWorker(workerID WorkerID, revertTarget RevertWorkerTarget) {
	workerId0, workerId1, workerId2, workerId3 := lower_WorkerID(workerID)
	revertTarget0, revertTarget1 := lower_RevertWorkerTarget(revertTarget)
	wasmimport_RevertWorker((uint64)(workerId0), (uint64)(workerId1), (*uint8)(workerId2), (uint32)(workerId3), (uint32)(revertTarget0), (uint64)(revertTarget1))
	return
}

// ResolveComponentID represents the imported function "resolve-component-id".
//
// Get the component-id for a given component reference.
// Returns none when no component with the specified reference exists.
// The syntax of the component reference is implementation dependent.
//
// Golem OSS: "{component_name}"
// Golem Cloud:
// 1: "{component_name}" -> will resolve in current account and project
// 2: "{project_name}/{component_name}" -> will resolve in current account
// 3: "{account_id}/{project_name}/{component_name}"
//
//	resolve-component-id: func(component-reference: string) -> option<component-id>
//
//go:nosplit
func ResolveComponentID(componentReference string) (result cm.Option[ComponentID]) {
	componentReference0, componentReference1 := cm.LowerString(componentReference)
	wasmimport_ResolveComponentID((*uint8)(componentReference0), (uint32)(componentReference1), &result)
	return
}

// ResolveWorkerID represents the imported function "resolve-worker-id".
//
// Get the worker-id for a given component and worker name.
// Returns none when no component for the specified reference exists.
//
//	resolve-worker-id: func(component-reference: string, worker-name: string) -> option<worker-id>
//
//go:nosplit
func ResolveWorkerID(componentReference string, workerName string) (result cm.Option[WorkerID]) {
	componentReference0, componentReference1 := cm.LowerString(componentReference)
	workerName0, workerName1 := cm.LowerString(workerName)
	wasmimport_ResolveWorkerID((*uint8)(componentReference0), (uint32)(componentReference1), (*uint8)(workerName0), (uint32)(workerName1), &result)
	return
}

// ResolveWorkerIDStrict represents the imported function "resolve-worker-id-strict".
//
// Get the worker-id for a given component and worker name.
// Returns none when no component for the specified component-reference or no worker
// with the specified worker-name exists.
//
//	resolve-worker-id-strict: func(component-reference: string, worker-name: string)
//	-> option<worker-id>
//
//go:nosplit
func ResolveWorkerIDStrict(componentReference string, workerName string) (result cm.Option[WorkerID]) {
	componentReference0, componentReference1 := cm.LowerString(componentReference)
	workerName0, workerName1 := cm.LowerString(workerName)
	wasmimport_ResolveWorkerIDStrict((*uint8)(componentReference0), (uint32)(componentReference1), (*uint8)(workerName0), (uint32)(workerName1), &result)
	return
}

// Fork represents the imported function "fork".
//
// Forks the current worker at the current execution point. The new worker gets the
// `new-name` worker name,
// and this worker continues running as well. The return value is going to be different
// in this worker and
// the forked worker.
//
//	fork: func(new-name: string) -> fork-result
//
//go:nosplit
func Fork(newName string) (result ForkResult) {
	newName0, newName1 := cm.LowerString(newName)
	result0 := wasmimport_Fork((*uint8)(newName0), (uint32)(newName1))
	result = (ForkResult)((uint32)(result0))
	return
}
