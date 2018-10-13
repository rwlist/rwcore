package app

type Module interface {
	Init(app *App) error
}

func (app *App) AddModule(module Module) error {
	err := module.Init(app)
	if err != nil {
		return err
	}
	app.modules = append(app.modules, module)
	return nil
}
