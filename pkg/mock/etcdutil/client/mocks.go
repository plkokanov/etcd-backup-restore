// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/etcd-backup-restore/pkg/etcdutil/client (interfaces: Factory,ClusterCloser,KVCloser,MaintenanceCloser)

// Package client is a generated GoMock package.
package client

import (
	context "context"
	io "io"
	reflect "reflect"

	client "github.com/gardener/etcd-backup-restore/pkg/etcdutil/client"
	gomock "github.com/golang/mock/gomock"
	clientv3 "go.etcd.io/etcd/clientv3"
)

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// NewCluster mocks base method.
func (m *MockFactory) NewCluster() (client.ClusterCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCluster")
	ret0, _ := ret[0].(client.ClusterCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewCluster indicates an expected call of NewCluster.
func (mr *MockFactoryMockRecorder) NewCluster() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCluster", reflect.TypeOf((*MockFactory)(nil).NewCluster))
}

// NewKV mocks base method.
func (m *MockFactory) NewKV() (client.KVCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewKV")
	ret0, _ := ret[0].(client.KVCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewKV indicates an expected call of NewKV.
func (mr *MockFactoryMockRecorder) NewKV() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewKV", reflect.TypeOf((*MockFactory)(nil).NewKV))
}

// NewMaintenance mocks base method.
func (m *MockFactory) NewMaintenance() (client.MaintenanceCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewMaintenance")
	ret0, _ := ret[0].(client.MaintenanceCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewMaintenance indicates an expected call of NewMaintenance.
func (mr *MockFactoryMockRecorder) NewMaintenance() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMaintenance", reflect.TypeOf((*MockFactory)(nil).NewMaintenance))
}

// NewWatcher mocks base method.
func (m *MockFactory) NewWatcher() (clientv3.Watcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewWatcher")
	ret0, _ := ret[0].(clientv3.Watcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewWatcher indicates an expected call of NewWatcher.
func (mr *MockFactoryMockRecorder) NewWatcher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewWatcher", reflect.TypeOf((*MockFactory)(nil).NewWatcher))
}

// MockClusterCloser is a mock of ClusterCloser interface.
type MockClusterCloser struct {
	ctrl     *gomock.Controller
	recorder *MockClusterCloserMockRecorder
}

// MockClusterCloserMockRecorder is the mock recorder for MockClusterCloser.
type MockClusterCloserMockRecorder struct {
	mock *MockClusterCloser
}

// NewMockClusterCloser creates a new mock instance.
func NewMockClusterCloser(ctrl *gomock.Controller) *MockClusterCloser {
	mock := &MockClusterCloser{ctrl: ctrl}
	mock.recorder = &MockClusterCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterCloser) EXPECT() *MockClusterCloserMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockClusterCloser) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClusterCloserMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClusterCloser)(nil).Close))
}

