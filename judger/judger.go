package judger

import (
	msgs "OnlineJudge/pbgen/messages"
)

type Judger struct {
	info *msgs.SubmitMQ
}

func NewJudger() *Judger {
	return &Judger{}
}

func (this *Judger) Reset() {
	*this = Judger{}
}

func (this *Judger) Init(sub *msgs.SubmitMQ) {
	this.info = sub
}

func (this *Judger) GetRunId() int64 {
	return this.info.GetRunId()
}

func (this *Judger) GetOJName() string {
	return this.info.GetOjName()
}

func (this *Judger) GetOJPid() string {
	return this.info.GetOjPid()
}

func (this *Judger) IsLocal() bool {
	if this.info.GetIsLocal() {
		return true
	}
	return false
}

func (this *Judger) IsVirtual() bool {
	if this.info.GetIsLocal() {
		return false
	}
	return true

}

func (this *Judger) IsSpj() bool {
	return this.info.GetIsSpj()
}

func (this *Judger) GetTestCasesBrief() []*msgs.TestCase {
	return this.info.GetTestcases()
}

func (this *Judger) GetLanguage() *msgs.SubmitLanguage {
	return this.info.GetLanguage()
}

// Need to talk to Daemon
func (this *Judger) GetSpjCode() *msgs.SpjCode {
	return nil
}

func (this *Judger) GetTestCase(id int64) *msgs.TestCase {
	return nil
}

func (this *Judger) UpdateStatus(status string, status_code string, cases int) {

}

func (this *Judger) UpdateUsage(time_used int, memory_used int) {

}

func (this *Judger) UpdateCEInfo(ce string) {

}
