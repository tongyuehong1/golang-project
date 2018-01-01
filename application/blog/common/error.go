package common

import "errors"

const (
	ErrSucceed      = 0   // 成功
	ErrMysqlQuery   = 500 // MySQL 错误
	ErrInvalidParam = 1   // 参数错误
	ErrMysqlNotFound = 501
)

var ErrNotFound = errors.New("ErrNotFound")