// MemberAdd mocks base method.
func (m *MockClusterCloser) MemberAdd(arg0 context.Context, arg1 []string) (*clientv3.MemberAddResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MemberAdd", arg0, arg1)
	ret0, _ := ret[0].(*clientv3.MemberAddResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MemberAdd indicates an expected call of MemberAdd.
func (mr *MockClusterCloserMockRecorder) MemberAdd(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemberAdd", reflect.TypeOf((*MockClusterCloser)(nil).MemberAdd), arg0, arg1)
}

// MemberAddAsLearner mocks base method.
func (m *MockClusterCloser) MemberAddAsLearner(arg0 context.Context, arg1 []string) (*clientv3.MemberAddResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MemberAddAsLearner", arg0, arg1)
	ret0, _ := ret[0].(*clientv3.MemberAddResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MemberAddAsLearner indicates an expected call of MemberAddAsLearner.
func (mr *MockClusterCloserMockRecorder) MemberAddAsLearner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemberAddAsLearner", reflect.TypeOf((*MockClusterCloser)(nil).MemberAddAsLearner), arg0, arg1)
}

// MemberList mocks base method.
func (m *MockClusterCloser) MemberList(arg0 context.Context) (*clientv3.MemberListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MemberList", arg0)
	ret0, _ := ret[0].(*clientv3.MemberListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MemberList indicates an expected call of MemberList.
func (mr *MockClusterCloserMockRecorder) MemberList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemberList", reflect.TypeOf((*MockClusterCloser)(nil).MemberList), arg0)
}

// MemberPromote mocks base method.
func (m *MockClusterCloser) MemberPromote(arg0 context.Context, arg1 uint64) (*clientv3.MemberPromoteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MemberPromote", arg0, arg1)
	ret0, _ := ret[0].(*clientv3.MemberPromoteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MemberPromote indicates an expected call of MemberPromote.
func (mr *MockClusterCloserMockRecorder) MemberPromote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemberPromote", reflect.TypeOf((*MockClusterCloser)(nil).MemberPromote), arg0, arg1)
}

// MemberRemove mocks base method.
func (m *MockClusterCloser) MemberRemove(arg0 context.Context, arg1 uint64) (*clientv3.MemberRemoveResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MemberRemove", arg0, arg1)
	ret0, _ := ret[0].(*clientv3.MemberRemoveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MemberRemove indicates an expected call of MemberRemove.
func (mr *MockClusterCloserMockRecorder) MemberRemove(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemberRemove", reflect.TypeOf((*MockClusterCloser)(nil).MemberRemove), arg0, arg1)
}

// MemberUpdate mocks base method.
func (m *MockClusterCloser) MemberUpdate(arg0 context.Context, arg1 uint64, arg2 []string) (*clientv3.MemberUpdateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MemberUpdate", arg0, arg1, arg2)
	ret0, _ := ret[0].(*clientv3.MemberUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MemberUpdate indicates an expected call of MemberUpdate.
func (mr *MockClusterCloserMockRecorder) MemberUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemberUpdate", reflect.TypeOf((*MockClusterCloser)(nil).MemberUpdate), arg0, arg1, arg2)
}

// MockKVCloser is a mock of KVCloser interface.
type MockKVCloser struct {
	ctrl     *gomock.Controller
	recorder *MockKVCloserMockRecorder
}

// MockKVCloserMockRecorder is the mock recorder for MockKVCloser.
type MockKVCloserMockRecorder struct {
	mock *MockKVCloser
}

// NewMockKVCloser creates a new mock instance.
func NewMockKVCloser(ctrl *gomock.Controller) *MockKVCloser {
	mock := &MockKVCloser{ctrl: ctrl}
	mock.recorder = &MockKVCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKVCloser) EXPECT() *MockKVCloserMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockKVCloser) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockKVCloserMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockKVCloser)(nil).Close))
}

// Compact mocks base method.
func (m *MockKVCloser) Compact(arg0 context.Context, arg1 int64, arg2 ...clientv3.CompactOption) (*clientv3.CompactResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Compact", varargs...)
	ret0, _ := ret[0].(*clientv3.CompactResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Compact indicates an expected call of Compact.
func (mr *MockKVCloserMockRecorder) Compact(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compact", reflect.TypeOf((*MockKVCloser)(nil).Compact), varargs...)
}

// Delete mocks base method.
func (m *MockKVCloser) Delete(arg0 context.Context, arg1 string, arg2 ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*clientv3.DeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockKVCloserMockRecorder) Delete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockKVCloser)(nil).Delete), varargs...)
}

// Do mocks base method.
func (m *MockKVCloser) Do(arg0 context.Context, arg1 clientv3.Op) (clientv3.OpResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(clientv3.OpResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockKVCloserMockRecorder) Do(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockKVCloser)(nil).Do), arg0, arg1)
}

// Get mocks base method.
func (m *MockKVCloser) Get(arg0 context.Context, arg1 string, arg2 ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*clientv3.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockKVCloserMockRecorder) Get(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKVCloser)(nil).Get), varargs...)
}

// Put mocks base method.
func (m *MockKVCloser) Put(arg0 context.Context, arg1, arg2 string, arg3 ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Put", varargs...)
	ret0, _ := ret[0].(*clientv3.PutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockKVCloserMockRecorder) Put(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockKVCloser)(nil).Put), varargs...)
}

// Txn mocks base method.
func (m *MockKVCloser) Txn(arg0 context.Context) clientv3.Txn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Txn", arg0)
	ret0, _ := ret[0].(clientv3.Txn)
	return ret0
}

// Txn indicates an expected call of Txn.
func (mr *MockKVCloserMockRecorder) Txn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Txn", reflect.TypeOf((*MockKVCloser)(nil).Txn), arg0)
}

