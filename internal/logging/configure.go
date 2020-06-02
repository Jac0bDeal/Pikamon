package logging

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Configure configures the logging for use with the bot
func Configure(level string) error {
	logLevel, err := log.ParseLevel(level)
	if err != nil {
		return err
	}
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
	log.SetLevel(logLevel)

	return nil
}
