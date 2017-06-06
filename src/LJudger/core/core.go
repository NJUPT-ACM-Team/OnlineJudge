package core

import (
	"OnlineJudge/irpc"
)

type Result struct {
	Status          string
	StatusCode      string
	TimeUsed        int
	MemoryUsed      int
	CEInfo          int
	TestcasesPassed int
}

const (
	MODE_SPJ  uint16 = 0x0001
	MODE_OI   uint16 = 0x0010
	MODE_ICPC uint16 = 0x0100
)

type Mode struct {
	UseSudo     bool
	RunDir      string
	InDir       string
	OutDir      string
	SrcPath     string
	SpjPath     string
	ResPath     string
	TimeLimit   int
	MemoryLimit int
	JudgingMode uint16 // SPJ: 0x0001, OI: 0x0010, ICPC: 0x0100
}

func (this *Mode) Set(m uint16) {
	this.JudgingMode |= m
}

func (this *Mode) Unset(m uint16) {
	this.JudgingMode = (this.JudgingMode | m) ^ m
}

func (this *Mode) SetSPJ() {
	this.Set(MODE_SPJ)
}

func (this *Mode) UnsetSPJ() {
	this.Unset(MODE_SPJ)
}

func (this *Mode) SetOI() {
	this.Set(MODE_OI)
}

func (this *Mode) UnsetOI() {
	this.Unset(MODE_OI)
}

func (this *Mode) SetICPC() {
	this.Set(MODE_ICPC)
}

func (this *Mode) UnsetICPC() {
	this.Unset(MODE_ICPC)
}

func (this *Mode) IsSPJ() bool {
	return (this.JudgingMode & MODE_SPJ) == MODE_SPJ
}

func (this *Mode) IsOI() bool {
	return (this.JudgingMode & MODE_OI) == MODE_OI
}

func (this *Mode) IsICPC() bool {
	return (this.JudgingMode & MODE_ICPC) == MODE_ICPC
}

// mode: SPJ, OI, ICPC

type Core interface {
	SetMode(*Mode)
	Run() error
	GetSubmissionStatus() (*irpc.SubmissionStatus, error)
}
