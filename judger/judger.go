package judger

import (
	"OnlineJudge/mq"
)

var jmq *mq.MQ

func Init() error {
	mq.Init()
	jmq := mq.New()
	if err := jmq.Connect(); err != nil {
		return err
	}
	return nil
}

func RunVJ(fn func([]byte)) error {
	return Run("v", fn)
}

func RunLJ(fn func([]byte)) error {
	return Run("l", fn)
}

func Run(oj string, fn func([]byte)) error {
	if err := Init(); err != nil {
		return err
	}
	switch oj {
	case "l":
		if err := jmq.DeclareLJ(); err != nil {
			return err
		}
		if err := jmq.LJReceiver(fn); err != nil {
			return err
		}
	case "v":
		if err := jmq.DeclareVJ(); err != nil {
			return err
		}
		if err := jmq.VJReceiver(fn); err != nil {
			return err
		}
	}
	forever := make(chan bool)
	<-forever
	return nil
}
