// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cockroachdb/cockroach/pkg/sql/schemachanger/scexec (interfaces: Catalog,Dependencies,Backfiller,BackfillTracker,IndexSpanSplitter,PeriodicProgressFlusher)

// Package scexec_test is a generated GoMock package.
package scexec_test

import (
	context "context"
	reflect "reflect"

	security "github.com/cockroachdb/cockroach/pkg/security"
	catalog "github.com/cockroachdb/cockroach/pkg/sql/catalog"
	scexec "github.com/cockroachdb/cockroach/pkg/sql/schemachanger/scexec"
	scmutationexec "github.com/cockroachdb/cockroach/pkg/sql/schemachanger/scexec/scmutationexec"
	catid "github.com/cockroachdb/cockroach/pkg/sql/sem/catid"
	gomock "github.com/golang/mock/gomock"
)

// MockCatalog is a mock of Catalog interface.
type MockCatalog struct {
	ctrl     *gomock.Controller
	recorder *MockCatalogMockRecorder
}

// MockCatalogMockRecorder is the mock recorder for MockCatalog.
type MockCatalogMockRecorder struct {
	mock *MockCatalog
}

// NewMockCatalog creates a new mock instance.
func NewMockCatalog(ctrl *gomock.Controller) *MockCatalog {
	mock := &MockCatalog{ctrl: ctrl}
	mock.recorder = &MockCatalogMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCatalog) EXPECT() *MockCatalogMockRecorder {
	return m.recorder
}

// AddSyntheticDescriptor mocks base method.
func (m *MockCatalog) AddSyntheticDescriptor(arg0 catalog.Descriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddSyntheticDescriptor", arg0)
}

// AddSyntheticDescriptor indicates an expected call of AddSyntheticDescriptor.
func (mr *MockCatalogMockRecorder) AddSyntheticDescriptor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSyntheticDescriptor", reflect.TypeOf((*MockCatalog)(nil).AddSyntheticDescriptor), arg0)
}

// GetFullyQualifiedName mocks base method.
func (m *MockCatalog) GetFullyQualifiedName(arg0 context.Context, arg1 catid.DescID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullyQualifiedName", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullyQualifiedName indicates an expected call of GetFullyQualifiedName.
func (mr *MockCatalogMockRecorder) GetFullyQualifiedName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullyQualifiedName", reflect.TypeOf((*MockCatalog)(nil).GetFullyQualifiedName), arg0, arg1)
}

// MustReadImmutableDescriptor mocks base method.
func (m *MockCatalog) MustReadImmutableDescriptor(arg0 context.Context, arg1 catid.DescID) (catalog.Descriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MustReadImmutableDescriptor", arg0, arg1)
	ret0, _ := ret[0].(catalog.Descriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MustReadImmutableDescriptor indicates an expected call of MustReadImmutableDescriptor.
func (mr *MockCatalogMockRecorder) MustReadImmutableDescriptor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustReadImmutableDescriptor", reflect.TypeOf((*MockCatalog)(nil).MustReadImmutableDescriptor), arg0, arg1)
}

// MustReadMutableDescriptor mocks base method.
func (m *MockCatalog) MustReadMutableDescriptor(arg0 context.Context, arg1 catid.DescID) (catalog.MutableDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MustReadMutableDescriptor", arg0, arg1)
	ret0, _ := ret[0].(catalog.MutableDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MustReadMutableDescriptor indicates an expected call of MustReadMutableDescriptor.
func (mr *MockCatalogMockRecorder) MustReadMutableDescriptor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustReadMutableDescriptor", reflect.TypeOf((*MockCatalog)(nil).MustReadMutableDescriptor), arg0, arg1)
}

// NewCatalogChangeBatcher mocks base method.
func (m *MockCatalog) NewCatalogChangeBatcher() scexec.CatalogChangeBatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCatalogChangeBatcher")
	ret0, _ := ret[0].(scexec.CatalogChangeBatcher)
	return ret0
}

// NewCatalogChangeBatcher indicates an expected call of NewCatalogChangeBatcher.
func (mr *MockCatalogMockRecorder) NewCatalogChangeBatcher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCatalogChangeBatcher", reflect.TypeOf((*MockCatalog)(nil).NewCatalogChangeBatcher))
}

// RemoveSyntheticDescriptor mocks base method.
func (m *MockCatalog) RemoveSyntheticDescriptor(arg0 catid.DescID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveSyntheticDescriptor", arg0)
}

