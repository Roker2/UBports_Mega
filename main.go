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
	"log"
	//"strconv"

	"github.com/nanu-c/qml-go"
	"github.com/t3rm1n4l/go-mega"
)

type user struct {
	Login string
	Password string
	mega *mega.Mega
	Output string
}

var Root qml.Object

func main() {
	err := qml.Run(run)
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	engine := qml.NewEngine()
	component, err := engine.LoadFile("qml/login.qml")
	if err != nil {
		return err
	}

	u := user{Login: "Login", Password: "Password", Output: "Text", mega: mega.New()}
	context := engine.Context()
	context.SetVar("u", &u)

	win := component.CreateWindow(nil)
	Root = win.Root()
	win.Show()
	win.Wait()

	return nil
}

func (u *user) SignIn() bool {
	log.Println(u.Login)
	log.Println(u.Password)
	err := u.mega.Login(u.Login, u.Password)
	if err != nil {
		u.Output = err.Error()
	    qml.Changed(u, &u.Output)
        return false
	} else {
		u.Output = "Work."
	    qml.Changed(u, &u.Output)
        return true
	}
}
