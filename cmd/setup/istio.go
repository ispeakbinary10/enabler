package setup

import (
	"fmt"
	"github.com/keitaroinc/enabler/cmd/colors"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var istioCmd = &cobra.Command{
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
			fmt.Println(string(colors.RED), "Istio pre-check failed...aborting install.")
			os.Exit(126)
		}
		fmt.Println("Installing istio, please wait...")

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
			fmt.Println(string(colors.RED), "Istio installation failed.")
			os.Exit(126)
		} else {
			fmt.Println("Istio installed.")
		}
	},
}