// RemoveSyntheticDescriptor indicates an expected call of RemoveSyntheticDescriptor.
func (mr *MockCatalogMockRecorder) RemoveSyntheticDescriptor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSyntheticDescriptor", reflect.TypeOf((*MockCatalog)(nil).RemoveSyntheticDescriptor), arg0)
}

// MockDependencies is a mock of Dependencies interface.
type MockDependencies struct {
	ctrl     *gomock.Controller
	recorder *MockDependenciesMockRecorder
}

// MockDependenciesMockRecorder is the mock recorder for MockDependencies.
type MockDependenciesMockRecorder struct {
	mock *MockDependencies
}

// NewMockDependencies creates a new mock instance.
func NewMockDependencies(ctrl *gomock.Controller) *MockDependencies {
	mock := &MockDependencies{ctrl: ctrl}
	mock.recorder = &MockDependenciesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDependencies) EXPECT() *MockDependenciesMockRecorder {
	return m.recorder
}

// BackfillProgressTracker mocks base method.
func (m *MockDependencies) BackfillProgressTracker() scexec.BackfillTracker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BackfillProgressTracker")
	ret0, _ := ret[0].(scexec.BackfillTracker)
	return ret0
}

// BackfillProgressTracker indicates an expected call of BackfillProgressTracker.
func (mr *MockDependenciesMockRecorder) BackfillProgressTracker() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackfillProgressTracker", reflect.TypeOf((*MockDependencies)(nil).BackfillProgressTracker))
}

// Catalog mocks base method.
func (m *MockDependencies) Catalog() scexec.Catalog {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Catalog")
	ret0, _ := ret[0].(scexec.Catalog)
	return ret0
}

// Catalog indicates an expected call of Catalog.
func (mr *MockDependenciesMockRecorder) Catalog() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Catalog", reflect.TypeOf((*MockDependencies)(nil).Catalog))
}

// EventLogger mocks base method.
func (m *MockDependencies) EventLogger() scexec.EventLogger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventLogger")
	ret0, _ := ret[0].(scexec.EventLogger)
	return ret0
}

// EventLogger indicates an expected call of EventLogger.
func (mr *MockDependenciesMockRecorder) EventLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventLogger", reflect.TypeOf((*MockDependencies)(nil).EventLogger))
}

// IndexBackfiller mocks base method.
func (m *MockDependencies) IndexBackfiller() scexec.Backfiller {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexBackfiller")
	ret0, _ := ret[0].(scexec.Backfiller)
	return ret0
}

// IndexBackfiller indicates an expected call of IndexBackfiller.
func (mr *MockDependenciesMockRecorder) IndexBackfiller() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexBackfiller", reflect.TypeOf((*MockDependencies)(nil).IndexBackfiller))
}

// IndexSpanSplitter mocks base method.
func (m *MockDependencies) IndexSpanSplitter() scexec.IndexSpanSplitter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexSpanSplitter")
	ret0, _ := ret[0].(scexec.IndexSpanSplitter)
	return ret0
}

// IndexSpanSplitter indicates an expected call of IndexSpanSplitter.
func (mr *MockDependenciesMockRecorder) IndexSpanSplitter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexSpanSplitter", reflect.TypeOf((*MockDependencies)(nil).IndexSpanSplitter))
}

// IndexValidator mocks base method.
func (m *MockDependencies) IndexValidator() scexec.IndexValidator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexValidator")
	ret0, _ := ret[0].(scexec.IndexValidator)
	return ret0
}

// IndexValidator indicates an expected call of IndexValidator.
func (mr *MockDependenciesMockRecorder) IndexValidator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexValidator", reflect.TypeOf((*MockDependencies)(nil).IndexValidator))
}

// Partitioner mocks base method.
func (m *MockDependencies) Partitioner() scmutationexec.Partitioner {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Partitioner")
	ret0, _ := ret[0].(scmutationexec.Partitioner)
	return ret0
}

// Partitioner indicates an expected call of Partitioner.
func (mr *MockDependenciesMockRecorder) Partitioner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Partitioner", reflect.TypeOf((*MockDependencies)(nil).Partitioner))
}

// PeriodicProgressFlusher mocks base method.
func (m *MockDependencies) PeriodicProgressFlusher() scexec.PeriodicProgressFlusher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeriodicProgressFlusher")
	ret0, _ := ret[0].(scexec.PeriodicProgressFlusher)
	return ret0
}

