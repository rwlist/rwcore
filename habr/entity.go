package habr

import (
	"github.com/google/wire"
	"github.com/rwlist/rwcore/habr/client"
)

var All = wire.NewSet(
	NewService,
	client.NewClient,
	client.NewReaderDailyTop,
)
