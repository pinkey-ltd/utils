package captcha

import (
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

func NewDrive(cnf *Config) (base64Captcha.Driver, error) {
	switch cnf.DriveType {
	case Digit:
		return base64Captcha.NewDriverDigit(cnf.Height, cnf.Width, cnf.Length, 0.6, 8), nil
	case String:
		return base64Captcha.NewDriverString(cnf.Height, cnf.Width, 20, 100, cnf.Length, "", nil, nil, nil), nil
	case Chinese:
		return base64Captcha.NewDriverChinese(cnf.Height, cnf.Width, 20, 1, cnf.Length, "", nil, nil, nil), nil
	case Math:
		return base64Captcha.NewDriverMath(cnf.Height, cnf.Width, 4, 20, nil, nil, nil), nil
	case Custom:
		driver := &base64Captcha.DriverString{
			Length:          4,
			Height:          cnf.Height,
			Width:           cnf.Width,
			ShowLineOptions: 2,
			NoiseCount:      0,
			Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		}
		return driver, nil
	default:
		return base64Captcha.NewDriverDigit(80, 240, 6, 0.6, 8), nil
	}
}
