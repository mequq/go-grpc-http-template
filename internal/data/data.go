package data

import "github.com/google/wire"

var DataProviderSet = wire.NewSet(
	NewDataSource,
	NewTenantRepo,
)
