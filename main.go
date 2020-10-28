/*
 * Copyright (C) 2020  Dmitry Minko
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; version 3.
 *
 * MEGA is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"./downloader"
	"./user"
	"github.com/nanu-c/qml-go"
	"github.com/t3rm1n4l/go-mega"
	"log"
)

var Root qml.Object
var Engine *qml.Engine

func main() {
	err := qml.Run(run)
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	downloader.Register()
	Engine = qml.NewEngine()
	component, err := Engine.LoadFile("qml/MEGA.qml")
	if err != nil {
		return err
	}

	u := user.User{Login: "Login", Password: "Password", Mega: mega.New()}
	context := Engine.Context()
	context.SetVar("u", &u)

	win := component.CreateWindow(nil)
	Root = win.Root()
	win.Show()
	win.Wait()

	return nil
}

func showProgress(ch chan int) {
	bytesread := 0
	for {
		b := 0
		ok := false
		select {
		case b, ok = <-ch:
			if ok == false {
				return
			}
		}
		bytesread += b
	}
}
