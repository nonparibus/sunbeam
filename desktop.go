package main

import (
	"fmt"
	"os"
	"path"

	"github.com/rkoesters/xdg/desktop"
)

func ScanDesktopEntries() (map[string]*desktop.Entry, error) {
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
				println(fmt.Sprintf("failed to parse %s", entryPath))
				continue
			}
			if desktopEntry.Type != desktop.Application {
				continue
			}
			entryMap[entryPath] = desktopEntry
		}
	}

	return entryMap, nil
}
