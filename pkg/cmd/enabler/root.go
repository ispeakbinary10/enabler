package enabler

import (
	"github.com/keitaroinc/enabler/pkg/cmd"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/kind"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/platform"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/preflight"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/setup"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/apps"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cfgFile, kubeCtx string

func NewCommand(logger *logrus.Logger, streams cmd.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Args: cobra.NoArgs,
		Use:   "enabler",
		Short: "Enabler CLI for ease of setup of microservice based apps",
		Long: ``,
	}
	cmd.SetOut(streams.Out)
	cmd.SetErr(streams.ErrOut)

	// set flags
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/enabler/enabler.yaml)")
	cmd.PersistentFlags().StringVarP(&kubeCtx, "kube-context", "", "keitaro", "The kubernetes context to use (defaults to \"keitaro\")")

	// register commands
	cmd.AddCommand(preflight.NewCommand(logger, streams))
	cmd.AddCommand(setup.NewCommand(logger, streams))
	cmd.AddCommand(apps.NewCommand(logger, streams))
	cmd.AddCommand(kind.NewCommand(logger, streams))
	cmd.AddCommand(platform.NewCommand(logger, streams))

	return cmd
}

