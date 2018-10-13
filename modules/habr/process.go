package habr

import "github.com/davecgh/go-spew/spew"

func (m *Module) process() {
	db := m.app.DB.Copy()
	defer db.Close()

	for v := range m.reader.Read() {
		spew.Dump(v)
	}
}
