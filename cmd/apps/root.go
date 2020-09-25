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
package apps

import (
	"github.com/spf13/cobra"
)

// kindCmd represents the kind command
var MainCmd = &cobra.Command{
	Use:   "apps",
	Short: "Application commands",
	Long: `Application specific commands such as creation of kubernetes
    objects such as namespaces, configmaps etc. 

The name of the cluster is taken from the global flag --kube-context
which defaults to "keitaro"`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// MainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// MainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// initialize and register sub-commands e.g.:
	// MainCmd.AddCommand(sub)
	MainCmd.AddCommand(namespaceCmd)
}
