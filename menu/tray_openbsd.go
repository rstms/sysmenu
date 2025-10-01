//go:build openbsd

package menu

type SystrayMenuItem struct {
	Title   string
	Tooltip string
}

func (m *Menu) startup() error {
	return nil
}

func (m *Menu) shutdown() error {
	return nil
}

/*
func (m *MenuItem) start() {
}

func (m *MenuItem) stop() {
}
*/
