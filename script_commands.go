package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"regexp"

	"github.com/adrg/xdg"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ScriptCommand struct {
	Path string
	ScriptMetadatas
}

type ScriptMetadatas struct {
	SchemaVersion        int    `validate:"required,eq=1"`
	Title                string `validate:"required"`
	Mode                 string `validate:"required,oneof=silent filter search"`
	PackageName          string
	Icon                 string
	IconDark             string
	CurrentDirectoryPath string
	NeedsConfirmation    bool
	Author               string
	AutorUrl             string
	Description          string
}

func (s *ScriptCommand) toSearchItem() SearchItem {
	var primaryAction Action
	var accessoryTitle string
	if s.Mode == "filter" || s.Mode == "search" {
		primaryAction = Action{
			Title:   "Open Command",
			Command: NewPushListCommand(s.Path, []string{}, string(s.Mode)),
		}
		accessoryTitle = "Command"
	} else {
		primaryAction = Action{
			Title:   "Run Script",
			Command: NewRunScriptCommand(s.Path, []string{}, string(s.Mode)),
		}
		accessoryTitle = "Script Command"
	}
	return SearchItem{
		IconSource:     s.Icon,
		Title:          s.Title,
		AccessoryTitle: accessoryTitle,
		Actions: []Action{
			primaryAction,
		},
	}
}

func extractRaycastMetadatas(content string) map[string]string {
	r := regexp.MustCompile("@raycast.([A-Za-z0-9]+)\\s([\\S ]+)")
	groups := r.FindAllStringSubmatch(content, -1)

	metadataMap := make(map[string]string)
	for _, group := range groups {
		metadataMap[group[1]] = group[2]
	}

	return metadataMap
}

func ParseScript(script_path string) (ScriptCommand, error) {
	content, err := ioutil.ReadFile(script_path)
	if err != nil {
		return ScriptCommand{}, err
	}

	metadatas := extractRaycastMetadatas(string(content))

	var schemaVersion int
	err = json.Unmarshal([]byte(metadatas["schemaVersion"]), &schemaVersion)
	if err != nil {
		return ScriptCommand{}, err
	}

	var needsConfirmation bool
	json.Unmarshal([]byte(metadatas["schemaVersion"]), &needsConfirmation)

	scripCommand := ScriptCommand{Path: script_path, ScriptMetadatas: ScriptMetadatas{
		SchemaVersion:        schemaVersion,
		Title:                metadatas["title"],
		Mode:                 metadatas["mode"],
		PackageName:          metadatas["packageName"],
		Icon:                 metadatas["icon"],
		IconDark:             metadatas["iconDark"],
		CurrentDirectoryPath: metadatas["currentDirectoryPath"],
		NeedsConfirmation:    needsConfirmation,
		Author:               metadatas["author"],
		AutorUrl:             metadatas["autorUrl"],
		Description:          metadatas["description"],
	}}

	err = validate.Struct(scripCommand)
	if err != nil {
		println(err)
		return ScriptCommand{}, err
	}

	return scripCommand, nil
}

func ScanScriptDir(scriptDir string) ([]ScriptCommand, error) {
	dirEntries, _ := os.ReadDir(scriptDir)

	scriptCommands := make([]ScriptCommand, 0, len(dirEntries))
	for _, dirEntry := range dirEntries {
		scriptPath := path.Join(scriptDir, dirEntry.Name())
		scriptCommand, err := ParseScript(scriptPath)
		if err != nil {
			return nil, err
		}
		scriptCommands = append(scriptCommands, scriptCommand)
	}

	return scriptCommands, nil
}

func ScanScriptDirs() (scriptCommands []ScriptCommand, err error) {
	for _, dir := range xdg.DataDirs {
		scriptDir := path.Join(dir, "raycast")
		if _, err := os.Stat(scriptDir); os.IsNotExist(err) {
			continue
		}
		dirCommands, err := ScanScriptDir(scriptDir)
		if err != nil {
			return nil, err
		}
		scriptCommands = append(scriptCommands, dirCommands...)
	}

	return
}

func IsExecOwner(mode os.FileMode) bool {
	return mode&0100 != 0
}
