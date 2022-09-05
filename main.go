package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

type Args struct {
	show bool
}

func parseArgs() Args {
	args := Args{}
	flag.BoolVar(&args.show, "show", false, "Show Raycast Window")
	flag.Parse()
	return args
}

func main() {
	args := parseArgs()
	if args.show {
		fmt.Print("Showing")
		os.Exit(0)
	}

	currentTheme, err := currentTheme()
	if err != nil {
		currentTheme = "Adwaita"
	}

	iconFinder := NewIconFinder()
	themes := []string{
		"hicolor",
		currentTheme,
	}
	for _, theme := range themes {
		for _, dir := range IconsDirectories() {
			themedir := path.Join(dir, theme)
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
	app.loadRootItems()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "Raycast",
		Assets: assets,
		// Frameless: true,
		// Buggy on linux
		// DisableResize: true,
		AssetsHandler: NewFileLoader(iconFinder),

		// AlwaysOnTop: true,
		Width:  750,
		Height: 475,

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
