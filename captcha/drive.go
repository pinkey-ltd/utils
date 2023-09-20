package captcha

import (
	"github.com/google/wire"
	"github.com/mojocn/base64Captcha"
)

type DriveType int

const (
	Digit DriveType = 1 << iota
	String
	Math
	Chinese
	Custom
)

func NewDrive(cnf *Config) (*base64Captcha.DriverDigit, error) {
	switch cnf.DriveType {
	case Digit:
		return base64Captcha.NewDriverDigit(cnf.Height, cnf.Width, 6, 0.6, 8), nil
	case String:
		return nil, nil
	default:
		return base64Captcha.NewDriverDigit(80, 240, 6, 0.6, 8), nil
	}
}

var DriveSet = wire.NewSet(NewDrive)
