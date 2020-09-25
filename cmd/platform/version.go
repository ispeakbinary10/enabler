package platform

import "github.com/spf13/cobra"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show platform version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
