package articles

import (
	"github.com/rwlist/rwcore/app"
)

type Module struct {
	app *app.App
}

func (m *Module) Init(app *app.App) error {
	m.app = app
	app.Router.Mount("/articles", m.Router())
	return nil
}
