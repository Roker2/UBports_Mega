package uploader

import (
	"../user"
	"github.com/nanu-c/qml-go"
	"log"
	"os"
)

type Uploader struct {
	Percent float32

	SetterPercent float32

	setterPercentChanged int
	setterPercentSet float32

	getterPercent float32
	getterPercentChanged int
}

func (u *Uploader) SetSetterPercent(p float32) {
	u.setterPercentChanged++
	u.setterPercentSet = p
}

func (u *Uploader) GetterPercent() float32 {
	return u.getterPercent
}

func (u *Uploader) SetGetterPercent(p float32) {
	u.getterPercentChanged++
	u.getterPercent = p
}

func (u *Uploader) ChangePercent(new float32) (old float32) {
	old = u.Percent
	u.Percent = new
	return
}

func (u *Uploader) NotifyPercentChanged() {
	qml.Changed(u, &u.Percent)
}

func (u *Uploader) UploadFile(user *user.User, srcpath string) {
	var ch *chan int
	ch = new(chan int)
	*ch = make(chan int)
	fi, err := os.Stat(srcpath)
	if err != nil {
		log.Println(err)
		return
	}
	go func() {
		bytesread := 0
		for {
			b := 0
			ok := false
			select {
			case b, ok = <-*ch:
				if ok == false {
					if u.Percent != 100 {
						u.Percent = 100
						log.Println(u.Percent)
						u.NotifyPercentChanged()
					}
					return
				}
			}
			bytesread += b
			u.Percent = float32(bytesread) * 100 / float32(fi.Size())
			log.Println(u.Percent)
			u.NotifyPercentChanged()
		}
	}()
	go func() {
		//err := u.Mega.DownloadFile(u.GetCurrentNode(), "/tmp/" + u.GetCurrentNodeName(), ch)
		_, err := user.Mega.UploadFile("", user.GetCurrentNode(), "", ch)
		if err != nil {
			log.Println(err)
		}
	}()
}

func Register() {
	qml.RegisterTypes("GoMegaUploader", 1, 0, []qml.TypeSpec{{
		Init: func(v *Uploader, obj qml.Object) {
			log.Println("Init Uploader")
		},
	}})
}
