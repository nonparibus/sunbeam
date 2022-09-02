package main

import (
	"encoding/json"
	"io"
	"os/exec"
)

type PopLauncher struct {
	cmd *exec.Cmd
	io.ReadCloser
	io.WriteCloser
	*json.Encoder
	*json.Decoder
}

func NewPopLauncher(cmd *exec.Cmd) (launcher *PopLauncher) {
	return &PopLauncher{
		cmd: cmd,
	}
}

func (p *PopLauncher) Start() error {
	stdin, err := p.cmd.StdinPipe()
	if err != nil {
		return nil
	}
	p.WriteCloser = stdin
	p.Encoder = json.NewEncoder(stdin)

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

func (p *PopLauncher) Close() {
	p.ReadCloser.Close()
	p.WriteCloser.Close()
	p.cmd.Process.Kill()
}
