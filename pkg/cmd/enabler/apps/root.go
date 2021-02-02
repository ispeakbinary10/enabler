package apps

import (
	"github.com/keitaroinc/enabler/pkg/cmd"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/apps/namespace"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewCommand(log *logrus.Logger, streams cmd.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apps",
		Short: "Application commands",
		Long: `Application specific commands such as creation of kubernetes objects such as namespaces, configmaps etc.
The name of the cluster is taken from the global flag --kube-context which defaults to "keitaro"`,

		Run: func(cmd *cobra.Command, args []string) {
			cmd.AddCommand(namespace.NewCommand(log, streams))
			cmd.Help()
		},
	}
	return cmd
}
