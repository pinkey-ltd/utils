package utils

import (
	"github.com/google/wire"
	"github.com/pinkey-ltd/utils/config"
)

// ProviderSet .
var ProviderSet = wire.NewSet(NewCorn, config.NewConfig)
