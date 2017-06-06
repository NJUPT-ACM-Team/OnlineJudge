package core

import (
	"OnlineJudge/irpc"

	"encoding/json"
	"io/ioutil"
	"log"
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

type Json struct {
	Memory  int
	PassNum int
	Result  string
	Time    int
	CE      string
}

func NewCore(path string) Core {
	return &CoreImpl{CorePath: path}
}

func (this *CoreImpl) SetMode(m *Mode) {
	this.Mode = m
}

func GetStatusByStatusCode(sc string) string {
	var status string
	switch strings.ToLower(sc) {
	case "ac":
		status = "Accepted"
	case "ce":
		status = "Compile Error"
	case "tle":
		status = "Time Limit Exceed"
	case "mle":
		status = "Memory Limit Exceed"
	case "ole":
		status = "Output Limit Exceed"
	case "wa":
		status = "Wrong Answer"
	case "pe":
		status = "Presentation Error"
	case "se":
		status = "System Error"
	case "re":
		status = "Runtime Error"
	default:
		status = "Unknown"
	}
	return status
}

func (this *CoreImpl) GetSubmissionStatus() (*irpc.SubmissionStatus, error) {
	// return &irpc.SubmissionStatus{Status: this.out}
	// TODO: parse result
	data, err := ioutil.ReadFile(this.Mode.ResPath)
	if err != nil {
		return nil, err
	}
	c := &Json{}
	if err := json.Unmarshal(data, c); err != nil {
		return nil, err
	}
	return &irpc.SubmissionStatus{
		Status:          GetStatusByStatusCode(strings.ToLower(c.Result)),
		StatusCode:      strings.ToLower(c.Result),
		TimeUsed:        int32(c.Time),
		MemoryUsed:      int32(c.Memory),
		CEInfo:          c.CE,
		TestcasesPassed: int32(c.PassNum),
	}, nil
}

func (this *CoreImpl) Run() error {
	// wg := new(sync.WaitGroup)
	cmds := []string{}
	if this.Mode.UseSudo {
		cmds = append(cmds, "sudo")
	}
	cmds = append(cmds, this.CorePath)
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
	log.Println("cmd:", cmd)
	out, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
