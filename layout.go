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

func layout(g *gocui.Gui) error {
	init := false
	maxX, maxY := g.Size()
	if v, err := g.SetView("output", 0, 0, maxX-1, maxY-3, gocui.BOTTOM); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Autoscroll = true
		v.Wrap = true
		init = true
	}
	if v, err := g.SetView("input", 0, maxY-3, maxX-1, maxY-1, gocui.TOP); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Editable = true
	}
	g.SetCurrentView("input")
	if init {
		if err := initcmd(g); err != nil {
			return err
		}
	}
	return nil
}
