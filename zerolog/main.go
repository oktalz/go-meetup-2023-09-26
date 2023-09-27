package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// START OMIT
type LineInfoHook struct{}

func (h LineInfoHook) Run(e *zerolog.Event, l zerolog.Level, msg string) {
	_, file, line, ok := runtime.Caller(3)
	parts := strings.Split(file, "/")
	file = parts[len(parts)-1]
	if ok {
		e.Str("src", fmt.Sprintf("%s:%d", file, line))
	}
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.DateTime,
	}).Hook(LineInfoHook{})

	log.Info().Str("msg", "Hello world!").Msg("")
}

// END OMIT
