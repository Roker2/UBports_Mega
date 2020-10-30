package user

import (
	"../stack"
	"github.com/t3rm1n4l/go-mega"
	"log"
	"strings"
)

type User struct {
	Login       string
	Password    string
	Mega        *mega.Mega
	nodeStack   stack.Stack
	dicHashNode map[string]*mega.Node
}

func (u *User) SignIn() bool {
	//log.Println("Login: " + u.Login)
	//log.Println("Password: " + u.Password)
	err := u.Mega.Login(u.Login, u.Password)
	if err != nil {
		log.Println(err)
		return false
	} else {
		//log.Println("Work")PushRoot()
		u.PushRoot()
		return true
	}
}

func (u *User) GetFiles() string {
	nodes, err := u.Mega.FS.GetChildren(u.nodeStack.Peek())
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

func (u *User) GetHashes() string {
	nodes, err := u.Mega.FS.GetChildren(u.nodeStack.Peek())
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

func (u *User) RegenerateDictionary() {
	nodes, err := u.Mega.FS.GetChildren(u.nodeStack.Peek())
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

func (u *User) GetCurrentNode() *mega.Node {
	return u.nodeStack.Peek()
}

func (u *User) GetCurrentNodeName() string {
	return u.nodeStack.Peek().GetName()
}

func (u *User) GetCurrentNodeHash() string {
	return u.nodeStack.Peek().GetHash()
}

func (u *User) GetCurrentNodeSize() int64 {
	return u.nodeStack.Peek().GetSize()
}

func (u *User) CurrentNodeIsFolder() bool {
	return u.nodeStack.Peek().GetType() == mega.FOLDER
}

func (u *User) DeleteCurrentNode() {
	err := u.Mega.Delete(u.nodeStack.Peek(), true)
	if err != nil {
		log.Println(err)
	}
}

func (u *User) RenameCurrentNode(newName string) {
	err := u.Mega.Rename(u.nodeStack.Peek(), newName)
	if err != nil {
		log.Println(err)
	}
}

func (u *User) PushNode(hash string) {
	//log.Println(hash)
	u.nodeStack.Push(u.dicHashNode[hash])
}

func (u *User) PopNode() {
	u.nodeStack.Pop()
}

func (u *User) GetNumberOfChildren() int {
	nodes, err := u.Mega.FS.GetChildren(u.nodeStack.Peek())
	if err != nil {
		log.Println(err)
		return -1
	}
	return len(nodes)
}

func (u *User) DownloadCurrentNode() {
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
			log.Println(float32(bytesread) * 100 / float32(u.nodeStack.Peek().GetSize()))
		}
	}()
	err := u.Mega.DownloadFile(u.nodeStack.Peek(), "/tmp/" + u.nodeStack.Peek().GetName(), ch)
	if err != nil {
		log.Println(err)
	}
}

func (u *User) StackIsEmpty() bool {
	return u.nodeStack.Len() == 0
}

func (u *User) PushRoot() {
	u.nodeStack.Push(u.Mega.FS.GetRoot())
}
