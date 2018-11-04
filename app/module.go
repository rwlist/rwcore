package app

import "log"

type Module interface {
	Init(app *App) error
}

func (app *App) AddModule(module Module) error {
	err := module.Init(app)
	if err != nil {
		log.Println(err)
		return err
	}
	app.modules = append(app.modules, module)
	return nil
}
