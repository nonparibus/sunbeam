package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
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
			{Title: "Run Script", Command: NewRunCommand(s.Path)},
		},
	}
}

func extractValue(buf []byte, prefix string) []byte {
	r := regexp.MustCompile(fmt.Sprintf("%s (\\w+)$", prefix))
	return r.Find(buf)
}

func ParseScript(script_path string) ScriptCommand {
	content, _ := os.ReadFile(script_path)
	return ScriptCommand{Path: script_path, ScriptMetadatas: ScriptMetadatas{
		Title:       string(extractValue(content, "@raycast.title")),
		Description: string(extractValue(content, "@raycast.description")),
		PackageName: string(extractValue(content, "@raycast.packageName")),
		Mode:        string(extractValue(content, "@raycast.mode")),
	}}
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
