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
	"io"
	"os/exec"

	"github.com/awesome-gocui/gocui"
)

var cmd *exec.Cmd

func initcmd(g *gocui.Gui) error {
	reader, writer := io.Pipe()
	cmd.Stdout = writer
	cmd.Stderr = writer
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	go func() {
		var buffer [512]byte
		var n int
		var err error
		for err == nil {
			n, err = reader.Read(buffer[:])
			g.Update(updateView(buffer, n))
		}
	}()
	g.Update(func(g *gocui.Gui) error {
		vin, err := g.View("input")
		if err != nil {
			return err
		}
		vin.Editor = gocui.EditorFunc(editorwriter(g, stdin))
		return nil
	})
	err = cmd.Start()
	if err != nil {
		return err
	}
	return nil
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
