// Copyright (C) 2021  Sönke Lambert

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"github.com/awesome-gocui/gocui"
)

func keybinds(g *gocui.Gui) error {
	if err := g.SetKeybinding("input", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("input", gocui.KeyPgup, gocui.ModNone, scroll(-10)); err != nil {
		return err
	}
	if err := g.SetKeybinding("input", gocui.KeyArrowUp, gocui.ModNone, scroll(-1)); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.MouseWheelUp, gocui.ModNone, scroll(-1)); err != nil {
		return err
	}
	if err := g.SetKeybinding("input", gocui.KeyPgdn, gocui.ModNone, scroll(10)); err != nil {
		return err
	}
	if err := g.SetKeybinding("input", gocui.KeyArrowDown, gocui.ModNone, scroll(1)); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.MouseWheelDown, gocui.ModNone, scroll(1)); err != nil {
		return err
	}
	if err := g.SetKeybinding("input", gocui.KeyCtrlSpace, gocui.ModNone, scroll(0)); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.MouseMiddle, gocui.ModNone, scroll(0)); err != nil {
		return err
	}
	return nil
}

func scroll(n int) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		v, err := g.View("output")
		if v == nil || err != nil {
			return err
		}
		if n == 0 {
			v.Autoscroll = true
			v.Title = ""
			return nil
		}
		v.Autoscroll = false
		v.Title = "scrolling"
		_, y := v.Origin()
		y += n
		if y < 0 {
			y = 0
		}
		v.SetOrigin(0, y)
		return nil
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
