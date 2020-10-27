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
	"./stack"
	"github.com/nanu-c/qml-go"
	"github.com/t3rm1n4l/go-mega"
	"log"
	"strings"
)

type user struct {
	Login string
	Password string
	mega *mega.Mega
	nodeStack stack.Stack
	dicHashNode map[string]*mega.Node
	Percent float32
}

var Root qml.Object
var Engine *qml.Engine

func main() {
	err := qml.Run(run)
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	Engine = qml.NewEngine()
	component, err := Engine.LoadFile("qml/MEGA.qml")
	if err != nil {
		return err
	}

	u := user{Login: "Login", Password: "Password", mega: mega.New()}
	context := Engine.Context()
	context.SetVar("u", &u)

	win := component.CreateWindow(nil)
	Root = win.Root()
	win.Show()
	win.Wait()

	return nil
}

func (u *user) SignIn() bool {
	//log.Println("Login: " + u.Login)
	//log.Println("Password: " + u.Password)
	err := u.mega.Login(u.Login, u.Password)
	if err != nil {
        log.Println(err)
        return false
	} else {
        //log.Println("Work")
        u.nodeStack.Push(u.mega.FS.GetRoot())
        return true
	}
}

func (u *user) GetFiles() string {
    nodes, err := u.mega.FS.GetChildren(u.nodeStack.Peek())
    if err != nil {
		log.Println(err)
		return ""
	}
	var paths string
	for _, node := range nodes {
		paths += node.GetName() + "|"
		//log.Println(node.GetName())
	}
	paths = strings.TrimSuffix(paths, "|")
	return paths
}

func (u *user) GetHashes() string {
	nodes, err := u.mega.FS.GetChildren(u.nodeStack.Peek())
	if err != nil {
		log.Println(err)
		return ""
	}
	var hashes string
	dic := make(map[string]*mega.Node)
	for _, node := range nodes {
		dic[node.GetHash()] = node
		hashes += node.GetHash() + "|"
	}
	hashes = strings.TrimSuffix(hashes, "|")
	u.dicHashNode = dic
	return hashes
}

func (u *user) RegenerateDictionary() {
	nodes, err := u.mega.FS.GetChildren(u.nodeStack.Peek())
	if err != nil {
		log.Println(err)
		return
	}
	dic := make(map[string]*mega.Node)
	for _, node := range nodes {
		dic[node.GetHash()] = node
	}
	u.dicHashNode = dic
}

func (u *user) GetCurrentNodeName() string {
	return u.nodeStack.Peek().GetName()
}

func (u *user) GetCurrentNodeHash() string {
	return u.nodeStack.Peek().GetHash()
}

func (u *user) PushNode(hash string) {
	//log.Println(hash)
	u.nodeStack.Push(u.dicHashNode[hash])
}

func (u *user) PopNode() {
	u.nodeStack.Pop()
}

func (u *user) GetNumberOfChildren() int {
	nodes, err := u.mega.FS.GetChildren(u.nodeStack.Peek())
	if err != nil {
		log.Println(err)
		return -1
	}
	return len(nodes)
}

func (u *user) DownloadCurrentNode() {
	var ch *chan int
	ch = new(chan int)
	*ch = make(chan int)
	go u.showProgress(*ch, u.nodeStack.Peek().GetSize())
	err := u.mega.DownloadFile(u.nodeStack.Peek(), "/tmp/" + u.nodeStack.Peek().GetName(), ch)
	if err != nil {
		log.Println(err)
	}
}

func (u *user) showProgress(ch chan int, size int64) {
	bytesread := 0
	u.Percent = float32(0)
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
		u.Percent = 100 * float32(bytesread) / float32(size)
		log.Println(u.Percent)
	}
}
