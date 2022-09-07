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
	rootItems []ListItem
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) loadRootItems() (err error) {
	searchItems := make([]ListItem, 0)

	scriptCommands, err := ScanScriptDirs()
	if err != nil {
		return
	}
	for _, scriptCommand := range scriptCommands {
		searchItems = append(searchItems, scriptCommand.toSearchItem())
	}

	entryMap, err := ScanDesktopEntries()
	if err != nil {
		return
	}
	for desktopEntryPath, desktopEntry := range entryMap {
		searchItems = append(searchItems, ListItem{
			Title:          desktopEntry.Name,
			AccessoryTitle: "Application",
			IconSource:     desktopEntry.Icon,
			Actions: []Action{
				NewOpenAction("Open Application", desktopEntryPath),
				NewCopyToClipboardAction("Copy Path", desktopEntryPath),
			},
		})
	}

	a.rootItems = searchItems
	return
}

func (a *App) RootItems() []ListItem {
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

func (a *App) RunListCommand(scriptPath string, args []string) (res ScriptResponse, err error) {
	err = os.Chmod(scriptPath, 0755)
	if err != nil {
		return
	}

	cmd := exec.Command(scriptPath, args...)
	cmd.Dir = path.Dir(scriptPath)
	output, err := cmd.Output()
	if err != nil {
		return
	}
	json.Unmarshal(output, &res)
	for _, item := range res.List.Items {
		for i, action := range item.Actions {
			switch action.Type {
			case "callback":
				item.Actions[i] = NewRunCommandAction(action.Title, scriptPath, action.Args...)
			}
		}
	}

	err = validate.Struct(res)
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
