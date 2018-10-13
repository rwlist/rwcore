package lists

import (
	"github.com/rwlist/rwcore/app"
)

type Lists struct {
	app *app.App
}

func (l *Lists) Init(app *app.App) {
	l.app = app

	app.Router.Mount("/lists", l.Router())
}
