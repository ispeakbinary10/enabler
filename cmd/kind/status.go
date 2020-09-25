package kind

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the status of the kind cluster",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// check if the kind cluster exists
		kubeContext := cmd.Flag("kube-context").Value
		err := getKind(kubeContext.String())
		if err != nil {
			fmt.Println(fmt.Sprintf("Kind cluster %s doesn't exist, terminating.", kubeContext))
			if err, ok := err.(*exec.ExitError); ok {
				os.Exit(err.ExitCode())
			}
		}
		// check the status of the cluster
		err = getClusterInfo(kubeContext.String())
		if err != nil {
			fmt.Println(fmt.Sprintf("Kind cluster %s not running, please start the cluster.", kubeContext))
			if err, ok := err.(*exec.ExitError); ok {
				os.Exit(err.ExitCode())
			} else {
				os.Exit(2)
			}
		}
		fmt.Println(fmt.Sprintf("Kind cluster %s is running.", kubeContext))
	},
}
