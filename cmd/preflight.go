/*
Copyright © 2020 Keitaro Inc dev@keitaro.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/keitaroinc/enabler/cmd/colors"
)

// TODO: Specify versions and dependencies via config file?

// preflightCmd represents the preflight command
var preflightCmd = &cobra.Command{
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
			fmt.Println(string(colors.RED), "java is not present on the machine, terminating...")
			os.Exit(126)
		}
		version := strings.Split(string(cmdOut), " ")
		if strings.HasPrefix(version[1], "11") {
			fmt.Println(string(colors.WHITE), "java jdk 11", string(colors.GREEN), "✓")
		} else {
			fmt.Println(string(colors.RED), "Java JDK 11 needed, please change the version of java on your machine.")
			os.Exit(1)
		}

		// figure out if there's a better way to check for required dependencies???

		// check if docker is present on the system
		command = exec.Command("docker", "version", "-f", "{{.Server.Version}}")
		cmdOut, err = command.Output()
		if err != nil {
			// java is not present in the system
			fmt.Println(string(colors.RED), "docker is not present on the machine, terminating...")
			os.Exit(126)
		} else {
			fmt.Println(string(colors.WHITE), "docker "+strings.TrimSpace(string(cmdOut)), string(colors.GREEN), "✓")
		}
		// check if helm is present on the system
		command = exec.Command("helm", "version", "--short")
		cmdOut, err = command.Output()
		if err != nil {
			// java is not present in the system
			fmt.Println(string(colors.RED), "helm is not present on the machine, terminating...")
			os.Exit(126)
		}
		if strings.HasPrefix(strings.TrimSpace(string(cmdOut)), "v3") {
			fmt.Println(string(colors.WHITE), "helm "+strings.TrimSpace(string(cmdOut)), string(colors.GREEN), "✓")
		} else {
			fmt.Println(string(colors.RED), "helm 3 needed, please install it on the machine.")
			os.Exit(1)
		}
		// check if kind is present on the system
		command = exec.Command("kind", "version")
		cmdOut, err = command.Output()
		if err != nil {
			// kind is not present in the system
			fmt.Println(string(colors.RED), "kind is not present on the machine, terminating...")
			os.Exit(126)
		} else {
			fmt.Println(string(colors.WHITE), strings.TrimSpace(string(cmdOut)), string(colors.GREEN), "✓")
		}
		// check if skaffold is present on the system
		command = exec.Command("skaffold", "version")
		cmdOut, err = command.Output()
		if err != nil {
			// kind is not present in the system
			fmt.Println(string(colors.RED), "skaffold is not present on the machine, terminating...")
			os.Exit(126)
		} else {
			fmt.Println(string(colors.WHITE), "skaffold "+strings.TrimSpace(string(cmdOut)), string(colors.GREEN), "✓")
		}
		// check if kubectl is present on the system
		command = exec.Command("kubectl", "version", "--client=true", "--short=true")
		cmdOut, err = command.Output()
		if err != nil {
			// kind is not present in the system
			fmt.Println(string(colors.RED), "kubectl is not present on the machine, terminating...")
			os.Exit(126)
		} else {
			fmt.Println(string(colors.WHITE), "kubectl "+strings.TrimSpace(strings.ToLower(string(cmdOut))), string(colors.GREEN), "✓")
		}
		// check if istioctl is present on the system
		command = exec.Command("istioctl", "version", "-s", "--remote=false")
		cmdOut, err = command.Output()
		if err != nil {
			// kind is not present in the system
			fmt.Println(string(colors.RED), "istioctl is not present on the machine, terminating...")
			os.Exit(126)
		}
		version = strings.Split(strings.TrimSpace(string(cmdOut)), ".")
		minorVer, err := strconv.Atoi(version[1])
		if err != nil {
			fmt.Println(string(colors.RED), "unable to parse istio version, terminating...")
			os.Exit(126)
		}
		if minorVer >= 5 {
			fmt.Println(string(colors.WHITE), "istio "+strings.TrimSpace(string(cmdOut)), string(colors.GREEN), "✓")
		} else {
			fmt.Println(string(colors.RED), "istio 1.5 or greater needed, please update the version of istio on your machine.")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(preflightCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// preflightCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// preflightCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