// MockMaintenanceCloser is a mock of MaintenanceCloser interface.
type MockMaintenanceCloser struct {
	ctrl     *gomock.Controller
	recorder *MockMaintenanceCloserMockRecorder
}

// MockMaintenanceCloserMockRecorder is the mock recorder for MockMaintenanceCloser.
type MockMaintenanceCloserMockRecorder struct {
	mock *MockMaintenanceCloser
}

// NewMockMaintenanceCloser creates a new mock instance.
func NewMockMaintenanceCloser(ctrl *gomock.Controller) *MockMaintenanceCloser {
	mock := &MockMaintenanceCloser{ctrl: ctrl}
	mock.recorder = &MockMaintenanceCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMaintenanceCloser) EXPECT() *MockMaintenanceCloserMockRecorder {
	return m.recorder
}

// AlarmDisarm mocks base method.
func (m *MockMaintenanceCloser) AlarmDisarm(arg0 context.Context, arg1 *clientv3.AlarmMember) (*clientv3.AlarmResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlarmDisarm", arg0, arg1)
	ret0, _ := ret[0].(*clientv3.AlarmResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlarmDisarm indicates an expected call of AlarmDisarm.
func (mr *MockMaintenanceCloserMockRecorder) AlarmDisarm(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlarmDisarm", reflect.TypeOf((*MockMaintenanceCloser)(nil).AlarmDisarm), arg0, arg1)
}

// AlarmList mocks base method.
func (m *MockMaintenanceCloser) AlarmList(arg0 context.Context) (*clientv3.AlarmResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlarmList", arg0)
	ret0, _ := ret[0].(*clientv3.AlarmResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlarmList indicates an expected call of AlarmList.
func (mr *MockMaintenanceCloserMockRecorder) AlarmList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlarmList", reflect.TypeOf((*MockMaintenanceCloser)(nil).AlarmList), arg0)
}

// Close mocks base method.
func (m *MockMaintenanceCloser) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockMaintenanceCloserMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMaintenanceCloser)(nil).Close))
}

// Defragment mocks base method.
func (m *MockMaintenanceCloser) Defragment(arg0 context.Context, arg1 string) (*clientv3.DefragmentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Defragment", arg0, arg1)
	ret0, _ := ret[0].(*clientv3.DefragmentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Defragment indicates an expected call of Defragment.
func (mr *MockMaintenanceCloserMockRecorder) Defragment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Defragment", reflect.TypeOf((*MockMaintenanceCloser)(nil).Defragment), arg0, arg1)
}

// HashKV mocks base method.
func (m *MockMaintenanceCloser) HashKV(arg0 context.Context, arg1 string, arg2 int64) (*clientv3.HashKVResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashKV", arg0, arg1, arg2)
	ret0, _ := ret[0].(*clientv3.HashKVResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashKV indicates an expected call of HashKV.
func (mr *MockMaintenanceCloserMockRecorder) HashKV(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashKV", reflect.TypeOf((*MockMaintenanceCloser)(nil).HashKV), arg0, arg1, arg2)
}

// MoveLeader mocks base method.
func (m *MockMaintenanceCloser) MoveLeader(arg0 context.Context, arg1 uint64) (*clientv3.MoveLeaderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveLeader", arg0, arg1)
	ret0, _ := ret[0].(*clientv3.MoveLeaderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MoveLeader indicates an expected call of MoveLeader.
func (mr *MockMaintenanceCloserMockRecorder) MoveLeader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveLeader", reflect.TypeOf((*MockMaintenanceCloser)(nil).MoveLeader), arg0, arg1)
}

// Snapshot mocks base method.
func (m *MockMaintenanceCloser) Snapshot(arg0 context.Context) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshot", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Snapshot indicates an expected call of Snapshot.
func (mr *MockMaintenanceCloserMockRecorder) Snapshot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockMaintenanceCloser)(nil).Snapshot), arg0)
}

// Status mocks base method.
func (m *MockMaintenanceCloser) Status(arg0 context.Context, arg1 string) (*clientv3.StatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0, arg1)
	ret0, _ := ret[0].(*clientv3.StatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockMaintenanceCloserMockRecorder) Status(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockMaintenanceCloser)(nil).Status), arg0, arg1)
}
