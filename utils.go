package utils

import "github.com/google/wire"

// ProviderSet .
var ProviderSet = wire.NewSet(NewCaptcha)
