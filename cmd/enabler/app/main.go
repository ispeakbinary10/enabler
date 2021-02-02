package app

import (
	"github.com/sirupsen/logrus"
	"os"

	"github.com/keitaroinc/enabler/pkg/cmd"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler"
)

func Main() {
	if err := Run(cmd.NewLogger(), cmd.StandardIOStreams(), os.Args[1:]); err != nil {
		os.Exit(1)
	}
}

// Run invokes the enabler root command, returning the error.

func Run(logger *logrus.Logger, streams cmd.IOStreams, args []string) error {

	// actually run the command
	c := enabler.NewCommand(logger, streams)
	c.SetArgs(args)
	if err := c.Execute(); err != nil {
		logError(logger, err)
		return err
	}
	return nil
}

func logError(logger *logrus.Logger, err error) {

	logger.Errorf("ERROR: %v", err)

}