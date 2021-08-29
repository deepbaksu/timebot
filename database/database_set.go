package database

import "github.com/google/wire"

var ProdSet = wire.NewSet(ProvideMongoClient)
