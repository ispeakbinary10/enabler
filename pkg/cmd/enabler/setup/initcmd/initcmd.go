package initcmd

import (
	"github.com/keitaroinc/enabler/pkg/cmd"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os/exec"
)

func NewCommand(log *logrus.Logger, streams cmd.IOStreams) *cobra.Command {
	var plugin string
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize infrastructure services",
		Long: `You can use the setup command to download and configure all infrastructure services such as: kind, kubectl, istioctl, helm and skaffold.`,

		Run: func(cmd *cobra.Command, args []string) {
			if &plugin != nil {
				command := exec.Command(plugin)
				cmdOut, err := command.Output()
				if err != nil {
					log.Fatal(err)
				} else {
					log.Info(string(cmdOut))
				}
				log.Info("executed sometginf")
			} else {
				defaultCommand(log)
			}
		},
	}
	cmd.Flags().StringVarP(&plugin, "plugin", "", "asdf", "repository location")
	return cmd
}

