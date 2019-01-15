package ctx

import (
	"context"
	"github.com/globalsign/mgo"
)

func DB(ctx context.Context) *mgo.Database {
	return ctx.Value("db").(*mgo.Database)
}