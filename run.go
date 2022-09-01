package main

import (
	"encoding/json"
	"io"
	"os/exec"
)

type PopLauncher struct {
	cmd    *exec.Cmd
	stdin  io.ReadCloser
	stdout io.WriteCloser
}

func NewPopLauncher(path string) *PopLauncher {
	cmd := exec.Command(path)

	return &PopLauncher{
		cmd: cmd,
	}
}

func (p *PopLauncher) Run() (*json.Encoder, *json.Decoder, error) {
	stdin, err := p.cmd.StdinPipe()
	if err != nil {
		return nil, nil, err
	}

	stdout, err := p.cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}

	return json.NewEncoder(stdin), json.NewDecoder(stdout), err
}

func (p *PopLauncher) Shutdown() {
	p.cmd.Process.Kill()
	p.stdin.Close()
	p.stdout.Close()
}
