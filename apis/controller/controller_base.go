package controller

import "errors"

// ControllerBase definition
type ControllerBase struct{}

func NewControllerBase() *ControllerBase {
	return &ControllerBase{}
}

func (c ControllerBase) CommonFuncOnError(err error) error {
	return errors.New("ABC")
}
