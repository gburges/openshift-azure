// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package mock_cluster is a generated GoMock package.
package mock_cluster

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	api "github.com/openshift/openshift-azure/pkg/api"
)

// MockUpgrader is a mock of Upgrader interface
type MockUpgrader struct {
	ctrl     *gomock.Controller
	recorder *MockUpgraderMockRecorder
}

// MockUpgraderMockRecorder is the mock recorder for MockUpgrader
type MockUpgraderMockRecorder struct {
	mock *MockUpgrader
}

// NewMockUpgrader creates a new mock instance
func NewMockUpgrader(ctrl *gomock.Controller) *MockUpgrader {
	mock := &MockUpgrader{ctrl: ctrl}
	mock.recorder = &MockUpgraderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpgrader) EXPECT() *MockUpgraderMockRecorder {
	return m.recorder
}

// CreateClients mocks base method
func (m *MockUpgrader) CreateClients(ctx context.Context, cs *api.OpenShiftManagedCluster, disableKeepAlives bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClients", ctx, cs, disableKeepAlives)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateClients indicates an expected call of CreateClients
func (mr *MockUpgraderMockRecorder) CreateClients(ctx, cs, disableKeepAlives interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClients", reflect.TypeOf((*MockUpgrader)(nil).CreateClients), ctx, cs, disableKeepAlives)
}

// CreateConfigStorageAccount mocks base method
func (m *MockUpgrader) CreateConfigStorageAccount(ctx context.Context, cs *api.OpenShiftManagedCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateConfigStorageAccount", ctx, cs)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateConfigStorageAccount indicates an expected call of CreateConfigStorageAccount
func (mr *MockUpgraderMockRecorder) CreateConfigStorageAccount(ctx, cs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConfigStorageAccount", reflect.TypeOf((*MockUpgrader)(nil).CreateConfigStorageAccount), ctx, cs)
}

// EnrichCSFromVault mocks base method
func (m *MockUpgrader) EnrichCSFromVault(ctx context.Context, cs *api.OpenShiftManagedCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnrichCSFromVault", ctx, cs)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnrichCSFromVault indicates an expected call of EnrichCSFromVault
func (mr *MockUpgraderMockRecorder) EnrichCSFromVault(ctx, cs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnrichCSFromVault", reflect.TypeOf((*MockUpgrader)(nil).EnrichCSFromVault), ctx, cs)
}

// Initialize mocks base method
func (m *MockUpgrader) Initialize(ctx context.Context, cs *api.OpenShiftManagedCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", ctx, cs)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize
func (mr *MockUpgraderMockRecorder) Initialize(ctx, cs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockUpgrader)(nil).Initialize), ctx, cs)
}

// InitializeUpdateBlob mocks base method
func (m *MockUpgrader) InitializeUpdateBlob(cs *api.OpenShiftManagedCluster, suffix string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeUpdateBlob", cs, suffix)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitializeUpdateBlob indicates an expected call of InitializeUpdateBlob
func (mr *MockUpgraderMockRecorder) InitializeUpdateBlob(cs, suffix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeUpdateBlob", reflect.TypeOf((*MockUpgrader)(nil).InitializeUpdateBlob), cs, suffix)
}

// WaitForHealthzStatusOk mocks base method
func (m *MockUpgrader) WaitForHealthzStatusOk(ctx context.Context, cs *api.OpenShiftManagedCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForHealthzStatusOk", ctx, cs)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForHealthzStatusOk indicates an expected call of WaitForHealthzStatusOk
func (mr *MockUpgraderMockRecorder) WaitForHealthzStatusOk(ctx, cs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForHealthzStatusOk", reflect.TypeOf((*MockUpgrader)(nil).WaitForHealthzStatusOk), ctx, cs)
}

// HealthCheck mocks base method
func (m *MockUpgrader) HealthCheck(ctx context.Context, cs *api.OpenShiftManagedCluster) *api.PluginError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck", ctx, cs)
	ret0, _ := ret[0].(*api.PluginError)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck
func (mr *MockUpgraderMockRecorder) HealthCheck(ctx, cs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockUpgrader)(nil).HealthCheck), ctx, cs)
}

// SortedAgentPoolProfilesForRole mocks base method
func (m *MockUpgrader) SortedAgentPoolProfilesForRole(cs *api.OpenShiftManagedCluster, role api.AgentPoolProfileRole) []api.AgentPoolProfile {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SortedAgentPoolProfilesForRole", cs, role)
	ret0, _ := ret[0].([]api.AgentPoolProfile)
	return ret0
}

// SortedAgentPoolProfilesForRole indicates an expected call of SortedAgentPoolProfilesForRole
func (mr *MockUpgraderMockRecorder) SortedAgentPoolProfilesForRole(cs, role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SortedAgentPoolProfilesForRole", reflect.TypeOf((*MockUpgrader)(nil).SortedAgentPoolProfilesForRole), cs, role)
}

// WaitForNodesInAgentPoolProfile mocks base method
func (m *MockUpgrader) WaitForNodesInAgentPoolProfile(ctx context.Context, cs *api.OpenShiftManagedCluster, app *api.AgentPoolProfile, suffix string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForNodesInAgentPoolProfile", ctx, cs, app, suffix)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForNodesInAgentPoolProfile indicates an expected call of WaitForNodesInAgentPoolProfile
func (mr *MockUpgraderMockRecorder) WaitForNodesInAgentPoolProfile(ctx, cs, app, suffix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForNodesInAgentPoolProfile", reflect.TypeOf((*MockUpgrader)(nil).WaitForNodesInAgentPoolProfile), ctx, cs, app, suffix)
}

// UpdateMasterAgentPool mocks base method
func (m *MockUpgrader) UpdateMasterAgentPool(ctx context.Context, cs *api.OpenShiftManagedCluster, app *api.AgentPoolProfile) *api.PluginError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMasterAgentPool", ctx, cs, app)
	ret0, _ := ret[0].(*api.PluginError)
	return ret0
}

