package resp

import (
	"errors"

	"github.com/crochee/lib/e"
	"github.com/crochee/lib/log"
	"github.com/gin-gonic/gin"
)

type ResponseCode struct {
	// 返回码
	// Required: true
	// Example: 20000000
	Code int `json:"code"`
	// 返回信息描述
	// Required: true
	// Example: success
	Msg string `json:"message"`
	// 具体描述信息
	Result string `json:"result"`
}

// Error gin response with Code
func Error(ctx *gin.Context, code e.Code) {
	ctx.JSON(code.StatusCode(), ResponseCode{
		Code: code.Code(),
		Msg:  code.Error(),
	})
}

// A Wrapper provides context around another error.
type Wrapper interface {
	// Unwrap returns the next error in the error chain.
	// If there is no next error, Unwrap returns nil.
	Unwrap() error
}

// Errors gin Response with error
func Errors(ctx *gin.Context, err error) {
	log.FromContext(ctx.Request.Context()).Errorf("%+v", err)
	for err != nil {
		wrapper, ok := err.(Wrapper)
		if !ok {
			break
		}
		err = wrapper.Unwrap()
	}
	if err == nil {
		Error(ctx, e.ErrInternalServerError)
		return
	}
	var errorCode e.Code
	if errors.As(err, &errorCode) {
		Error(ctx, errorCode)
		return
	}
	Error(ctx, e.ErrInternalServerError)
}

func ErrorParam(ctx *gin.Context, err error) {
	log.FromContext(ctx.Request.Context()).Errorf("parse param failed.%+v", err)
	ctx.JSON(e.ErrInvalidParam.StatusCode(), ResponseCode{
		Code:   e.ErrInvalidParam.Code(),
		Msg:    e.ErrInvalidParam.Error(),
		Result: err.Error(),
	})
}