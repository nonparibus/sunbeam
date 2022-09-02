package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/adrg/xdg"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

type FileLoader struct {
	http.Handler
}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

var iconThemes = []string{
	"Adwaita",
	"hicolor",
}

func getSystemIconPath(iconName string) (string, error) {
	return "/usr/share/icons/hicolor/24x24/apps/Thunar.png", nil
	var iconPath string
	for _, xdgDataDir := range xdg.DataDirs {
		for _, theme := range iconThemes {
			iconPath = fmt.Sprintf("%s/icons/%s/24x24/apps/%s", xdgDataDir, theme, iconName)
			if _, err := os.Stat(iconPath); !os.IsNotExist(err) {
				return iconPath, nil
			}
			iconPath = fmt.Sprintf("%s/icons/%s/apps/24/%s", xdgDataDir, theme, iconName)
			if _, err := os.Stat(iconPath); !os.IsNotExist(err) {
				return iconPath, nil
			}
		}
	}
	return "", fmt.Errorf("No icon found for %s", iconName)
}

func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	iconName := strings.TrimPrefix(req.URL.Path, "/")
	// iconType := req.URL.Query().Get("type")
	iconPath, err := getSystemIconPath(iconName)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not find icon %s", iconName)))
	}

	fileData, err := os.ReadFile(iconPath)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not load icon %s", iconName)))
	}

	res.Header().Set("content-type", "image/png")
	res.Write(fileData)
}

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
		AssetsHandler: NewFileLoader(),

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
