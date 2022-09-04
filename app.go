package main

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path"

	"github.com/adrg/xdg"
	"github.com/atotto/clipboard"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) RootItems() ([]SearchItem, error) {
	entryMap, err := ScanDesktopEntries()
	if err != nil {
		return nil, err
	}

	searchItems := make([]SearchItem, 0, len(entryMap))
	for desktopEntryPath, desktopEntry := range entryMap {
		searchItems = append(searchItems, SearchItem{
			Key:            desktopEntryPath,
			Title:          desktopEntry.Name,
			AccessoryTitle: "Application",
			Icon:           desktopEntry.Icon,
			Actions: []Action{
				{Title: "Open Application", Command: NewOpenCommand(desktopEntryPath)},
				{Title: "Copy Path", Command: NewCopyToClipboardCommand(desktopEntryPath)},
			},
		})
	}

	scriptDir := path.Join(xdg.DataHome, "raycast", "scripts")
	scriptCommands, err := ScanScriptDir(scriptDir)
	for _, scriptCommand := range scriptCommands {
		searchItems = append(searchItems, SearchItem{
			Title:          scriptCommand.Title,
			Subtitle:       scriptCommand.PackageName,
			AccessoryTitle: "Script Command",
			Actions: []Action{
				{Title: "Run Script", Command: RunScriptCommand(scriptCommand.Path)},
			},
		})
	}

	return searchItems, nil
}

func (a *App) OpenFile(filePath string) error {
	return exec.Command("xdg-open", filePath).Run()
}

func (a *App) OpenInBrowser(url string) error {
	runtime.BrowserOpenURL(a.ctx, url)
	return nil
}

func (a *App) CopyToClipboard(content string) error {
	return clipboard.WriteAll(content)
}

func (a *App) RunScript(scriptPath string, args []string) (err error) {
	err = os.Chmod(scriptPath, 0755)
	if err != nil {
		return
	}
	cmd := exec.Command(scriptPath, args...)
	cmd.Dir = path.Dir(scriptPath)
	err = cmd.Run()
	return

}

func (a *App) RunListCommand(scriptPath string) (items []SearchItem, err error) {
	err = os.Chmod(scriptPath, 0755)
	if err != nil {
		return
	}

	cmd := exec.Command(scriptPath)
	cmd.Dir = path.Dir(scriptPath)
	output, err := cmd.Output()
	if err != nil {
		return
	}

	err = json.Unmarshal(output, &items)
	return
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	var err error
	a.ctx = ctx

	dbusAPI := NewDbusAPI(a.ctx)
	go dbusAPI.Listen()
	if err != nil {
		runtime.LogFatalf(ctx, "Unable to start dbus api: %s", err)
	}
}

func (a *App) ready(ctx context.Context) {
	runtime.LogDebugf(a.ctx, "Starting emit Loop")
}

func (a *App) shutdown(ctx context.Context) {
}
