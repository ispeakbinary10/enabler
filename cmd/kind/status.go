package kind

import "github.com/spf13/cobra"

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the status of the kind cluster",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
