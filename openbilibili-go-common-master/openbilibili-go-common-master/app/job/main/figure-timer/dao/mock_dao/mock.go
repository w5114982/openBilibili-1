// Code generated by MockGen. DO NOT EDIT.
// Source: go-common/app/job/main/figure-timer/dao (interfaces: DaoInt)

// Package mock_dao is a generated GoMock package.
package mock_dao

import (
	context "context"
	model "go-common/app/job/main/figure-timer/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDaoInt is a mock of DaoInt interface
type MockDaoInt struct {
	ctrl     *gomock.Controller
	recorder *MockDaoIntMockRecorder
}

// MockDaoIntMockRecorder is the mock recorder for MockDaoInt
type MockDaoIntMockRecorder struct {
	mock *MockDaoInt
}

// NewMockDaoInt creates a new mock instance
func NewMockDaoInt(ctrl *gomock.Controller) *MockDaoInt {
	mock := &MockDaoInt{ctrl: ctrl}
	mock.recorder = &MockDaoIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDaoInt) EXPECT() *MockDaoIntMockRecorder {
	return m.recorder
}

// ActionCounter mocks base method
func (m *MockDaoInt) ActionCounter(arg0 context.Context, arg1, arg2 int64) (*model.ActionCounter, error) {
	ret := m.ctrl.Call(m, "ActionCounter", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.ActionCounter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActionCounter indicates an expected call of ActionCounter
func (mr *MockDaoIntMockRecorder) ActionCounter(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionCounter", reflect.TypeOf((*MockDaoInt)(nil).ActionCounter), arg0, arg1, arg2)
}

// CalcRecords mocks base method
func (m *MockDaoInt) CalcRecords(arg0 context.Context, arg1, arg2, arg3 int64) ([]*model.FigureRecord, error) {
	ret := m.ctrl.Call(m, "CalcRecords", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*model.FigureRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalcRecords indicates an expected call of CalcRecords
func (mr *MockDaoIntMockRecorder) CalcRecords(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalcRecords", reflect.TypeOf((*MockDaoInt)(nil).CalcRecords), arg0, arg1, arg2, arg3)
}

// Close mocks base method
func (m *MockDaoInt) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockDaoIntMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDaoInt)(nil).Close))
}

// Figure mocks base method
func (m *MockDaoInt) Figure(arg0 context.Context, arg1 int64) (*model.Figure, error) {
	ret := m.ctrl.Call(m, "Figure", arg0, arg1)
	ret0, _ := ret[0].(*model.Figure)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Figure indicates an expected call of Figure
func (mr *MockDaoIntMockRecorder) Figure(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Figure", reflect.TypeOf((*MockDaoInt)(nil).Figure), arg0, arg1)
}

// FigureCache mocks base method
func (m *MockDaoInt) FigureCache(arg0 context.Context, arg1 int64) (*model.Figure, error) {
	ret := m.ctrl.Call(m, "FigureCache", arg0, arg1)
	ret0, _ := ret[0].(*model.Figure)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FigureCache indicates an expected call of FigureCache
func (mr *MockDaoIntMockRecorder) FigureCache(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FigureCache", reflect.TypeOf((*MockDaoInt)(nil).FigureCache), arg0, arg1)
}

// Figures mocks base method
func (m *MockDaoInt) Figures(arg0 context.Context, arg1 int64, arg2 int) ([]*model.Figure, bool, error) {
	ret := m.ctrl.Call(m, "Figures", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*model.Figure)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Figures indicates an expected call of Figures
func (mr *MockDaoIntMockRecorder) Figures(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Figures", reflect.TypeOf((*MockDaoInt)(nil).Figures), arg0, arg1, arg2)
}

// InsertRankHistory mocks base method
func (m *MockDaoInt) InsertRankHistory(arg0 context.Context, arg1 *model.Rank) (int64, error) {
	ret := m.ctrl.Call(m, "InsertRankHistory", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertRankHistory indicates an expected call of InsertRankHistory
func (mr *MockDaoIntMockRecorder) InsertRankHistory(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertRankHistory", reflect.TypeOf((*MockDaoInt)(nil).InsertRankHistory), arg0, arg1)
}

// PendingMidsCache mocks base method
func (m *MockDaoInt) PendingMidsCache(arg0 context.Context, arg1, arg2 int64) ([]int64, error) {
	ret := m.ctrl.Call(m, "PendingMidsCache", arg0, arg1, arg2)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingMidsCache indicates an expected call of PendingMidsCache
func (mr *MockDaoIntMockRecorder) PendingMidsCache(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingMidsCache", reflect.TypeOf((*MockDaoInt)(nil).PendingMidsCache), arg0, arg1, arg2)
}

// Ping mocks base method
func (m *MockDaoInt) Ping(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Ping", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
func (mr *MockDaoIntMockRecorder) Ping(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockDaoInt)(nil).Ping), arg0)
}

// PutCalcRecord mocks base method
func (m *MockDaoInt) PutCalcRecord(arg0 context.Context, arg1 *model.FigureRecord, arg2 int64) error {
	ret := m.ctrl.Call(m, "PutCalcRecord", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutCalcRecord indicates an expected call of PutCalcRecord
func (mr *MockDaoIntMockRecorder) PutCalcRecord(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCalcRecord", reflect.TypeOf((*MockDaoInt)(nil).PutCalcRecord), arg0, arg1, arg2)
}

// RemoveCache mocks base method
func (m *MockDaoInt) RemoveCache(arg0 context.Context, arg1 int64) error {
	ret := m.ctrl.Call(m, "RemoveCache", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCache indicates an expected call of RemoveCache
func (mr *MockDaoIntMockRecorder) RemoveCache(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCache", reflect.TypeOf((*MockDaoInt)(nil).RemoveCache), arg0, arg1)
}

// SetFigureCache mocks base method
func (m *MockDaoInt) SetFigureCache(arg0 context.Context, arg1 *model.Figure) error {
	ret := m.ctrl.Call(m, "SetFigureCache", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFigureCache indicates an expected call of SetFigureCache
func (mr *MockDaoIntMockRecorder) SetFigureCache(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFigureCache", reflect.TypeOf((*MockDaoInt)(nil).SetFigureCache), arg0, arg1)
}

// UpsertFigure mocks base method
func (m *MockDaoInt) UpsertFigure(arg0 context.Context, arg1 *model.Figure) (int64, error) {
	ret := m.ctrl.Call(m, "UpsertFigure", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertFigure indicates an expected call of UpsertFigure
func (mr *MockDaoIntMockRecorder) UpsertFigure(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertFigure", reflect.TypeOf((*MockDaoInt)(nil).UpsertFigure), arg0, arg1)
}

// UpsertRank mocks base method
func (m *MockDaoInt) UpsertRank(arg0 context.Context, arg1 *model.Rank) (int64, error) {
	ret := m.ctrl.Call(m, "UpsertRank", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertRank indicates an expected call of UpsertRank
func (mr *MockDaoIntMockRecorder) UpsertRank(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRank", reflect.TypeOf((*MockDaoInt)(nil).UpsertRank), arg0, arg1)
}

// UserInfo mocks base method
func (m *MockDaoInt) UserInfo(arg0 context.Context, arg1, arg2 int64) (*model.UserInfo, error) {
	ret := m.ctrl.Call(m, "UserInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserInfo indicates an expected call of UserInfo
func (mr *MockDaoIntMockRecorder) UserInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserInfo", reflect.TypeOf((*MockDaoInt)(nil).UserInfo), arg0, arg1, arg2)
}