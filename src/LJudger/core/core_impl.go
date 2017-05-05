package core

import (
	"OnlineJudge/irpc"

	"os/exec"
	"strconv"
	"strings"
	//	"sync"
)

type CoreImpl struct {
	CorePath string
	Mode     *Mode
	out      string
}

func NewCore(path string) Core {
	return &CoreImpl{CorePath: path}
}

func (this *CoreImpl) SetMode(m *Mode) {
	this.Mode = m
}

func (this *CoreImpl) GetSubmissionStatus() *irpc.SubmissionStatus {
	// return &irpc.SubmissionStatus{Status: this.out}
	// TODO: parse result
	return &irpc.SubmissionStatus{Status: "Accepted", StatusCode: "ac", TimeUsed: 90}
}

func (this *CoreImpl) Run() error {
	// wg := new(sync.WaitGroup)
	cmds := []string{}
	cmds = append(cmds, "sudo", this.CorePath)
	cmds = append(cmds, "-c", this.Mode.SrcPath)
	cmds = append(cmds, "-t", strconv.Itoa(this.Mode.TimeLimit))
	cmds = append(cmds, "-m", strconv.Itoa(this.Mode.MemoryLimit))
	cmds = append(cmds, "-d", this.Mode.RunDir)
	cmd := strings.Join(cmds, " ")

	out, err := exec_command(cmd)
	if err != nil {
		return err
	}
	this.out = out
	return nil
}

func exec_command(cmd string) (string, error) {
	out, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
