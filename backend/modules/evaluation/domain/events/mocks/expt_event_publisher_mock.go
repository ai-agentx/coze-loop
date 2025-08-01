// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coze-dev/coze-loop/backend/modules/evaluation/domain/events (interfaces: ExptEventPublisher)

// Package mocks is a generated GoMock package.
package mocks

import (
	"context"
	"reflect"
	"time"

	"go.uber.org/mock/gomock"

	"github.com/coze-dev/coze-loop/backend/modules/evaluation/domain/entity"
)

// MockExptEventPublisher is a mock of ExptEventPublisher interface.
type MockExptEventPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockExptEventPublisherMockRecorder
}

// MockExptEventPublisherMockRecorder is the mock recorder for MockExptEventPublisher.
type MockExptEventPublisherMockRecorder struct {
	mock *MockExptEventPublisher
}

// NewMockExptEventPublisher creates a new mock instance.
func NewMockExptEventPublisher(ctrl *gomock.Controller) *MockExptEventPublisher {
	mock := &MockExptEventPublisher{ctrl: ctrl}
	mock.recorder = &MockExptEventPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExptEventPublisher) EXPECT() *MockExptEventPublisherMockRecorder {
	return m.recorder
}

// BatchPublishExptRecordEvalEvent mocks base method.
func (m *MockExptEventPublisher) BatchPublishExptRecordEvalEvent(arg0 context.Context, arg1 []*entity.ExptItemEvalEvent, arg2 *time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchPublishExptRecordEvalEvent", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchPublishExptRecordEvalEvent indicates an expected call of BatchPublishExptRecordEvalEvent.
func (mr *MockExptEventPublisherMockRecorder) BatchPublishExptRecordEvalEvent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchPublishExptRecordEvalEvent", reflect.TypeOf((*MockExptEventPublisher)(nil).BatchPublishExptRecordEvalEvent), arg0, arg1, arg2)
}

// PublishExptAggrCalculateEvent mocks base method.
func (m *MockExptEventPublisher) PublishExptAggrCalculateEvent(arg0 context.Context, arg1 []*entity.AggrCalculateEvent, arg2 *time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishExptAggrCalculateEvent", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishExptAggrCalculateEvent indicates an expected call of PublishExptAggrCalculateEvent.
func (mr *MockExptEventPublisherMockRecorder) PublishExptAggrCalculateEvent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishExptAggrCalculateEvent", reflect.TypeOf((*MockExptEventPublisher)(nil).PublishExptAggrCalculateEvent), arg0, arg1, arg2)
}

// PublishExptOnlineEvalResult mocks base method.
func (m *MockExptEventPublisher) PublishExptOnlineEvalResult(arg0 context.Context, arg1 *entity.OnlineExptEvalResultEvent, arg2 *time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishExptOnlineEvalResult", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishExptOnlineEvalResult indicates an expected call of PublishExptOnlineEvalResult.
func (mr *MockExptEventPublisherMockRecorder) PublishExptOnlineEvalResult(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishExptOnlineEvalResult", reflect.TypeOf((*MockExptEventPublisher)(nil).PublishExptOnlineEvalResult), arg0, arg1, arg2)
}

// PublishExptRecordEvalEvent mocks base method.
func (m *MockExptEventPublisher) PublishExptRecordEvalEvent(arg0 context.Context, arg1 *entity.ExptItemEvalEvent, arg2 *time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishExptRecordEvalEvent", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishExptRecordEvalEvent indicates an expected call of PublishExptRecordEvalEvent.
func (mr *MockExptEventPublisherMockRecorder) PublishExptRecordEvalEvent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishExptRecordEvalEvent", reflect.TypeOf((*MockExptEventPublisher)(nil).PublishExptRecordEvalEvent), arg0, arg1, arg2)
}

// PublishExptScheduleEvent mocks base method.
func (m *MockExptEventPublisher) PublishExptScheduleEvent(arg0 context.Context, arg1 *entity.ExptScheduleEvent, arg2 *time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishExptScheduleEvent", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishExptScheduleEvent indicates an expected call of PublishExptScheduleEvent.
func (mr *MockExptEventPublisherMockRecorder) PublishExptScheduleEvent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishExptScheduleEvent", reflect.TypeOf((*MockExptEventPublisher)(nil).PublishExptScheduleEvent), arg0, arg1, arg2)
}
