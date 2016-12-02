package base

import (
	"errors"
	"fmt"
	"regexp"
)

type Pid struct {
	OJId   int
	OJName string
	OJPid  int
}

func ParseSid(sid string) (*Pid, error) {
	regex, err := regexp.Compile(`(\w+)#(\d+)`)
	if err != nil {
		return nil, err
	}

	first := regex.FindSubmatch([]byte(sid))
	if len(first) != 3 {
		return nil, errors.New("invalid sid")
	}

	pid := &Pid{}
	fmt.Sscanf(string(first[1]), "%s", &pid.OJName)
	fmt.Sscanf(string(first[2]), "%d", &pid.OJPid)
	return pid, nil
}

func GenSid(pid *Pid) string {
	return fmt.Sprintf("%s#%d", pid.OJName, pid.OJPid)
}
