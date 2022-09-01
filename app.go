package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx     context.Context
	encoder *json.Encoder
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) Search(query string) {
	runtime.LogDebug(a.ctx, fmt.Sprintf("User Query: %s", query))
	a.encoder.Encode(SearchCommand{query})
}

func SendResponses(ctx context.Context, decoder *json.Decoder) {
	for {
		var response UpdateResponse
		err := decoder.Decode(&response)
		if err != nil {
			continue
		}
		runtime.LogDebug(ctx, fmt.Sprintf("Sent %d items", len(response.Update)))
		runtime.EventsEmit(ctx, "update", response.Update)
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	homedir, _ := os.UserHomeDir()
	launcherPath := path.Join(homedir, ".local", "bin", "pop-launcher")
	launcher := NewPopLauncher(launcherPath)
	err := launcher.Start()
	if err != nil {
		log.Fatal("Unable to start launcher")
	}

	a.encoder = launcher.Encoder
	go SendResponses(a.ctx, launcher.Decoder)
}
