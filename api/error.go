package api

import (
	"Activity/constant"
	"fmt"
)

// Error 自定义错误类型
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error 实现error接口
func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

// NewError 创建错误
func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// 预定义错误
var (
	ErrSystem                  = NewError(constant.ErrSystem, constant.ErrMsgSystem)
	ErrInvalidParam            = NewError(constant.ErrInvalidParam, constant.ErrMsgInvalidParam)
	ErrActivityNotFound        = NewError(constant.ErrActivityNotFound, constant.ErrMsgActivityNotFound)
	ErrActivityEnded           = NewError(constant.ErrActivityEnded, constant.ErrMsgActivityEnded)
	ErrActivityNotStarted      = NewError(constant.ErrActivityNotStarted, constant.ErrMsgActivityNotStarted)
	ErrGameNotFound            = NewError(constant.ErrGameNotFound, constant.ErrMsgGameNotFound)
	ErrGameClosed              = NewError(constant.ErrGameClosed, constant.ErrMsgGameClosed)
	ErrUserAlreadyParticipated = NewError(constant.ErrUserAlreadyParticipated, constant.ErrMsgUserAlreadyParticipated)
	ErrPrizeStockEmpty         = NewError(constant.ErrPrizeStockEmpty, constant.ErrMsgPrizeStockEmpty)
	ErrUserNotPosted           = NewError(constant.ErrUserNotPosted, constant.ErrMsgUserNotPosted)
	ErrUserNotCheckedIn        = NewError(constant.ErrUserNotCheckedIn, constant.ErrMsgUserNotCheckedIn)
)
