// Code generated by MockGen. DO NOT EDIT.
// Source: secretsmanager.go

// Package mock is a generated GoMock package.
package mock

import (
	secretsmanager "github.com/aws/aws-sdk-go/service/secretsmanager"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// SecretsManagerClient is a mock of SecretsManagerClient interface
type SecretsManagerClient struct {
	ctrl     *gomock.Controller
	recorder *SecretsManagerClientMockRecorder
}

// SecretsManagerClientMockRecorder is the mock recorder for SecretsManagerClient
type SecretsManagerClientMockRecorder struct {
	mock *SecretsManagerClient
}

// NewSecretsManagerClient creates a new mock instance
func NewSecretsManagerClient(ctrl *gomock.Controller) *SecretsManagerClient {
	mock := &SecretsManagerClient{ctrl: ctrl}
	mock.recorder = &SecretsManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *SecretsManagerClient) EXPECT() *SecretsManagerClientMockRecorder {
	return m.recorder
}

// GetSecretValue mocks base method
func (m *SecretsManagerClient) GetSecretValue(arg0 *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretValue", arg0)
	ret0, _ := ret[0].(*secretsmanager.GetSecretValueOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretValue indicates an expected call of GetSecretValue
func (mr *SecretsManagerClientMockRecorder) GetSecretValue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretValue", reflect.TypeOf((*SecretsManagerClient)(nil).GetSecretValue), arg0)
}
