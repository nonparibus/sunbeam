package main

import (
	"context"
	"os"
	"os/exec"
	"path"

	"github.com/atotto/clipboard"
	"github.com/rkoesters/xdg/desktop"
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
	homedir, _ := os.UserHomeDir()
	directories := []string{
		"/usr/share/applications",
		"/usr/local/share/applications",
		path.Join(homedir, ".local", "share", "applications"),
	}

	entryMap := make(map[string]*desktop.Entry)
	for _, directory := range directories {
		dirEntries, _ := os.ReadDir(directory)
		for _, dirEntry := range dirEntries {
			entryPath := path.Join(directory, dirEntry.Name())
			file, _ := os.Open(entryPath)
			desktopEntry, err := desktop.New(file)
			if err != nil {
				continue
			}
			if desktopEntry.Terminal {
				continue
			}
			entryMap[entryPath] = desktopEntry
		}
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

func RunScript(scriptPath string, args []string) (err error) {
	cmd := exec.Command(scriptPath, args...)
	return cmd.Run()
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
