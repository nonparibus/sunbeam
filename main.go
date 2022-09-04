package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	var err error

	currentTheme, err := currentTheme()
	if err != nil {
		currentTheme = "Adwaita"
	}

	iconFinder := NewIconFinder()
	themes := []string{
		"hicolor",
		"Humanity",
		currentTheme,
	}
	for _, theme := range themes {
		for _, dir := range IconsDirectories() {
			themedir := path.Join(dir, theme)
			println(themedir)
			if _, err := os.Stat(themedir); os.IsNotExist(err) {
				continue
			}
			err = iconFinder.loadThemeIcons(fmt.Sprintf("/usr/share/icons/%s", theme))
			if err != nil {
				log.Fatalf("Theme not found: %s", theme)
			}

		}
	}

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title:     "Raycast",
		Assets:    assets,
		Frameless: true,
		// Buggy on linux
		// DisableResize: true,
		AssetsHandler: NewFileLoader(iconFinder),

		AlwaysOnTop: true,
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
