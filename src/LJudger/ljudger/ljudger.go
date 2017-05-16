package ljudger

import (
	"OnlineJudge/base"
	//	"OnlineJudge/irpc"
	"LJudger/core"
	"OnlineJudge/judger"

	"fmt"
	"log"
	"path"
)

const (
	COREPATH   = "/home/kevince/OnlineJudgeCore/judger"
	JUDGEROOT  = "/tmp/testoj"
	RESULTFILE = "result.json"
)

type LocalJudger struct {
	Dir      string
	Jdi      judger.JudgerInterface
	JudgeDir string
	RunDir   string
	InDir    string
	OutDir   string
	SrcPath  string
	SpjPath  string
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
	this.InDir = path.Join(this.RunDir, "in")
	this.OutDir = path.Join(this.RunDir, "out")

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
	var src_file string
	switch this.Jdi.GetLanguage().GetLang() {
	case "c++":
		src_file = "src." + sfx
	case "java":
		src_file = "Main.java"
	}
	src_path := path.Join(this.JudgeDir, src_file)
	if err := base.WriteFile(src_path, []byte(this.Jdi.GetCode())); err != nil {
		return err
	}
	this.SrcPath = src_path

	// TODO:spj code
	/*
		if this.Jdi.IsSpj() {
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

func NewCoreMode(lj *LocalJudger) *core.Mode {
	mode := &core.Mode{
		InDir:       lj.InDir,
		OutDir:      lj.OutDir,
		SrcPath:     lj.SrcPath,
		SpjPath:     lj.SpjPath,
		ResPath:     path.Join(lj.RunDir, RESULTFILE),
		RunDir:      lj.RunDir,
		TimeLimit:   lj.Jdi.GetTimeLimit(),
		MemoryLimit: lj.Jdi.GetMemoryLimit(),
	}

	// TODO:set mode
	if lj.Jdi.IsSpj() {
		mode.SetSPJ()
	}

	return mode
}

func EntryPoint(jdi judger.JudgerInterface) {
	log.Println("received run_id=", jdi.GetRunId())
	lj := NewLocalJudger(JUDGEROOT, jdi)
	if err := lj.InitDir(); err != nil {
		log.Fatal(err)
	}
	if err := lj.PrepareData(); err != nil {
		log.Fatal(err)
	}

	mode := NewCoreMode(lj)
	core := core.NewCore(COREPATH)
	core.SetMode(mode)
	if err := core.Run(); err != nil {
		// TODO: set SE
		log.Fatal(err)
	}
	subs, err := core.GetSubmissionStatus()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(subs)
	// TODO: set result

	if err := jdi.UpdateStatus(subs); err != nil {
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
