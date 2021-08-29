package security

import "github.com/google/wire"

var ProdSet = wire.NewSet(ProvideSecurityService)
var TestSet = wire.NewSet(ProvideFakeSecurity)
