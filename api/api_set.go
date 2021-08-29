package api

import "github.com/google/wire"

var ProdSet = wire.NewSet(ProvideRouter)
