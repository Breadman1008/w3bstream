// This is a generated source file. DO NOT EDIT
// Source: status/error__generated.go

package status

import (
	"github.com/iotexproject/Bumblebee/kit/statusx"
)

var _ statusx.Error = (*Error)(nil)

func (v Error) StatusErr() *statusx.StatusErr {
	return &statusx.StatusErr{
		Key:       v.Key(),
		Code:      v.Code(),
		Msg:       v.Msg(),
		CanBeTalk: v.CanBeTalk(),
	}
}

func (v Error) Unwrap() error {
	return v.StatusErr()
}

func (v Error) Error() string {
	return v.StatusErr().Error()
}

func (v Error) StatusCode() int {
	return statusx.StatusCodeFromCode(int(v))
}

func (v Error) Code() int {
	if with, ok := (interface{})(v).(statusx.ServiceCode); ok {
		return with.ServiceCode() + int(v)
	}
	return int(v)

}

func (v Error) Key() string {
	switch v {
	case BadRequest:
		return "BadRequest"
	case Unauthorized:
		return "Unauthorized"
	case NotFound:
		return "NotFound"
	case Conflict:
		return "Conflict"
	case InternalServerError:
		return "InternalServerError"
	}
	return "UNKNOWN"
}

func (v Error) Msg() string {
	switch v {
	case BadRequest:
		return "BadRequest"
	case Unauthorized:
		return "Unauthorized"
	case NotFound:
		return "NotFound"
	case Conflict:
		return "Conflict conflict error"
	case InternalServerError:
		return "InternalServerError 内部错误"
	}
	return "-"
}

func (v Error) CanBeTalk() bool {
	switch v {
	case BadRequest:
		return false
	case Unauthorized:
		return true
	case NotFound:
		return false
	case Conflict:
		return false
	case InternalServerError:
		return false
	}
	return false
}
