//go:build linux

package menu

import (
	_ "embed"
	"fmt"
	"github.com/getlantern/systray"
	"log"
	"runtime"
	"sync"
)

//go:embed icon.ico
var iconData []byte

func (m *Menu) run() {
	// Ensure the program is run with a Windows GUI context
	runtime.LockOSThread()
	systray.Run(m.onReady, m.onExit)
}

func (m *Menu) shutdown() {
	if m.started {
		if !m.stopped {
			systray.Quit()
			m.stopped = true
		}
	}
}

func (m *Menu) onReady() {
	log.Println("Menu.onReady")
	// Set the icon and tooltip
	systray.SetTitle(m.Title)
	systray.SetTooltip(m.Title)
	systray.SetIcon(iconData)

	for _, item := range m.items {
		switch item.Type {
		case MenuItemClickable, MenuItemQuit:
			item.start(systray.AddMenuItem(item.Title, item.Tooltip))
		case MenuItemCheckbox:
			item.start(systray.AddMenuItemCheckbox(item.Title, item.Tooltip, item.Checked))
		case MenuItemSeparator:
			systray.AddSeparator()
		}
	}
}

func (m *Menu) onExit() {
	log.Println("onExit: calling Stop")
	m.Stop()
}
