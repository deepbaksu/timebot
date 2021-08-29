package application

import (
	"github.com/google/wire"
)

var ProdSet = wire.NewSet(ProvideApplication)
