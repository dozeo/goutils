package goutils

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ErikDubbelboer/gspt"
	"os"
	"os/exec"
	"time"
)

var Process Processes
var Command Processes // deprecated

type Processes struct {
}

func (c *Processes) Execute(timeout int, stdin []byte, parms ...string) ([]byte, []byte, error) {
	if len(parms) < 1 {
		return nil, nil, errors.New("execute: command missing")
	}
	cmd := exec.Command(parms[0], parms[1:]...)
	var bout bytes.Buffer
	var berr bytes.Buffer
	cmd.Stdout = &bout
	cmd.Stderr = &berr
	in, ierr := cmd.StdinPipe()
	if ierr != nil {
		return nil, nil, errors.New("execute: stdin failed")
	}
	cmd.Start()
	done := make(chan error)
	go func() {
		in.Write(stdin)
		in.Close()
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(time.Second * time.Duration(timeout)):
		go func() { <-done }() // allow goroutine to exit
		if err := cmd.Process.Kill(); err != nil {
			return bout.Bytes(), berr.Bytes(), errors.New(ErrorTimeoutKill)
		}
		return bout.Bytes(), berr.Bytes(), errors.New(ErrorTimeout)
	case status := <-done:
		out := berr.Bytes()
		if status != nil {
			return bout.Bytes(), out, status
		} else {
			return bout.Bytes(), out, nil
		}
	}
}

func (c *Processes) SetProcTitle(title string) {
	gspt.SetProcTitle(title)
}

func (c *Processes) AppendProcTitleAppend(title string) {
	gspt.SetProcTitle(fmt.Sprintf("%s%s", os.Args[0], title))
}
