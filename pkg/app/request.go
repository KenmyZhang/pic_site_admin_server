package app

import (
	"errors"
	"github.com/astaxie/beego/validation"
)

// MarkErrors logs error logs
func MarkErrors(errs []*validation.Error) error {
	errMsg := ""
	for _, err := range errs {
		errMsg = errMsg + err.Key + err.Message
	}

	return errors.New(errMsg)
}
