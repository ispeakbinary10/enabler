package kind

import "github.com/spf13/cobra"

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a kind cluster",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}