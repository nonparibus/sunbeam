package main

import (
	"fmt"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/net/context"
)

const intro = `
<node>
	<interface name="com.github.pomdtr.Raycast">
		<method name="Toggle">
		</method>
		<method name="Sleep">
			<arg direction="in" type="u"/>
		</method>
	</interface>` + introspect.IntrospectDataString + `</node> `

type DbusAPI struct {
	ctx context.Context
}

func NewDbusAPI(ctx context.Context) *DbusAPI {
	return &DbusAPI{ctx}
}

func (api *DbusAPI) Toggle() *dbus.Error {
	runtime.EventsEmit(api.ctx, "toggle")
	return nil
}

func (api *DbusAPI) Foo() (string, *dbus.Error) {
	return string("Foo"), nil
}

func (api *DbusAPI) Listen() error {
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
	fmt.Println("Listening on com.github.pomdtr.Raycast / /com/github/pomdtr/Raycast ...")

	select {}
}
