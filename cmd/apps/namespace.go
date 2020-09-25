package apps

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var nsName string

var namespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "Create namespace",
	Long:  `Create a namespace with auto-injection`,
	Run: func(cmd *cobra.Command, args []string) {
		kubeContext := cmd.Flag("kube-context").Value
		if nsName != "" {
			// check if the namespace exists
			command := exec.Command("kubectl", "get", "ns", nsName,
				"--context", fmt.Sprintf("kind-%s", kubeContext),
			)
			_, err := command.Output()
			if err != nil {
				// the namespace doesn't exist, try to create it
				command = exec.Command("kubectl", "create", "ns", nsName,
					"--context", fmt.Sprintf("kind-%s", kubeContext),
				)
				_, err := command.Output()
				if err != nil {
					// unable to create the namespace, exit with original exit code
					fmt.Println(fmt.Sprintf("Unable to create namespace: %s", nsName))
					if err, ok := err.(*exec.ExitError); ok {
						os.Exit(err.ExitCode())
					}
				}
				fmt.Println(fmt.Sprintf("Created a namespace for: %s", nsName))
				// label the created namespace
				command = exec.Command("kubectl", "label", "namespace", nsName,
					"istio-injection=enabled",
					"--context", fmt.Sprintf("kind-%s", kubeContext),
				)
				_, err = command.Output()
				if err != nil {
					// unable to label the namespace, exit with original exit code
					fmt.Println(fmt.Sprintf("Unable to label the namespace: %s", nsName))
					if err, ok := err.(*exec.ExitError); ok {
						os.Exit(err.ExitCode())
					}
				}
				fmt.Println(fmt.Sprintf("Labeled %s namespace for istio injection.", nsName))
			} else {
				fmt.Println(fmt.Sprintf("Skipped creation of %s since it already exists.", nsName))
			}
		} else {
			cmd.Help()
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// namesoaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	namespaceCmd.Flags().StringVarP(&nsName, "name", "n", "", "The name of the k8s namespace")
}