package main

import (
	"embed"
	"os"
	"os/exec"
	"path"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	homedir, _ := os.UserHomeDir()
	launcherPath := path.Join(homedir, ".local", "bin", "pop-launcher")
	launcherCmd := exec.Command(launcherPath)
	launcher := NewPopLauncher(launcherCmd)

	// Create an instance of the app structure
	app := NewApp(launcher)

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "raycast-linux",
		Assets:    assets,
		Frameless: true,
		// Buggy on linux
		// DisableResize: true,

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
