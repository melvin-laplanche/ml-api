// Code generated by mockery v1.0.0
package mockmailer

import mailer "github.com/Nivl/go-rest-tools/notifiers/mailer"
import mock "github.com/stretchr/testify/mock"

// Mailer is an autogenerated mock type for the Mailer type
type Mailer struct {
	mock.Mock
}

// Send provides a mock function with given fields: msg
func (_m *Mailer) Send(msg *mailer.Message) error {
	ret := _m.Called(msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(*mailer.Message) error); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendStackTrace provides a mock function with given fields: trace, endpoint, message, id
func (_m *Mailer) SendStackTrace(trace []byte, endpoint string, message string, id string) error {
	ret := _m.Called(trace, endpoint, message, id)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, string, string, string) error); ok {
		r0 = rf(trace, endpoint, message, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}