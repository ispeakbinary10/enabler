package platform

import (
	"github.com/keitaroinc/enabler/cmd/util"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"os"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show platform version",
	Long: `Check versions of microservices in git submodules
    You can provide a comma separated list of submodules
    or you can use 'all' for all submodules`,
	Run: func(cmd *cobra.Command, args []string) {
		log := util.NewLogger("INFO", nil)
		// init repo
		r, err := git.PlainOpen(repoPath)
		if err != nil {
			log.Fatal(err)
		}
		// fetch repo HEAD
		ref, err := r.Head()
		if err != nil {
			log.Fatal(err)
		}
		// fetch repo config
		c, _ := r.Config()
		if origin, ok := c.Remotes["origin"]; ok {
			log.Infof("Checking platform version for repository (%s@%s)", origin.URLs[0], ref.Hash().String())
		}

		// fetch worktree
		w, err := r.Worktree()
		if err != nil {
			log.Fatal(err)
		}

		var subModulesList git.Submodules
		if subModules == "all" {
			subModulesList, err = w.Submodules()
			if err != nil {
				log.Error(err)
			}
			for _, sub := range subModulesList {
				r, err := sub.Repository()
				if err != nil {
					continue
				}
				ref, err := r.Head()
				if err != nil {
					continue
				}
				log.Infof("%s@%s", sub.Config().URL, ref.Hash().String())
			}
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	cwd, _ := os.Getwd()
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	versionCmd.Flags().StringVarP(&repoPath, "repo-path", "", cwd, "repository location")
	versionCmd.Flags().StringVarP(&subModules, "submodules", "", "all", "submodules to update")
}