// PeriodicProgressFlusher indicates an expected call of PeriodicProgressFlusher.
func (mr *MockDependenciesMockRecorder) PeriodicProgressFlusher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeriodicProgressFlusher", reflect.TypeOf((*MockDependencies)(nil).PeriodicProgressFlusher))
}

// Statements mocks base method.
func (m *MockDependencies) Statements() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Statements")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Statements indicates an expected call of Statements.
func (mr *MockDependenciesMockRecorder) Statements() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Statements", reflect.TypeOf((*MockDependencies)(nil).Statements))
}

// TransactionalJobRegistry mocks base method.
func (m *MockDependencies) TransactionalJobRegistry() scexec.TransactionalJobRegistry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionalJobRegistry")
	ret0, _ := ret[0].(scexec.TransactionalJobRegistry)
	return ret0
}

// TransactionalJobRegistry indicates an expected call of TransactionalJobRegistry.
func (mr *MockDependenciesMockRecorder) TransactionalJobRegistry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionalJobRegistry", reflect.TypeOf((*MockDependencies)(nil).TransactionalJobRegistry))
}

// User mocks base method.
func (m *MockDependencies) User() security.SQLUsername {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "User")
	ret0, _ := ret[0].(security.SQLUsername)
	return ret0
}

// User indicates an expected call of User.
func (mr *MockDependenciesMockRecorder) User() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockDependencies)(nil).User))
}

// MockBackfiller is a mock of Backfiller interface.
type MockBackfiller struct {
	ctrl     *gomock.Controller
	recorder *MockBackfillerMockRecorder
}

// MockBackfillerMockRecorder is the mock recorder for MockBackfiller.
type MockBackfillerMockRecorder struct {
	mock *MockBackfiller
}

// NewMockBackfiller creates a new mock instance.
func NewMockBackfiller(ctrl *gomock.Controller) *MockBackfiller {
	mock := &MockBackfiller{ctrl: ctrl}
	mock.recorder = &MockBackfillerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackfiller) EXPECT() *MockBackfillerMockRecorder {
	return m.recorder
}

