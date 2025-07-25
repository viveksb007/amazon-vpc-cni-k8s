// Code generated by MockGen. DO NOT EDIT.
// Source: driver.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	driver "github.com/aws/amazon-vpc-cni-k8s/cmd/routed-eni-cni-plugin/driver"
	sgpp "github.com/aws/amazon-vpc-cni-k8s/pkg/sgpp"
	logger "github.com/aws/amazon-vpc-cni-k8s/pkg/utils/logger"
	gomock "github.com/golang/mock/gomock"
)

// MockNetworkAPIs is a mock of NetworkAPIs interface.
type MockNetworkAPIs struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkAPIsMockRecorder
}

// MockNetworkAPIsMockRecorder is the mock recorder for MockNetworkAPIs.
type MockNetworkAPIsMockRecorder struct {
	mock *MockNetworkAPIs
}

// NewMockNetworkAPIs creates a new mock instance.
func NewMockNetworkAPIs(ctrl *gomock.Controller) *MockNetworkAPIs {
	mock := &MockNetworkAPIs{ctrl: ctrl}
	mock.recorder = &MockNetworkAPIsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkAPIs) EXPECT() *MockNetworkAPIsMockRecorder {
	return m.recorder
}

// SetupBranchENIPodNetwork mocks base method.
func (m *MockNetworkAPIs) SetupBranchENIPodNetwork(vethMetadata driver.VirtualInterfaceMetadata, netnsPath string, vlanID int, eniMAC, subnetGW string, parentIfIndex, mtu int, podSGEnforcingMode sgpp.EnforcingMode, log logger.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupBranchENIPodNetwork", vethMetadata, netnsPath, vlanID, eniMAC, subnetGW, parentIfIndex, mtu, podSGEnforcingMode, log)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupBranchENIPodNetwork indicates an expected call of SetupBranchENIPodNetwork.
func (mr *MockNetworkAPIsMockRecorder) SetupBranchENIPodNetwork(vethMetadata, netnsPath, vlanID, eniMAC, subnetGW, parentIfIndex, mtu, podSGEnforcingMode, log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupBranchENIPodNetwork", reflect.TypeOf((*MockNetworkAPIs)(nil).SetupBranchENIPodNetwork), vethMetadata, netnsPath, vlanID, eniMAC, subnetGW, parentIfIndex, mtu, podSGEnforcingMode, log)
}

// SetupPodNetwork mocks base method.
func (m *MockNetworkAPIs) SetupPodNetwork(vethMetadata []driver.VirtualInterfaceMetadata, netnsPath string, mtu int, log logger.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupPodNetwork", vethMetadata, netnsPath, mtu, log)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupPodNetwork indicates an expected call of SetupPodNetwork.
func (mr *MockNetworkAPIsMockRecorder) SetupPodNetwork(vethMetadata, netnsPath, mtu, log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupPodNetwork", reflect.TypeOf((*MockNetworkAPIs)(nil).SetupPodNetwork), vethMetadata, netnsPath, mtu, log)
}

// TeardownBranchENIPodNetwork mocks base method.
func (m *MockNetworkAPIs) TeardownBranchENIPodNetwork(vethMetadata driver.VirtualInterfaceMetadata, vlanID int, podSGEnforcingMode sgpp.EnforcingMode, log logger.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeardownBranchENIPodNetwork", vethMetadata, vlanID, podSGEnforcingMode, log)
	ret0, _ := ret[0].(error)
	return ret0
}

// TeardownBranchENIPodNetwork indicates an expected call of TeardownBranchENIPodNetwork.
func (mr *MockNetworkAPIsMockRecorder) TeardownBranchENIPodNetwork(vethMetadata, vlanID, podSGEnforcingMode, log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeardownBranchENIPodNetwork", reflect.TypeOf((*MockNetworkAPIs)(nil).TeardownBranchENIPodNetwork), vethMetadata, vlanID, podSGEnforcingMode, log)
}

// TeardownPodNetwork mocks base method.
func (m *MockNetworkAPIs) TeardownPodNetwork(vethMetadata []driver.VirtualInterfaceMetadata, log logger.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeardownPodNetwork", vethMetadata, log)
	ret0, _ := ret[0].(error)
	return ret0
}

// TeardownPodNetwork indicates an expected call of TeardownPodNetwork.
func (mr *MockNetworkAPIsMockRecorder) TeardownPodNetwork(vethMetadata, log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeardownPodNetwork", reflect.TypeOf((*MockNetworkAPIs)(nil).TeardownPodNetwork), vethMetadata, log)
}
