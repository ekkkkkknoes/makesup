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
	"fmt"
	"io"

	"github.com/awesome-gocui/gocui"
)

func editorwriter(g *gocui.Gui, writer io.Writer) func(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	return func(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
		switch key {
		case gocui.KeyEnter:
			g.Update(func(g *gocui.Gui) error {
				_, err := writer.Write([]byte(v.ViewBuffer()))
				if err != nil {
					return nil // cmd has exited/closed stdin if this fails. makesup should keep running
				}
				writer.Write([]byte("\n"))
				vout, err := g.View("output")
				if err != nil {
					return err
				}
				fmt.Fprintf(vout, ">>> %s\n", v.ViewBuffer())
				v.Clear()
				v.SetCursor(0, 0)
				return nil
			})
		default:
			gocui.DefaultEditor.Edit(v, key, ch, mod)
		}
	}
}
