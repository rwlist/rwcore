package admin

import (
	"github.com/google/wire"
)

var All = wire.NewSet(
	NewController,
	NewRouter,
)
