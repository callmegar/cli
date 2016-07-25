// Automatically generated by MockGen. DO NOT EDIT!
// Source: fromenv.go

package flags

import (
	flag "flag"
	gomock "github.com/golang/mock/gomock"
)

// Mock of FromEnv interface
type MockFromEnv struct {
	ctrl     *gomock.Controller
	recorder *_MockFromEnvRecorder
}

// Recorder for MockFromEnv (not exported)
type _MockFromEnvRecorder struct {
	mock *MockFromEnv
}

func NewMockFromEnv(ctrl *gomock.Controller) *MockFromEnv {
	mock := &MockFromEnv{ctrl: ctrl}
	mock.recorder = &_MockFromEnvRecorder{mock}
	return mock
}

func (_m *MockFromEnv) EXPECT() *_MockFromEnvRecorder {
	return _m.recorder
}

func (_m *MockFromEnv) Prefix() string {
	ret := _m.ctrl.Call(_m, "Prefix")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockFromEnvRecorder) Prefix() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Prefix")
}

func (_m *MockFromEnv) Fill() {
	_m.ctrl.Call(_m, "Fill")
}

func (_mr *_MockFromEnvRecorder) Fill() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fill")
}

func (_m *MockFromEnv) Filled() map[string]string {
	ret := _m.ctrl.Call(_m, "Filled")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

func (_mr *_MockFromEnvRecorder) Filled() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Filled")
}

func (_m *MockFromEnv) AllKeys() []string {
	ret := _m.ctrl.Call(_m, "AllKeys")
	ret0, _ := ret[0].([]string)
	return ret0
}

func (_mr *_MockFromEnvRecorder) AllKeys() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AllKeys")
}

func (_m *MockFromEnv) AllFlags() []*flag.Flag {
	ret := _m.ctrl.Call(_m, "AllFlags")
	ret0, _ := ret[0].([]*flag.Flag)
	return ret0
}

func (_mr *_MockFromEnvRecorder) AllFlags() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AllFlags")
}
