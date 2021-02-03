package create

import (
	"github.com/keitaroinc/enabler/pkg/cmd"
	"github.com/keitaroinc/enabler/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

var config string

func NewCommand(log *logrus.Logger, streams cmd.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a kind cluster",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if config != "" {
				// check if the kind cluster already exists
				kubeContext := cmd.Flag("kube-context").Value
				err := util.GetKind(kubeContext.String(), log)
				if err != nil {
					log.Infof("Kind cluster %s doesn't exist, continue with creation...", kubeContext)
				} else {
					log.Infof("Kind cluster %s already exists, skipping creation...", kubeContext)
					os.Exit(0)
				}
				// check if cluster config file exist
				// TODO: validate yaml
				configPath := config
				if !filepath.IsAbs(config) {
					configPath, _ = filepath.Abs(config)
				}
				if _, err := os.Stat(configPath); os.IsNotExist(err) {
					log.Fatalf("Config file (%s) doesn't exist, aborting cluster creation...", configPath)
				}
				// create the kind cluster
				command := exec.Command("kind", "create", "cluster",
					"--name", kubeContext.String(),
					"--config", configPath,
				)
				_, err = command.Output()
				if err != nil {
					// unable to create the cluster
					log.Errorf("Unable to create the cluster: %s", kubeContext)
					if err, ok := err.(*exec.ExitError); ok {
						os.Exit(err.ExitCode())
					}
				}
				log.Infof("Kind cluster %s created.", kubeContext)
			} else {
				cmd.Help()
			}
		},
	}
	cmd.Flags().StringVarP(&config, "config-file", "", "kind-cluster.yaml", "cluster config path")
	return cmd
}
