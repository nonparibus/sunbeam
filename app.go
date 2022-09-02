package main

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx      context.Context
	launcher *PopLauncher
}

// NewApp creates a new App application struct
func NewApp(launcher *PopLauncher) *App {
	return &App{launcher: launcher}
}

// TODO: return most used apps if query is empty
func (a *App) Search(query string) {
	runtime.LogDebug(a.ctx, fmt.Sprintf("User Query: %s", query))
	a.launcher.Encode(SearchRequest{query})
}

func (a *App) ListRecentApplications(query string) {

}

func (a *App) Activate(itemID int) {
	runtime.LogDebug(a.ctx, fmt.Sprintf("Activate Item: %d", itemID))
	a.launcher.Encode(ActivateRequest{itemID})
}

func (a *App) OpenApp(path string) (err error) {
	cmd := exec.Command("xdg-open", path)
	err = cmd.Run()
	if err != nil {
		runtime.LogErrorf(a.ctx, "An error occured when opening %s", err)
	}
	return
}

func (a *App) Complete(itemID int) {
	runtime.LogDebug(a.ctx, fmt.Sprintf("Complete Item: %d", itemID))
	a.launcher.Encode(CompleteRequest{itemID})
}

func (a *App) Quit(itemID int) {
	a.launcher.Encode(QuitRequest{itemID})
}

func (a *App) Context(itemID int) {
	runtime.LogDebug(a.ctx, fmt.Sprintf("Context Item: %d", itemID))
	a.launcher.Encode(ContextRequest{itemID})
}

func (a *App) ActivateContext(itemID int, contextID int) {
	runtime.LogDebug(a.ctx, fmt.Sprintf("Activate Context for item: %d", itemID))
	a.launcher.Encode(ActivateContextRequest{ActivateContext{
		itemID,
		contextID,
	}})
}

func (a *App) emitUpdates() {
	var update interface{}
	for {
		err := a.launcher.Decode(&update)
		if err != nil {
			runtime.LogFatalf(a.ctx, "Decoding error: %s", err)
			panic(err)
		}
		runtime.EventsEmit(a.ctx, "update", update)
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	err := a.launcher.Start()
	if err != nil {
		runtime.LogFatalf(ctx, "Unable to start launcher: %s", err)
	}
}

func (a *App) ready(ctx context.Context) {
	runtime.LogDebugf(a.ctx, "Starting emit Loop")
	go a.emitUpdates()
}

func (a *App) shutdown(ctx context.Context) {
	a.launcher.Close()
}
