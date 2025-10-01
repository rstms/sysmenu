//go:build windows

package menu

import (
	_ "embed"
	"fmt"
	"github.com/getlantern/systray"
	"log"
	"runtime"
)

func (m *Menu) startup() error {
	if m.debug {
		log.Println("Menu.startup")
	}
	if m.started {
		return Fatalf("already started")
	}
	m.wg.Add(1)
	go func() {
		if m.debug {
			log.Println("Menu.EventLoop started")
			defer log.Println("Menu.EventLoop exiting")
		}
		defer m.wg.Done()
		runtime.LockOSThread()
		systray.Run(m.onReady, m.onExit)
	}()
	m.started = true
	return nil
}

func (m *Menu) shutdown() error {
	if m.debug {

		log.Println("Menu.shutdown")
	}
	if !m.started {
		return Fatalf("never started")
	}
	if m.stopped {
		return Fatalf("already stoped")
	}
	for _, item := range m.items {
		item.stop()
	}
	m.stopped = true
	m.exitHandler <- struct{}{}
	return nil
}

func (m *Menu) onReady() {
	if m.debug {
		log.Println("Menu.onReady")
	}
	// Set the icon and tooltip
	systray.SetTitle(m.Title)
	systray.SetTooltip(m.Title)
	systray.SetIcon(m.iconData)

	for _, item := range m.items {
		switch item.Type {
		case MenuItemClickable, MenuItemQuit:
			item.start(systray.AddMenuItem(item.Title, item.Tooltip))
		case MenuItemCheckbox:
			item.start(systray.AddMenuItemCheckbox(item.Title, item.Tooltip, item.checked))
		case MenuItemSeparator:
			systray.AddSeparator()
		default:
			panic(fmt.Sprintf("unexpected item: %v", item))
		}
	}
}

func (m *Menu) onExit() {
	if m.debug {
		log.Println("Menu.onExit")
	}
	m.Stop()
}

func (i *MenuItem) start(trayItem *systray.MenuItem) {
	if i.menu.debug {
		log.Printf("MenuItem.start %d %s\n", i.Id, i.Title)
	}
	i.trayItem = trayItem
	i.menu.wg.Add(1)
	go i.handler()
	for _, subItem := range i.subItems {
		switch subItem.Type {
		case MenuItemClickable, MenuItemQuit:
			subItem.start(i.trayItem.AddSubMenuItem(subItem.Title, subItem.Tooltip))
		case MenuItemCheckbox:
			subItem.start(i.trayItem.AddSubMenuItemCheckbox(subItem.Title, subItem.Tooltip, subItem.checked))
		case MenuItemSeparator:
			systray.AddSeparator()
		default:
			panic(fmt.Sprintf("unexpected subItem: %v", subItem))
		}
	}
}

func (i *MenuItem) handler() {
	defer i.menu.wg.Done()
	if i.menu.debug {
		defer log.Printf("MenuItem handler exit %d %s\n", i.Id, i.Title)
		log.Printf("MenuItem handler start %d %s\n", i.Id, i.Title)
	}
	for {
		select {
		case <-i.trayItem.ClickedCh:
			if i.menu.debug {
				log.Printf("received ClickedCh:  %d %s\n", i.Id, i.Title)
			}
			if i.Id == i.menu.qid {
				if i.menu.debug {
					log.Println("Quit Item clicked; calling systray.Quit()")
				}
				systray.Quit()
			}
			i.menu.clickMux <- i
		case <-i.exitHandler:
			return
		}
	}
}

func (i *MenuItem) stop() {
	if i.menu.debug {
		log.Printf("MenuItem.stop %d %s\n", i.Id, i.Title)
	}
	i.exitHandler <- struct{}{}
	for _, subItem := range i.subItems {
		subItem.stop()
	}
}
