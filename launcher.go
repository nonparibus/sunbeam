package main

import (
	"encoding/json"
	"os/exec"
)

type PopLauncher struct {
	cmd *exec.Cmd
	*json.Encoder
	*json.Decoder
}

func NewPopLauncher(path string) (launcher *PopLauncher) {
	return &PopLauncher{
		cmd: exec.Command(path),
	}
}

func (launcher *PopLauncher) Start() (err error) {
	stdin, err := launcher.cmd.StdinPipe()
	if err != nil {
		return
	}
	launcher.Encoder = json.NewEncoder(stdin)

	stdout, err := launcher.cmd.StdoutPipe()
	if err != nil {
		return
	}
	launcher.Decoder = json.NewDecoder(stdout)

	return launcher.cmd.Start()
}

func (p *PopLauncher) Stop() {
	p.cmd.Process.Kill()
}
