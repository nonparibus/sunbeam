package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/net/context"
)

const intro = `
<node>
	<interface name="com.github.pomdtr.Raycast">
		<method name="Show">
		</method>
		<method name="Sleep">
			<arg direction="in" type="u"/>
		</method>
	</interface>` + introspect.IntrospectDataString + `</node> `

type DbusServer struct {
	ctx context.Context
}

func NewDbusAPI(ctx context.Context) *DbusServer {
	return &DbusServer{ctx}
}

func (api *DbusServer) Show() *dbus.Error {
	runtime.WindowShow(api.ctx)
	return nil
}

func (api *DbusServer) Foo() (string, *dbus.Error) {
	return string("Foo"), nil
}

func (api *DbusServer) Listen() error {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return err
	}
	defer conn.Close()

	conn.Export(api, "/com/github/pomdtr/Raycast", "com.github.pomdtr.Raycast")
	conn.Export(introspect.Introspectable(intro), "/com/github/pomdtr/Raycast",
		"org.freedesktop.DBus.Introspectable")

	reply, err := conn.RequestName("com.github.pomdtr.Raycast",
		dbus.NameFlagDoNotQueue)
	if err != nil {
		return err
	}

	if reply != dbus.RequestNameReplyPrimaryOwner {
		return fmt.Errorf("name already taken")
	}
	runtime.LogInfo(api.ctx, "Listening on com.github.pomdtr.Raycast / /com/github/pomdtr/Raycast ...")
	select {}
}

func ShowWindow() {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}
	defer conn.Close()

	var s string
	obj := conn.Object("com.github.pomdtr.Raycast", "/com/github/pomdtr/Raycast")
	call := obj.Call("com.github.pomdtr.Raycast.Show", 0)
	if call.Err != nil {
		fmt.Fprintln(os.Stderr, "Failed to call Show function (is the server example running?):", err)
		os.Exit(1)
	}

	fmt.Println("Result from calling Show function on com.github.pomdtr.Raycast interface:")
	fmt.Println(s)
}
