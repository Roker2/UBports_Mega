package ui

import (
	"../user"
	"github.com/nanu-c/qml-go"
	"github.com/t3rm1n4l/go-mega"
	"log"
)

var Win *qml.Window
var Engine *qml.Engine

func SetEngine() {
	Engine = qml.NewEngine()
}

func InitModels() {
	u := user.User{Login: "Login", Password: "Password", Mega: mega.New()}
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
