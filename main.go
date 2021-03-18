// Copyright (C) <year>  <name of author>

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
	"log"
	"os"
	"os/exec"

	"github.com/awesome-gocui/gocui"
)

var cmd *exec.Cmd

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <command> [args...]", os.Args[0])
		os.Exit(1)
	}
	cmd = exec.Command(os.Args[1], os.Args[2:]...)
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Mouse = true

	g.SetManagerFunc(layout)

	if err := keybinds(g); err != nil {
		log.Panicln(err)
	}
	if err := g.MainLoop(); err != nil && !gocui.IsQuit(err) {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	init := false
	maxX, maxY := g.Size()
	if vout, err := g.SetView("output", 0, 0, maxX-1, maxY-3, gocui.BOTTOM); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		vout.Autoscroll = true
		vout.Wrap = true
		init = true
	}
	if init {
		if err := initcmd(g); err != nil {
			return err
		}
	}
	return nil
}

func keybinds(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.MouseWheelUp, gocui.ModNone, scroll(-1)); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.MouseWheelDown, gocui.ModNone, scroll(1)); err != nil {
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
		v.SetOrigin(0, y+n)
		return nil
	}
}

func updateView(buffer [512]byte, n int) func(g *gocui.Gui) error {
	return func(g *gocui.Gui) error {
		vout, err := g.View("output")
		if err != nil {
			return err
		}
		_, err = vout.Write(buffer[:n])
		return err
	}
}

func initcmd(g *gocui.Gui) error {
	reader, writer := io.Pipe()
	cmd.Stdout = writer
	cmd.Stderr = writer
	cmd.Stdin = nil
	go func() {
		var buffer [512]byte
		var n int
		var err error
		for err == nil {
			n, err = reader.Read(buffer[:])
			g.Update(updateView(buffer, n))
		}
	}()
	err := cmd.Start()
	if err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
