package common

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func NewLogger(handlerName string) zerolog.Logger {
	return zerolog.New(zerolog.ConsoleWriter{
		Out: os.Stdout,
		FormatTimestamp: func(i interface{}) string {
			return time.Now().Format("2006/01/02 15:04:05") // Format YYYY/MM/DD HH:MM:SS
		},
		FormatLevel: func(i interface{}) string {
			if i == nil {
				return ""
			}
			return i.(string)
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("[%s] %s", handlerName, i.(string))
		},
		NoColor: false,
	}).With().Logger()
}
