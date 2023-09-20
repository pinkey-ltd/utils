package captcha

import (
	"github.com/google/wire"
	"github.com/mojocn/base64Captcha"
)

type StoreType int

const (
	LocalStore StoreType = 1 << iota
	RedisStore
	CustomStore
)

func NewStore(cnf *Config) (*base64Captcha.Store, error) {
	switch cnf.StoreType {
	case LocalStore:
		return &base64Captcha.DefaultMemStore, nil
	case RedisStore:
		return nil, nil
	default:
		return &base64Captcha.DefaultMemStore, nil
	}
}

var StoreSet = wire.NewSet(NewStore)
