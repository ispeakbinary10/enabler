package setup

import (
	"github.com/keitaroinc/enabler/pkg/cmd"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/setup/initcmd"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/setup/istio"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/setup/metallb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewCommand(log *logrus.Logger, streams cmd.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setup",
		Short: "Checks if required dependencies are installed in the system",
		Long: `You can use the setup command to download and configure all infrastructure services such as: kind, kubectl, istioctl, helm and skaffold.`,

		Run: func(cmd *cobra.Command, args []string) {
			cmd.AddCommand(initcmd.NewCommand(log, streams))
			cmd.AddCommand(istio.NewCommand(log, streams))
			cmd.AddCommand(metallb.NewCommand(log, streams))
			cmd.Help()
		},
	}
	return cmd
}
