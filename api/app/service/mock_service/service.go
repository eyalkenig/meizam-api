// Automatically generated by MockGen. DO NOT EDIT!
// Source: /Users/eyalkenig/Development/yotpo-workspace/yotpo-go-workspace/src/github.com/eyalkenig/meizam-api/api/app/service/service.go

package mock_service

import (
	models "github.com/eyalkenig/meizam-api/api/app/repository/mysql/models"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *_MockServiceRecorder
}

// Recorder for MockService (not exported)
type _MockServiceRecorder struct {
	mock *MockService
}

func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &_MockServiceRecorder{mock}
	return mock
}

func (_m *MockService) EXPECT() *_MockServiceRecorder {
	return _m.recorder
}

func (_m *MockService) CreateTeam(teamName string, externalEntityId *string, imageUrl *string) (*models.Team, error) {
	ret := _m.ctrl.Call(_m, "CreateTeam", teamName, externalEntityId, imageUrl)
	ret0, _ := ret[0].(*models.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockServiceRecorder) CreateTeam(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateTeam", arg0, arg1, arg2)
}

func (_m *MockService) ListTeams(limit int, offset int) ([]*models.Team, error) {
	ret := _m.ctrl.Call(_m, "ListTeams", limit, offset)
	ret0, _ := ret[0].([]*models.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockServiceRecorder) ListTeams(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTeams", arg0, arg1)
}

// Mock of Command interface
type MockCommand struct {
	ctrl     *gomock.Controller
	recorder *_MockCommandRecorder
}

// Recorder for MockCommand (not exported)
type _MockCommandRecorder struct {
	mock *MockCommand
}

func NewMockCommand(ctrl *gomock.Controller) *MockCommand {
	mock := &MockCommand{ctrl: ctrl}
	mock.recorder = &_MockCommandRecorder{mock}
	return mock
}

func (_m *MockCommand) EXPECT() *_MockCommandRecorder {
	return _m.recorder
}

func (_m *MockCommand) CreateTeam(teamName string, externalEntityId *string, imageUrl *string) (*models.Team, error) {
	ret := _m.ctrl.Call(_m, "CreateTeam", teamName, externalEntityId, imageUrl)
	ret0, _ := ret[0].(*models.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCommandRecorder) CreateTeam(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateTeam", arg0, arg1, arg2)
}

// Mock of Query interface
type MockQuery struct {
	ctrl     *gomock.Controller
	recorder *_MockQueryRecorder
}

// Recorder for MockQuery (not exported)
type _MockQueryRecorder struct {
	mock *MockQuery
}

func NewMockQuery(ctrl *gomock.Controller) *MockQuery {
	mock := &MockQuery{ctrl: ctrl}
	mock.recorder = &_MockQueryRecorder{mock}
	return mock
}

func (_m *MockQuery) EXPECT() *_MockQueryRecorder {
	return _m.recorder
}

func (_m *MockQuery) ListTeams(limit int, offset int) ([]*models.Team, error) {
	ret := _m.ctrl.Call(_m, "ListTeams", limit, offset)
	ret0, _ := ret[0].([]*models.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockQueryRecorder) ListTeams(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTeams", arg0, arg1)
}