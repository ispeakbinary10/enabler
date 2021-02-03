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
package kind

import (
	"github.com/keitaroinc/enabler/pkg/cmd"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/kind/create"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/kind/delete"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/kind/start"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/kind/status"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/kind/stop"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewCommand(log *logrus.Logger, streams cmd.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kind",
		Short: "Manage kind clusters",
		Long: `This command lets you manage kind clusters.

The name of the cluster is taken from the global flag --kube-context
which defaults to "keitaro"`,

		Run: func(cmd *cobra.Command, args []string) {
			cmd.AddCommand(create.NewCommand(log, streams))
			cmd.AddCommand(delete.NewCommand(log, streams))
			cmd.AddCommand(start.NewCommand(log, streams))
			cmd.AddCommand(stop.NewCommand(log, streams))
			cmd.AddCommand(status.NewCommand(log, streams))
			cmd.Help()
		},
	}
	return cmd
}
