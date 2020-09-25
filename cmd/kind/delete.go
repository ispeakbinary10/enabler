package kind

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete kind cluster",
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
		// delete the cluster
		command := exec.Command("kind", "delete", "cluster", "--name", kubeContext.String())
		_, err = command.Output()
		if err != nil {
			// unable to delete the cluster
			fmt.Println(fmt.Sprintf("Unable to delete kind cluster: %s", kubeContext))
			if err, ok := err.(*exec.ExitError); ok {
				os.Exit(err.ExitCode())
			}
		}
		fmt.Println(fmt.Sprintf("Kind cluster %s deleted.", kubeContext))
	},
}
