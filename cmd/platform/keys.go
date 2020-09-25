package platform

import "github.com/spf13/cobra"

var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Generate encryption keys used by the application services",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
