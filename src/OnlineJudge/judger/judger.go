package judger

import (
	"OnlineJudge/base"
	"OnlineJudge/irpc"
	msgs "OnlineJudge/pbgen/messages"

	"time"
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

func (this *Judger) GetTimeLimit() int {
	return int(this.info.GetTimeLimit())
}

func (this *Judger) GetSubmitTime() time.Time {
	t, err := base.UnmarshalTime(this.info.GetSubmitTime())
	if err != nil {
		return base.GetDefaultTime()
	}
	return t
}

func (this *Judger) UpdateResource(t, m int) error {
	subs := &irpc.SubmissionStatus{
		TimeUsed:   int32(t),
		MemoryUsed: int32(m),
	}
	return this.UpdateStatus(subs)
}

func (this *Judger) UpdateResult(s, sc string) error {
	subs := &irpc.SubmissionStatus{
		Status:     s,
		StatusCode: sc,
	}
	return this.UpdateStatus(subs)
}

func (this *Judger) GetMemoryLimit() int {
	return int(this.info.GetMemoryLimit())
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

func (this *Judger) GetCode() string {
	return this.info.GetCode()
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

func (this *Judger) UpdateStatus(subs *irpc.SubmissionStatus) error {
	subs.RunId = this.GetRunId()
	helper := irpc.NewHelper()
	if err := helper.Connect(); err != nil {
		return err
	}
	defer helper.Disconnect()

	helper.NewClient()

	// Set judging
	_, err := helper.UpdateSubmissionStatus(subs)
	if err != nil {
		return err
	}
	return nil
}

func (this *Judger) UpdateStatusJudging() error {
	subs := &irpc.SubmissionStatus{
		Status:     "Judging",
		StatusCode: "wt",
	}
	return this.UpdateStatus(subs)
}

func (this *Judger) UpdateCEInfo(ce string) error {
	subs := &irpc.SubmissionStatus{
		CEInfo: ce,
	}
	return this.UpdateStatus(subs)
}

func (this *Judger) UpdateUsage(time_used int, memory_used int) {

}
