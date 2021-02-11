package istio

import (
	"fmt"
	"github.com/keitaroinc/enabler/pkg/cmd"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os/exec"
)

func NewCommand(log *logrus.Logger, streams cmd.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "istio",
		Short: "Configure istio",
		Long:  `Verify system installation of istio and check if we are ready to setup istio`,
		Run: func(cmd *cobra.Command, args []string) {
			kubeContext := cmd.Flag("kube-context").Value
			// check if istio is present on the system
			command := exec.Command("istioctl", "verify-install", "--context", fmt.Sprintf("kind-%s", kubeContext))
			_, err := command.Output()
			if err != nil {
				// istio verification failed, exit
				log.Fatal("Istio pre-check failed...aborting install.")
			}
			log.Infof("Installing istio, please wait...")
			// TODO: Add spinner because setup takes a while
			// TODO: configure istio through config?
			command = exec.Command("istioctl", "manifest", "apply", "-y",
				"--set", "profile=default",
				"--set", "addonComponents.grafana.enabled=true",
				"--set", "addonComponents.kiali.enabled=true",
				"--set", "addonComponents.prometheus.enabled=true",
				"--set", "addonComponents.tracing.enabled=true",
				"--set", "values.kiali.dashboard.jaegerURL=http://jaeger-query:16686",
				"--set", "values.kiali.dashboard.grafanaURL=http://grafana:3000",
				"--context", fmt.Sprintf("kind-%s", kubeContext),
			)
			_, err = command.Output()
			if err != nil {
				// istio installation failed, exit
				log.Fatal("Istio installation failed.")
			} else {
				log.Infof("Istio installed.")
			}
		},
	}
	return cmd
}