package main

import (
	"bufio"
	"encoding/json"
	"io"
	"os/exec"
)

type PopLauncher struct {
	cmd *exec.Cmd
	io.ReadCloser
	io.WriteCloser
	*json.Encoder
	*bufio.Scanner
}

func NewPopLauncher(path string) (launcher *PopLauncher) {
	return &PopLauncher{
		cmd: exec.Command(path),
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
	p.Scanner = bufio.NewScanner(stdout)

	return nil
}

func (p *PopLauncher) Close() {
	p.ReadCloser.Close()
	p.WriteCloser.Close()
	p.cmd.Process.Kill()
}
