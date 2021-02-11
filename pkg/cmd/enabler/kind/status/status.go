package status

import (
	"github.com/keitaroinc/enabler/pkg/cmd"
	"github.com/keitaroinc/enabler/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func NewCommand(log *logrus.Logger, streams cmd.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Check the status of the kind cluster",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// check if the kind cluster exists
			kubeContext := cmd.Flag("kube-context").Value
			err := util.GetKind(kubeContext.String(), log)
			if err != nil {
				log.Errorf("Kind cluster %s doesn't exist, terminating.", kubeContext)
				if err, ok := err.(*exec.ExitError); ok {
					os.Exit(err.ExitCode())
				}
			}
			// check the status of the cluster
			err = util.GetClusterInfo(kubeContext.String())
			if err != nil {
				log.Errorf("Kind cluster %s not running, please start the cluster.", kubeContext)
				if err, ok := err.(*exec.ExitError); ok {
					os.Exit(err.ExitCode())
				} else {
					os.Exit(2)
				}
			}
			log.Infof("Kind cluster %s is running.", kubeContext)
		},
	}
	return cmd
}
