package main

import (
	"context"
	"os"
	"os/exec"
	"path"

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

func (a *App) RootItems() (Response, error) {
	entryMap, err := ScanDesktopEntries()
	if err != nil {
		return Response{}, err
	}

	searchItems := make([]SearchItem, 0, len(entryMap))
	for desktopEntryPath, desktopEntry := range entryMap {
		runtime.LogDebugf(a.ctx, "%s", desktopEntry)
		searchItems = append(searchItems, SearchItem{
			Key:            desktopEntryPath,
			Title:          desktopEntry.Name,
			AccessoryTitle: "Application",
			Icon:           desktopEntry.Icon,
			Actions: []Action{
				{Title: "Open Application", Command: NewOpenCommand(desktopEntryPath)},
				{Title: "Copy Name", Command: NewCopyToClipboardCommand(desktopEntry.Name)},
			},
		})
	}

	scriptCommands, err := ScanScriptDir()
	for _, scriptCommand := range scriptCommands {
		searchItems = append(searchItems, SearchItem{
			Title:          scriptCommand.Title,
			Subtitle:       scriptCommand.PackageName,
			AccessoryTitle: "Script Command",
			Actions: []Action{
				{Title: "Run Script", Command: NewRunCommand(scriptCommand.Path)},
			},
		})
	}

	return Response{
		Type:  Filter,
		Items: searchItems,
	}, nil
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
		return err
	}
	cmd := exec.Command(scriptPath, args...)
	cmd.Dir = path.Dir(scriptPath)
	err = cmd.Run()
	return err

}

func PushScript(scriptPath string) Response {
	return Response{}
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
