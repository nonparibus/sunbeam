package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	var err error

	launcherPath := "/usr/bin/pop-launcher"
	if os.Stat(launcherPath); os.IsNotExist(err) {
		log.Fatalf("Pop Launcher not found!")
	}
	launcherCmd := exec.Command(launcherPath)
	launcher := NewPopLauncher(launcherCmd)

	iconFinder := NewIconFinder()

	for _, theme := range []string{"hicolor", "Humanity", "Adwaita"} {
		err = iconFinder.loadThemeIcons(fmt.Sprintf("/usr/share/icons/%s", theme))
		if err != nil {
			log.Fatalf("Theme not found: %s", theme)
		}
	}

	if err != nil {
		log.Fatalf("Unable to load theme: %s", err)
	}

	// Create an instance of the app structure
	app := NewApp(launcher)

	// Create application with options
	err = wails.Run(&options.App{
		Title:     "raycast-linux",
		Assets:    assets,
		Frameless: true,
		// Buggy on linux
		// DisableResize: true,
		AssetsHandler: NewFileLoader(iconFinder),

		AlwaysOnTop: false,
		Width:       750,
		Height:      475,

		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		OnDomReady: app.ready,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
