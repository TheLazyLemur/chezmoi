package main

import (
	"github.com/mdirkse/i3ipc"
)

const (
	VERTICAL   = "vertical"
	HORIZONTAL = "horizontal"
	NONE       = "none"
)

type i3ipcClient struct {
	ipcsocket *i3ipc.IPCSocket
}

func NewI3ipcClient() *i3ipcClient {
	ipcsocket, err := i3ipc.GetIPCSocket()
	if err != nil {
		panic(err)
	}

	return &i3ipcClient{ipcsocket}
}

func (c *i3ipcClient) GetTree() (root i3ipc.I3Node, err error) {
	return c.ipcsocket.GetTree()
}

func (c *i3ipcClient) Close() error {
	return c.ipcsocket.Close()
}

func (c *i3ipcClient) splitVertical() {
	ok, err := c.ipcsocket.Command("split v")
	if err != nil {
		panic(err)
	}

	if !ok {
		panic("split failed")
	}
}

func (c *i3ipcClient) splitHorizontal() {
	ok, err := c.ipcsocket.Command("split h")
	if err != nil {
		panic(err)
	}

	if !ok {
		panic("split failed")
	}
}

func toggleSplit() {
	c := NewI3ipcClient()

	defer func() {
		err := c.Close()
		if err != nil {
			panic(err)
		}
	}()

	node, err := c.GetTree()
	if err != nil {
		panic(err)
	}

	win := node.FindFocused()

	switch win.Parent.Orientation {
	case VERTICAL:
		c.splitHorizontal()
	case HORIZONTAL:
		c.splitVertical()
	case NONE:
		c.splitHorizontal()
	default:
		panic("unknown orientation")
	}
}

func main() {
	toggleSplit()
}
