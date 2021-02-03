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
package platform

import (
	"github.com/keitaroinc/enabler/pkg/cmd"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/platform/info"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/platform/keys"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/platform/pinit"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/platform/release"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/platform/version"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewCommand(log *logrus.Logger, streams cmd.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "platform",
		Short: "Platform commands",
		Long: `This command lets you manage kind clusters.

The name of the cluster is taken from the global flag --kube-context
which defaults to "keitaro"`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(info.NewCommand(log, streams))
	cmd.AddCommand(keys.NewCommand(log, streams))
	cmd.AddCommand(pinit.NewCommand(log, streams))
	cmd.AddCommand(release.NewCommand(log, streams))
	cmd.AddCommand(version.NewCommand(log, streams))
	return cmd
}
