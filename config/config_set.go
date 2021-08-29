package config

import "github.com/google/wire"

var (
	ProdSet = wire.NewSet(ProvideProdConfig)
	TestSet = wire.NewSet(ProvideTestConfig)
)
