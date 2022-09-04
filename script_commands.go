package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"regexp"
)

type ScriptCommand struct {
	Path string
	ScriptMetadatas
}

type ScriptMetadatas struct {
	Prefix               string
	SchemaVersion        string
	Title                string
	Mode                 string
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
	return SearchItem{
		Icon:           s.Icon,
		Title:          s.Title,
		AccessoryTitle: "Script Command",
		Actions: []Action{
			{Title: "Run Script", Command: RunScriptCommand(s.Path)},
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
	return ScriptCommand{Path: script_path, ScriptMetadatas: ScriptMetadatas{
		Title:       metadatas["title"],
		PackageName: metadatas["packageName"],
		Mode:        metadatas["mode"],
	}}, nil
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

type ScriptRunner struct {
	cmd *exec.Cmd
	io.ReadCloser
	io.WriteCloser
	*json.Decoder
}

func NewScriptRunner(cmd *exec.Cmd) (launcher *ScriptRunner) {
	return &ScriptRunner{
		cmd: cmd,
	}
}

func (p *ScriptRunner) Start() error {
	stdin, err := p.cmd.StdinPipe()
	if err != nil {
		return nil
	}
	p.WriteCloser = stdin

	stdout, err := p.cmd.StdoutPipe()
	if err != nil {
		return nil
	}
	p.ReadCloser = stdout
	p.Decoder = json.NewDecoder(stdout)

	if err = p.cmd.Start(); err != nil {
		return err
	}

	return nil
}

func (p *ScriptRunner) Close() {
	p.cmd.Process.Kill()
}

func IsExecOwner(mode os.FileMode) bool {
	return mode&0100 != 0
}
