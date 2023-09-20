package captcha

import (
	"github.com/google/wire"
	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	*base64Captcha.Captcha
}

type Config struct {
	// local store default memory store
	// only supports single machine deployment
	// please customize Redis store for multiple servers
	StoreType StoreType
	DriveType DriveType
	//Height Captcha png height in pixel.
	Height int
	//Width Captcha png width in pixel.
	Width int
}

func NewConfig() (*Config, error) {
	return &Config{StoreType: LocalStore, DriveType: Digit, Height: 80, Width: 240}, nil
}

// Verify .
func (ca *Captcha) Verify(id string, code string) bool {
	return ca.Store.Verify(id, code, false)
}

func NewCaptcha(driver *base64Captcha.DriverDigit, store *base64Captcha.Store) (*Captcha, error) {
	ca := &Captcha{base64Captcha.NewCaptcha(driver, *store)}
	return ca, nil
}

var CaptchaSet = wire.NewSet(NewCaptcha, NewConfig)
