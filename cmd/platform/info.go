package platform

import "github.com/spf13/cobra"

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info on platform and platform components",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}