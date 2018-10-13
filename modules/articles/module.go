package articles

import (
	"github.com/rwlist/rwcore/app"
)

type Articles struct {
	app  *app.App
	impl impl
}

func (a *Articles) Init(app *app.App) {
	a.app = app
	a.impl = impl{}
	app.Router.Mount("/articles", a.Router())
}
