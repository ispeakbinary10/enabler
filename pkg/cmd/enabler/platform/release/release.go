package release

import (
	"github.com/keitaroinc/enabler/pkg/cmd"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewCommand(log *logrus.Logger, streams cmd.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "release",
		Short: "Show releasees",
		Long:  `Release platform by tagging platform repo and
    tagging all individual components (git submodules)
    using their respective SHA that the submodules point at`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Fatal("Not implemented!")
		},
	}
	return cmd
}