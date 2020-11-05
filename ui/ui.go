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

package ui

import (
	"../user"
	"github.com/nanu-c/qml-go"
	"log"
)

var Win *qml.Window
var Engine *qml.Engine

func SetEngine() {
	Engine = qml.NewEngine()
}

func InitModels() {
	u := user.User{Login: "Login", Password: "Password"}
	Engine.Context().SetVar("u", &u)
}

func SetComponent() error {
	component, err := Engine.LoadFile("qml/MEGA.qml")
	if err != nil {
		log.Println(err)
		return err
	}
	Win = component.CreateWindow(nil)
	return nil
}
