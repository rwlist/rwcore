package habr

import (
	"github.com/rwlist/rwcore/app"
)

type Module struct {
	app *app.App

	reader *ReaderDailyTop
}

func (m *Module) Init(app *app.App) error {
	m.app = app
	m.reader = NewReaderDailyTop()
	app.Router.Mount("/habr", m.Router())
	go m.process()
	return nil
}
