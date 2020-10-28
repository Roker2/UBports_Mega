package downloader

import (
	"../user"
	"github.com/nanu-c/qml-go"
	"github.com/t3rm1n4l/go-mega"
	"log"
)

type Downloader struct {
	node *mega.Node
	Percent float32

	SetterPercent float32

	setterPercentChanged int
	setterPercentSet float32

	getterPercent float32
	getterPercentChanged int
}

func (d *Downloader) Test() {
	log.Println("TEST")
}

func (d *Downloader) SetSetterPercent(p float32) {
	d.setterPercentChanged++
	d.setterPercentSet = p
}

func (d *Downloader) GetterPercent() float32 {
	return d.getterPercent
}

func (d *Downloader) SetGetterPercent(p float32) {
	d.getterPercentChanged++
	d.getterPercent = p
}

func (d *Downloader) ChangePercent(new float32) (old float32) {
	old = d.Percent
	d.Percent = new
	return
}

func (d *Downloader) NotifyPercentChanged() {
	qml.Changed(d, &d.Percent)
}

func (d *Downloader) DownloadNode(u *user.User) {
	var ch *chan int
	ch = new(chan int)
	*ch = make(chan int)
	go func() {
		bytesread := 0
		for {
			b := 0
			ok := false
			select {
			case b, ok = <-*ch:
				if ok == false {
					return
				}
			}
			bytesread += b
			d.Percent = float32(bytesread) * 100 / float32(d.node.GetSize())
			log.Println(d.Percent)
			d.NotifyPercentChanged()
		}
	}()
	err := u.Mega.DownloadFile(d.node, "/tmp/" + d.node.GetName(), ch)
	if err != nil {
		log.Println(err)
	}
}

func Register() {
	qml.RegisterTypes("GoMegaDownloader", 1, 0, []qml.TypeSpec{{
		Init: func(v *Downloader, obj qml.Object) {
			log.Println("Init Downloader")
		},
	}})
}
