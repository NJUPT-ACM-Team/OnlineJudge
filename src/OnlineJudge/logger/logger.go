package logger

import (
	"github.com/op/go-logging"

	"os"
)

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{module} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

var log = logging.MustGetLogger("crawler")

func init() {
	// logging.SetFormatter(format)
	logging.SetLevel(logging.INFO, "crawler")

	// For demo purposes, create two backend for os.Stderr.
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	// Only errors and more severe messages should be sent to backend1
	backend1Leveled := logging.AddModuleLevel(backend1Formatter)
	backend1Leveled.SetLevel(logging.DEBUG, "")
	backend2Leveled := logging.AddModuleLevel(backend2Formatter)
	backend2Leveled.SetLevel(logging.ERROR, "")

	// Set the backends to be used.
	logging.SetBackend(backend1Leveled, backend2Leveled)
}

func GetLogger(name string) *logging.Logger {
	return logging.MustGetLogger(name)
}
