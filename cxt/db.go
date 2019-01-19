package cxt

import (
	"context"
	"github.com/globalsign/mgo"
)

func DB(ctx context.Context) *mgo.Database {
	return ctx.Value(DbKey).(*mgo.Database)
}