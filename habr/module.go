package habr

import (
	"github.com/rwlist/rwcore/app"
	"github.com/rwlist/rwcore/hab"
)

type Module struct {
	app    *app.App
	reader *client.ReaderDailyTop
}

func (m *Module) Init(app *app.App) error {
	m.app = app
	m.reader = client.NewReaderDailyTop()
	app.Router.Mount("/habr", m.Router())
	go m.process()
	return nil
}
