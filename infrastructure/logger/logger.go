package logger

import (
	"flag"

	"github.com/rs/zerolog"
)

func Initialize() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()
	// Default level for this example is panic, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.PanicLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
