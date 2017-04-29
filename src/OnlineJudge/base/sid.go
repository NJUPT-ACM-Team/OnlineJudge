package base

import (
	"errors"
	"fmt"
	"regexp"
)

type Pid struct {
	OJId   int
	OJName string
	OJPid  string
}

func (this *Pid) OJPidToInt64() int64 {
	var id int64
	fmt.Sscanf(this.OJPid, "%d", &id)
	return id
}

func ParseSid(sid string) (*Pid, error) {
	regex, err := regexp.Compile(`(\w+)-(\w+)`)
	if err != nil {
		return nil, err
	}

	first := regex.FindSubmatch([]byte(sid))
	if len(first) != 3 {
		return nil, errors.New("invalid sid")
	}

	pid := &Pid{}
	fmt.Sscanf(string(first[1]), "%s", &pid.OJName)
	fmt.Sscanf(string(first[2]), "%s", &pid.OJPid)
	return pid, nil
}

func GenSid(pid *Pid) string {
	return fmt.Sprintf("%s-%s", pid.OJName, pid.OJPid)
}
