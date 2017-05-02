package ljudger

import (
	"OnlineJudge/base"
	//	"OnlineJudge/irpc"
	"OnlineJudge/judger"

	"fmt"
	"log"
	"path"
)

const (
	SRCFILE = "src"
)

// manual judge
type Result struct {
	Status     string
	StatusCode string
}

/*
func ManualJudge(oj string, pid string, src string, lang string) *Result {
	fmt.Printf("Problem Sid: %s-%s\n", oj, pid)
	fmt.Printf("Language:%s\nCode:\n%s\n", lang, src)
	fmt.Printf("1.ac\n2.wa\nchoice:")
	var in int
	fmt.Scanf("%d", &in)
	fmt.Println(in)
	switch in {
	case 1:
		return &Result{Status: "Accepted", StatusCode: "ac"}
	case 2:
		return &Result{Status: "Wrong Answer", StatusCode: "wa"}
	}
	return &Result{Status: "System Error", StatusCode: "se"}
}
*/

type LocalJudger struct {
	Dir      string
	Jdi      judger.JudgerInterface
	JudgeDir string
	RunDir   string
	InDir    string
	OutDir   string
}

func NewLocalJudger(dir string, jdi judger.JudgerInterface) *LocalJudger {
	return &LocalJudger{
		Dir: dir,
		Jdi: jdi,
	}
}

func (this *LocalJudger) InitDir() error {
	this.JudgeDir = path.Join(this.Dir, fmt.Sprintf("%d", this.Jdi.GetRunId()))
	this.RunDir = path.Join(this.JudgeDir, "run")
	this.InDir = path.Join(this.JudgeDir, "data", "in")
	this.OutDir = path.Join(this.JudgeDir, "data", "out")

	if base.DirExists(this.JudgeDir) {
		if err := base.RemoveDir(this.JudgeDir); err != nil {
			return err
		}
	}
	if err := base.MakeDirs(this.JudgeDir); err != nil {
		return err
	}
	if err := base.MakeDirs(this.RunDir); err != nil {
		return err
	}
	if err := base.MakeDirs(this.InDir); err != nil {
		return err
	}
	if err := base.MakeDirs(this.OutDir); err != nil {
		return err
	}
	return nil
}

func (this *LocalJudger) PrepareData() error {
	// source code
	sfx := this.Jdi.GetLanguage().GetSuffix()
	src_file := "src." + sfx
	src_path := path.Join(this.JudgeDir, src_file)
	if err := base.WriteFile(src_path, []byte(this.Jdi.GetCode())); err != nil {
		return err
	}

	// spj code
	/*
		if this.Jdi.GetIsSpj() {
		}
	*/

	// testing data
	for k, v := range this.Jdi.GetTestCasesBrief() {
		in_file := fmt.Sprintf("%d.in", k)
		out_file := fmt.Sprintf("%d.out", k)
		in_path := path.Join(this.InDir, in_file)
		out_path := path.Join(this.OutDir, out_file)
		if err := base.WriteFile(in_path, v.GetInput()); err != nil {
			return err
		}
		if err := base.WriteFile(out_path, v.GetOutput()); err != nil {
			return err
		}
	}
	return nil
}

func EntryPoint(jdi judger.JudgerInterface) {
	log.Println("received run_id=", jdi.GetRunId())
	lj := NewLocalJudger("/tmp/testoj", jdi)
	if err := lj.InitDir(); err != nil {
		log.Fatal(err)
	}
	if err := lj.PrepareData(); err != nil {
		log.Fatal(err)
	}
}

/*
func EntryPoint(jdi judger.JudgerInterface) {
	log.Println(jdi.GetRunId())
	log.Println(jdi.GetCode())

	helper := irpc.NewHelper()
	if err := helper.Connect(); err != nil {
		// Log the error
		log.Println(err)
		return
	}
	defer helper.Disconnect()

	helper.NewClient()

	// Set judging
	res, err := helper.UpdateSubmissionStatus(&irpc.SubmissionStatus{RunId: jdi.GetRunId(), Status: "Judging", StatusCode: "wt"})
	if err != nil {
		log.Println(err)
	}
	log.Println(res)

	// Use manual judge for demo
	j_res := ManualJudge(jdi.GetOJName(), jdi.GetOJPid(), jdi.GetCode(), jdi.GetLanguage().GetLang())
	subs := &irpc.SubmissionStatus{RunId: jdi.GetRunId(), Status: j_res.Status, StatusCode: j_res.StatusCode}

	res, err = helper.UpdateSubmissionStatus(subs)
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
}
*/

type VJudger interface {
	Init(judger.JudgerInterface) error
	Login(judger.JudgerInterface) error
	Submit(judger.JudgerInterface) error
	GetStatus(judger.JudgerInterface) error
	Run(judger.JudgerInterface) error
	Match(string) bool
	// Crawler(string) error
}
