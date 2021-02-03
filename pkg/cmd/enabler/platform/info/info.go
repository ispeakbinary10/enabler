package info

import (
	"fmt"
	"github.com/keitaroinc/enabler/pkg/cmd"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os/exec"
)

func NewCommand(log *logrus.Logger, streams cmd.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Get info on platform and platform components",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// get platform info
			kubeContext := cmd.Flag("kube-context").Value
			command := exec.Command("kubectl",
				"--context", fmt.Sprintf("kind-%s", kubeContext.String()),
				"-n", "istio-system",
				"get", "service", "istio-ingressgateway",
				"-o", "jsonpath={.status.loadBalancer.ingress[0].ip}")
			cmdOut, err := command.Output()
			if err != nil {
				log.Fatal("Unable to get platform info, exiting...")
			}
			log.Infof("Platform can be accessed through the URL: (http://%s)", string(cmdOut))
		},
	}
	return cmd
}