// UpdateMasterAgentPool indicates an expected call of UpdateMasterAgentPool
func (mr *MockUpgraderMockRecorder) UpdateMasterAgentPool(ctx, cs, app interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMasterAgentPool", reflect.TypeOf((*MockUpgrader)(nil).UpdateMasterAgentPool), ctx, cs, app)
}

// UpdateWorkerAgentPool mocks base method
func (m *MockUpgrader) UpdateWorkerAgentPool(ctx context.Context, cs *api.OpenShiftManagedCluster, app *api.AgentPoolProfile, suffix string) *api.PluginError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWorkerAgentPool", ctx, cs, app, suffix)
	ret0, _ := ret[0].(*api.PluginError)
	return ret0
}

// UpdateWorkerAgentPool indicates an expected call of UpdateWorkerAgentPool
func (mr *MockUpgraderMockRecorder) UpdateWorkerAgentPool(ctx, cs, app, suffix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkerAgentPool", reflect.TypeOf((*MockUpgrader)(nil).UpdateWorkerAgentPool), ctx, cs, app, suffix)
}

// EtcdBlobExists mocks base method
func (m *MockUpgrader) EtcdBlobExists(ctx context.Context, blobName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EtcdBlobExists", ctx, blobName)
	ret0, _ := ret[0].(error)
	return ret0
}

// EtcdBlobExists indicates an expected call of EtcdBlobExists
func (mr *MockUpgraderMockRecorder) EtcdBlobExists(ctx, blobName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EtcdBlobExists", reflect.TypeOf((*MockUpgrader)(nil).EtcdBlobExists), ctx, blobName)
}

// EtcdRestoreDeleteMasterScaleSet mocks base method
func (m *MockUpgrader) EtcdRestoreDeleteMasterScaleSet(ctx context.Context, cs *api.OpenShiftManagedCluster) *api.PluginError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EtcdRestoreDeleteMasterScaleSet", ctx, cs)
	ret0, _ := ret[0].(*api.PluginError)
	return ret0
}

// EtcdRestoreDeleteMasterScaleSet indicates an expected call of EtcdRestoreDeleteMasterScaleSet
func (mr *MockUpgraderMockRecorder) EtcdRestoreDeleteMasterScaleSet(ctx, cs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EtcdRestoreDeleteMasterScaleSet", reflect.TypeOf((*MockUpgrader)(nil).EtcdRestoreDeleteMasterScaleSet), ctx, cs)
}

// EtcdRestoreDeleteMasterScaleSetHashes mocks base method
func (m *MockUpgrader) EtcdRestoreDeleteMasterScaleSetHashes(ctx context.Context, cs *api.OpenShiftManagedCluster) *api.PluginError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EtcdRestoreDeleteMasterScaleSetHashes", ctx, cs)
	ret0, _ := ret[0].(*api.PluginError)
	return ret0
}

// EtcdRestoreDeleteMasterScaleSetHashes indicates an expected call of EtcdRestoreDeleteMasterScaleSetHashes
func (mr *MockUpgraderMockRecorder) EtcdRestoreDeleteMasterScaleSetHashes(ctx, cs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EtcdRestoreDeleteMasterScaleSetHashes", reflect.TypeOf((*MockUpgrader)(nil).EtcdRestoreDeleteMasterScaleSetHashes), ctx, cs)
}

// ResetUpdateBlob mocks base method
func (m *MockUpgrader) ResetUpdateBlob(cs *api.OpenShiftManagedCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetUpdateBlob", cs)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetUpdateBlob indicates an expected call of ResetUpdateBlob
func (mr *MockUpgraderMockRecorder) ResetUpdateBlob(cs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetUpdateBlob", reflect.TypeOf((*MockUpgrader)(nil).ResetUpdateBlob), cs)
}

// Reimage mocks base method
func (m *MockUpgrader) Reimage(ctx context.Context, cs *api.OpenShiftManagedCluster, scaleset, instanceID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reimage", ctx, cs, scaleset, instanceID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reimage indicates an expected call of Reimage
func (mr *MockUpgraderMockRecorder) Reimage(ctx, cs, scaleset, instanceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reimage", reflect.TypeOf((*MockUpgrader)(nil).Reimage), ctx, cs, scaleset, instanceID)
}

// ListVMHostnames mocks base method
func (m *MockUpgrader) ListVMHostnames(ctx context.Context, cs *api.OpenShiftManagedCluster) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVMHostnames", ctx, cs)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVMHostnames indicates an expected call of ListVMHostnames
func (mr *MockUpgraderMockRecorder) ListVMHostnames(ctx, cs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVMHostnames", reflect.TypeOf((*MockUpgrader)(nil).ListVMHostnames), ctx, cs)
}

// RunCommand mocks base method
func (m *MockUpgrader) RunCommand(ctx context.Context, cs *api.OpenShiftManagedCluster, scaleset, instanceID, command string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunCommand", ctx, cs, scaleset, instanceID, command)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunCommand indicates an expected call of RunCommand
func (mr *MockUpgraderMockRecorder) RunCommand(ctx, cs, scaleset, instanceID, command interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCommand", reflect.TypeOf((*MockUpgrader)(nil).RunCommand), ctx, cs, scaleset, instanceID, command)
}
