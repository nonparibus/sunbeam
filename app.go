package main

import (
	"context"
	"fmt"

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

func (a *App) Search(query string) {
	runtime.LogDebug(a.ctx, fmt.Sprintf("User Query: %s", query))
	a.launcher.Encode(SearchRequest{query})
}

func (a *App) Activate(itemID int) {
	a.launcher.Encode(ActivateRequest{itemID})
}

func (a *App) emitUpdates() {
	for a.launcher.Scan() {
		update := a.launcher.Text()
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

func (a *App) domReady(ctx context.Context) {
	go a.emitUpdates()
}

func (a *App) shutdown(ctx context.Context) {
	a.launcher.Close()
}
