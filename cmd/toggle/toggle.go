package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
)

func main() {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}
	defer conn.Close()

	var s string
	obj := conn.Object("com.github.pomdtr.Raycast", "/com/github/pomdtr/Raycast")
	call := obj.Call("com.github.pomdtr.Raycast.Toggle", 0)
	if call.Err != nil {
		fmt.Fprintln(os.Stderr, "Failed to call Toggle function (is the server example running?):", err)
		os.Exit(1)
	}

	fmt.Println("Result from calling Toggle function on com.github.pomdtr.Raycast interface:")
	fmt.Println(s)
}
