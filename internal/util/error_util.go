package util

import (
	"github.com/go-kratos/kratos/v2/errors"
)

// HttpErr422 参数校验错误
func HttpErr422(errMsg string) *errors.Error {
	err := errors.New(422, "VALID_PARAM_ERROR", errMsg)
	return err
}

// HttpErr500 当前服务请求错误
func HttpErr500(errMsg string, ErrorData map[string]string) *errors.Error {
	err := errors.New(500, "BIZ_ERROR", errMsg)
	if ErrorData != nil {
		err = err.WithMetadata(ErrorData)
	}
	return err
}

// HttpErr404 查找报错
func HttpErr404(errMsg string) *errors.Error {
	err := errors.New(404, "NOT_FOUND", errMsg)
	//err = err.WithMetadata(map[string]string{
	//	"foo": "bar",
	//})
	return err
}

// HttpErr200 200报错？
/*func HttpErr200(errMsg string) *errors.Error {
	err := errors.New(200, "BIZ_ERROR", errMsg)
	//err = err.WithMetadata(map[string]string{
	//	"foo": "bar",
	//})
	return err
}*/

// HttpErr403 权限报错
func HttpErr403(errMsg string) *errors.Error {
	err := errors.New(403, "AUTHORIZATION_ERROR", errMsg)
	//err = err.WithMetadata(map[string]string{
	//	"foo": "bar",
	//})
	return err
}
