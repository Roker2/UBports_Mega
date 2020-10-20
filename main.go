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
	"github.com/nanu-c/qml-go"
	"github.com/t3rm1n4l/go-mega"
	"log"
	"strings"
)

type user struct {
	Login string
	Password string
	mega *mega.Mega
	currentNode *mega.Node
	dicNameNode map[string]*mega.Node
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
	component, err := engine.LoadFile("qml/MEGA.qml")
	if err != nil {
		return err
	}

	u := user{Login: "Login", Password: "Password", mega: mega.New()}
	context := engine.Context()
	context.SetVar("u", &u)

	win := component.CreateWindow(nil)
	Root = win.Root()
	win.Show()
	win.Wait()

	return nil
}

func (u *user) SignIn() bool {
	log.Println("Login: " + u.Login)
	log.Println("Password: " + u.Password)
	err := u.mega.Login(u.Login, u.Password)
	if err != nil {
        log.Println(err)
        return false
	} else {
        log.Println("Work")
        u.currentNode = u.mega.FS.GetRoot()
        return true
	}
}

func (u *user) GetFiles() string {
    nodes, err := u.mega.FS.GetChildren(u.currentNode)
    if err != nil {
		log.Println(err)
		return ""
	}
	var paths string
	dic := make(map[string]*mega.Node)
	for _, node := range nodes {
		dic[node.GetName()] = node
		paths += node.GetName() + "|"
		log.Println(node.GetName())
	}
	paths = strings.TrimSuffix(paths, "|")
	u.dicNameNode = dic
	log.Println(u.dicNameNode)
	return paths
}

func (u *user) GetCurrentNodeName() string {
	return u.currentNode.GetName()
}

func (u *user) GetCurrentNodeHash() string {
	return u.currentNode.GetHash()
}
