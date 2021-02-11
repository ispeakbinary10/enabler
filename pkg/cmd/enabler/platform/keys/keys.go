package keys

import (
	"github.com/keitaroinc/enabler/pkg/cmd"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewCommand(log *logrus.Logger, streams cmd.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys",
		Short: "Generate encryption keys used by the application services",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			log.Fatal("Not implemented!")
		},
	}
	return cmd
}