// BackfillIndex mocks base method.
func (m *MockBackfiller) BackfillIndex(arg0 context.Context, arg1 scexec.BackfillProgress, arg2 scexec.BackfillProgressWriter, arg3 catalog.TableDescriptor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BackfillIndex", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// BackfillIndex indicates an expected call of BackfillIndex.
func (mr *MockBackfillerMockRecorder) BackfillIndex(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackfillIndex", reflect.TypeOf((*MockBackfiller)(nil).BackfillIndex), arg0, arg1, arg2, arg3)
}

// MaybePrepareDestIndexesForBackfill mocks base method.
func (m *MockBackfiller) MaybePrepareDestIndexesForBackfill(arg0 context.Context, arg1 scexec.BackfillProgress, arg2 catalog.TableDescriptor) (scexec.BackfillProgress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaybePrepareDestIndexesForBackfill", arg0, arg1, arg2)
	ret0, _ := ret[0].(scexec.BackfillProgress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaybePrepareDestIndexesForBackfill indicates an expected call of MaybePrepareDestIndexesForBackfill.
func (mr *MockBackfillerMockRecorder) MaybePrepareDestIndexesForBackfill(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybePrepareDestIndexesForBackfill", reflect.TypeOf((*MockBackfiller)(nil).MaybePrepareDestIndexesForBackfill), arg0, arg1, arg2)
}

// MockBackfillTracker is a mock of BackfillTracker interface.
type MockBackfillTracker struct {
	ctrl     *gomock.Controller
	recorder *MockBackfillTrackerMockRecorder
}

// MockBackfillTrackerMockRecorder is the mock recorder for MockBackfillTracker.
type MockBackfillTrackerMockRecorder struct {
	mock *MockBackfillTracker
}

// NewMockBackfillTracker creates a new mock instance.
func NewMockBackfillTracker(ctrl *gomock.Controller) *MockBackfillTracker {
	mock := &MockBackfillTracker{ctrl: ctrl}
	mock.recorder = &MockBackfillTrackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackfillTracker) EXPECT() *MockBackfillTrackerMockRecorder {
	return m.recorder
}

// FlushCheckpoint mocks base method.
func (m *MockBackfillTracker) FlushCheckpoint(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushCheckpoint", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FlushCheckpoint indicates an expected call of FlushCheckpoint.
func (mr *MockBackfillTrackerMockRecorder) FlushCheckpoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushCheckpoint", reflect.TypeOf((*MockBackfillTracker)(nil).FlushCheckpoint), arg0)
}

// FlushFractionCompleted mocks base method.
func (m *MockBackfillTracker) FlushFractionCompleted(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushFractionCompleted", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FlushFractionCompleted indicates an expected call of FlushFractionCompleted.
func (mr *MockBackfillTrackerMockRecorder) FlushFractionCompleted(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushFractionCompleted", reflect.TypeOf((*MockBackfillTracker)(nil).FlushFractionCompleted), arg0)
}

// GetBackfillProgress mocks base method.
func (m *MockBackfillTracker) GetBackfillProgress(arg0 context.Context, arg1 scexec.Backfill) (scexec.BackfillProgress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackfillProgress", arg0, arg1)
	ret0, _ := ret[0].(scexec.BackfillProgress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackfillProgress indicates an expected call of GetBackfillProgress.
func (mr *MockBackfillTrackerMockRecorder) GetBackfillProgress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackfillProgress", reflect.TypeOf((*MockBackfillTracker)(nil).GetBackfillProgress), arg0, arg1)
}

// SetBackfillProgress mocks base method.
func (m *MockBackfillTracker) SetBackfillProgress(arg0 context.Context, arg1 scexec.BackfillProgress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBackfillProgress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetBackfillProgress indicates an expected call of SetBackfillProgress.
func (mr *MockBackfillTrackerMockRecorder) SetBackfillProgress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBackfillProgress", reflect.TypeOf((*MockBackfillTracker)(nil).SetBackfillProgress), arg0, arg1)
}

// MockIndexSpanSplitter is a mock of IndexSpanSplitter interface.
type MockIndexSpanSplitter struct {
	ctrl     *gomock.Controller
	recorder *MockIndexSpanSplitterMockRecorder
}

// MockIndexSpanSplitterMockRecorder is the mock recorder for MockIndexSpanSplitter.
type MockIndexSpanSplitterMockRecorder struct {
	mock *MockIndexSpanSplitter
}

// NewMockIndexSpanSplitter creates a new mock instance.
func NewMockIndexSpanSplitter(ctrl *gomock.Controller) *MockIndexSpanSplitter {
	mock := &MockIndexSpanSplitter{ctrl: ctrl}
	mock.recorder = &MockIndexSpanSplitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIndexSpanSplitter) EXPECT() *MockIndexSpanSplitterMockRecorder {
	return m.recorder
}

// MaybeSplitIndexSpans mocks base method.
func (m *MockIndexSpanSplitter) MaybeSplitIndexSpans(arg0 context.Context, arg1 catalog.TableDescriptor, arg2 catalog.Index) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaybeSplitIndexSpans", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// MaybeSplitIndexSpans indicates an expected call of MaybeSplitIndexSpans.
func (mr *MockIndexSpanSplitterMockRecorder) MaybeSplitIndexSpans(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybeSplitIndexSpans", reflect.TypeOf((*MockIndexSpanSplitter)(nil).MaybeSplitIndexSpans), arg0, arg1, arg2)
}

// MockPeriodicProgressFlusher is a mock of PeriodicProgressFlusher interface.
type MockPeriodicProgressFlusher struct {
	ctrl     *gomock.Controller
	recorder *MockPeriodicProgressFlusherMockRecorder
}

// MockPeriodicProgressFlusherMockRecorder is the mock recorder for MockPeriodicProgressFlusher.
type MockPeriodicProgressFlusherMockRecorder struct {
	mock *MockPeriodicProgressFlusher
}

// NewMockPeriodicProgressFlusher creates a new mock instance.
func NewMockPeriodicProgressFlusher(ctrl *gomock.Controller) *MockPeriodicProgressFlusher {
	mock := &MockPeriodicProgressFlusher{ctrl: ctrl}
	mock.recorder = &MockPeriodicProgressFlusherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeriodicProgressFlusher) EXPECT() *MockPeriodicProgressFlusherMockRecorder {
	return m.recorder
}

// StartPeriodicUpdates mocks base method.
func (m *MockPeriodicProgressFlusher) StartPeriodicUpdates(arg0 context.Context, arg1 scexec.BackfillProgressFlusher) func() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartPeriodicUpdates", arg0, arg1)
	ret0, _ := ret[0].(func() error)
	return ret0
}

// StartPeriodicUpdates indicates an expected call of StartPeriodicUpdates.
func (mr *MockPeriodicProgressFlusherMockRecorder) StartPeriodicUpdates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartPeriodicUpdates", reflect.TypeOf((*MockPeriodicProgressFlusher)(nil).StartPeriodicUpdates), arg0, arg1)
}
