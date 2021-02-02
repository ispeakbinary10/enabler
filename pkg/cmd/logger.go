package cmd

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	// initialize new logger
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{FullTimestamp: true}

	// set level
	lvl, err := logrus.ParseLevel("debug")
	if err != nil {
		log.Error(errors.Wrapf(err, "Unable to parse log level (level=%s), setting default", "de"))
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.Debugf("Setting log level (level=%s)", lvl)
		log.SetLevel(lvl)
	}
	return log
}
