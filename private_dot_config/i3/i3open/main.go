package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mdirkse/i3ipc"
)

func main() {
	s := getState()

	cmd := ""
	flag.StringVar(&cmd, "cmd", "", "")
	flag.Parse()

	ipcsocket, err := i3ipc.GetIPCSocket()
	if err != nil {
		panic(err)
	}

	if strings.Contains(s, "1") {
		_, _ = ipcsocket.Command("split v")
		s = "2"
	} else if strings.Contains(s, "2") {
		_, _ = ipcsocket.Command("split h")
		s = "1"
	}

	t, _ := ipcsocket.GetTree()
	t.FindFocused()

	saveState(s)

	ok, err := ipcsocket.Command("exec --no-startup-id " + cmd)
	if err != nil {
		panic(err)
	}

	if ok {
		fmt.Println("ok")
	}
}

func getState() string {
	s, err := os.ReadFile("/home/dan/.config/i3/i3open/state.txt")
	if err != nil {
		panic(err)
	}

	return string(s)
}

func saveState(s string) {
	err := os.WriteFile("/home/dan/.config/i3/i3open/state.txt", []byte(s), 0644)
	if err != nil {
		panic(err)
	}
}
