package platform

import (
	"github.com/keitaroinc/enabler/cmd/util"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"os"
)

var repoPath, subModules string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize platform",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log := util.NewLogger("INFO", nil)

		// init repo
		r, err := git.PlainOpen(repoPath)
		if err != nil {
			log.Fatal(err)
		}
		// init worktree
		w, err := r.Worktree()
		if err != nil {
			log.Fatal(err)
		}
		// update submodules
		var subModulesList git.Submodules
		if subModules == "all" {
			subModulesList, err = w.Submodules()
			if err != nil {
				log.Error(err)
			}
			for _, sub := range subModulesList {
				log.Infof("Fetching latest changes for (%s)", sub.Config().Name)
				r, err := sub.Repository()
				if err != nil {
					continue
				}
				w, err := r.Worktree()
				if err != nil {
					continue
				}
				_ = w.Pull(&git.PullOptions{RemoteName: "origin"})
			}
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// namesoaceCmd.PersistentFlags().String("foo", "", "A help for foo")
	cwd, _ := os.Getwd()
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().StringVarP(&repoPath, "repo-path", "", cwd, "repository location")
	initCmd.Flags().StringVarP(&subModules, "submodules", "", "all", "submodules to update")
}
