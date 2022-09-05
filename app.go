package main

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path"

	"github.com/atotto/clipboard"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx       context.Context
	rootItems []SearchItem
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) loadRootItems() (err error) {
	entryMap, err := ScanDesktopEntries()
	if err != nil {
		return
	}

	searchItems := make([]SearchItem, 0, len(entryMap))
	for desktopEntryPath, desktopEntry := range entryMap {
		searchItems = append(searchItems, SearchItem{
			Title:          desktopEntry.Name,
			AccessoryTitle: "Application",
			IconSource:     desktopEntry.Icon,
			Actions: []Action{
				{Title: "Open Application", Icon: "/raycast/icon-app-window-16.svg", Command: NewOpenCommand(desktopEntryPath)},
				{Title: "Copy Path", Icon: "/raycast/icon-copy-clipboard-16.svg", Command: NewCopyToClipboardCommand(desktopEntryPath)},
			},
		})
	}

	scriptCommands, err := ScanScriptDirs()
	if err != nil {
		return
	}

	for _, scriptCommand := range scriptCommands {
		searchItems = append(searchItems, scriptCommand.toSearchItem())
	}

	a.rootItems = searchItems
	return
}

func (a *App) RootItems() []SearchItem {
	return a.rootItems
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
