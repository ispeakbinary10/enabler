package delete

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
		Use:   "delete",
		Short: "Delete kind cluster",
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
			// delete the cluster
			command := exec.Command("kind", "delete", "cluster", "--name", kubeContext.String())
			_, err = command.Output()
			if err != nil {
				// unable to delete the cluster
				log.Errorf("Unable to delete kind cluster: %s", kubeContext)
				if err, ok := err.(*exec.ExitError); ok {
					os.Exit(err.ExitCode())
				}
			}
			log.Infof("Kind cluster %s deleted.", kubeContext)
		},
	}
	return cmd
}
