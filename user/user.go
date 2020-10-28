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
	NodeStack   stack.Stack
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
		//log.Println("Work")
		u.NodeStack.Push(u.Mega.FS.GetRoot())
		return true
	}
}

func (u *User) GetFiles() string {
	nodes, err := u.Mega.FS.GetChildren(u.NodeStack.Peek())
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
	nodes, err := u.Mega.FS.GetChildren(u.NodeStack.Peek())
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
	nodes, err := u.Mega.FS.GetChildren(u.NodeStack.Peek())
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

func (u *User) GetCurrentNodeName() string {
	return u.NodeStack.Peek().GetName()
}

func (u *User) GetCurrentNodeHash() string {
	return u.NodeStack.Peek().GetHash()
}

func (u *User) PushNode(hash string) {
	//log.Println(hash)
	u.NodeStack.Push(u.dicHashNode[hash])
}

func (u *User) PopNode() {
	u.NodeStack.Pop()
}

func (u *User) GetNumberOfChildren() int {
	nodes, err := u.Mega.FS.GetChildren(u.NodeStack.Peek())
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
	err := u.Mega.DownloadFile(u.NodeStack.Peek(), "/tmp/" + u.NodeStack.Peek().GetName(), ch)
	if err != nil {
		log.Println(err)
	}
}
