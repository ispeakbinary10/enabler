package preflight

import (
	"github.com/keitaroinc/enabler/pkg/cmd"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func NewCommand(log *logrus.Logger, streams cmd.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "preflight",
		Short: "Checks if required dependencies are installed in the system",
		Long: `A longer description that spans multiple lines and likely contains examples
				and usage of using your command. For example:

				Cobra is a CLI library for Go that empowers applications.
				This application is a tool to generate the needed files
				to quickly create a Cobra application.`,

		Run: func(cmd *cobra.Command, args []string) {
			// check if java is present on the system
			command := exec.Command("java", "--version")
			cmdOut, err := command.Output()
			if err != nil {
				// java is not present in the system
				log.Fatal("java is not present on the machine, terminating...")
				os.Exit(126)
			}
			version := strings.Split(string(cmdOut), " ")
			if strings.HasPrefix(version[1], "11") {
				log.Info("java jdk 11 ✓")
			} else {
				log.Fatal("Java JDK 11 needed, please change the version of java on your machine.")
			}

			// figure out if there's a better way to check for required dependencies???

			// check if docker is present on the system
			command = exec.Command("docker", "version", "-f", "{{.Server.Version}}")
			cmdOut, err = command.Output()
			if err != nil {
				// docker is not present in the system
				log.Fatal("docker is not present on the machine, terminating...")
			} else {
				log.Infof("docker %s ✓", strings.TrimSpace(string(cmdOut)))
			}
			// check if helm is present on the system
			command = exec.Command("helm", "version", "--short")
			cmdOut, err = command.Output()
			if err != nil {
				// helm is not present in the system
				log.Fatal("helm is not present on the machine, terminating...")
			}
			if strings.HasPrefix(strings.TrimSpace(string(cmdOut)), "v3") {
				log.Infof("helm %s ✓", strings.TrimSpace(string(cmdOut)))
			} else {
				log.Fatal("helm 3 needed, please install it on the machine.")
			}
			// check if kind is present on the system
			command = exec.Command("kind", "version")
			cmdOut, err = command.Output()
			if err != nil {
				// kind is not present in the system
				log.Fatal("kind is not present on the machine, terminating...")
			} else {
				log.Infof("%s ✓", strings.TrimSpace(string(cmdOut)))
			}
			// check if skaffold is present on the system
			command = exec.Command("skaffold", "version")
			cmdOut, err = command.Output()
			if err != nil {
				// skaffold is not present in the system
				log.Fatal("skaffold is not present on the machine, terminating...")
			} else {
				log.Infof("skaffold %s ✓", strings.TrimSpace(string(cmdOut)))
			}
			// check if kubectl is present on the system
			command = exec.Command("kubectl", "version", "--client=true", "--short=true")
			cmdOut, err = command.Output()
			if err != nil {
				// kubectl is not present in the system
				log.Fatal("kubectl is not present on the machine, terminating...")
			} else {
				log.Infof("kubectl %s ✓", strings.TrimSpace(strings.ToLower(string(cmdOut))))
			}
			// check if istioctl is present on the system
			command = exec.Command("istioctl", "version", "-s", "--remote=false")
			cmdOut, err = command.Output()
			if err != nil {
				// istio is not present in the system
				log.Fatal("istioctl is not present on the machine, terminating...")
			}
			version = strings.Split(strings.TrimSpace(string(cmdOut)), ".")
			minorVer, err := strconv.Atoi(version[1])
			if err != nil {
				log.Fatal("unable to parse istio version, terminating...")
			}
			if minorVer >= 5 {
				log.Infof("istio %s ✓", strings.TrimSpace(string(cmdOut)))
			} else {
				log.Fatal("istio 1.5 or greater needed, please update the version of istio on your machine.")
			}
		},
	}
	return cmd
}
