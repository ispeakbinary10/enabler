package enabler

import (
	"github.com/keitaroinc/enabler/pkg/cmd"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/preflight"
	"github.com/keitaroinc/enabler/pkg/cmd/enabler/setup"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cfgFile, kubeCtx string

func NewCommand(logger *logrus.Logger, streams cmd.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Args: cobra.NoArgs,
		Use:   "enabler",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
	cmd.SetOut(streams.Out)
	cmd.SetErr(streams.ErrOut)

	// set flags
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/enabler/enabler.yaml)")
	cmd.PersistentFlags().StringVarP(&kubeCtx, "kube-context", "", "keitaro", "The kubernetes context to use (defaults to \"keitaro\")")

	// register commands
	cmd.AddCommand(preflight.NewCommand(logger, streams))
	cmd.AddCommand(setup.NewCommand(logger, streams))

	return cmd
}

