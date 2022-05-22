package tanda

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupLoggerFormat() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC1123Z}
	output.FormatLevel = func(i interface{}) string {
		logLevel := strings.ToUpper(fmt.Sprintf("%-6s", i))

		if i == "error" {
			return BRed + logLevel + Color_Off
		} else if i == "fatal" {
			return BRed + logLevel + Color_Off
		} else if i == "debug" {
			return BYellow + logLevel + Color_Off
		}
		return BCyan + logLevel + Color_Off
	}
	output.FormatTimestamp = func(i interface{}) string {
		// To remove timestamp from log
		// return ""
		return fmt.Sprintf("[%v]", zerolog.TimestampFunc().Format("15:04:05"))
	}

	log.Logger = log.Output(output)
}

func LogSoftwareDetail() {
	log.Info().Msg(fmt.Sprintf("tanda version %v", BuildVersion))
}

func LogDebugArguments(args []string) {
	log.Debug().Int("Arguments length", len(args)).Msg("")
	for i := 0; i < len(args); i++ {
		key := fmt.Sprintf("Arg%v", i)
		log.Debug().Str(key, args[i]).Msg("")
	}
}
