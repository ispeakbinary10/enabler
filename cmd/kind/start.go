package kind

import (
	"context"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/keitaroinc/enabler/cmd/util"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"time"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start kind cluster",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log := util.NewLogger("INFO", nil)
		// Kind creates containers with a label io.x-k8s.kind.cluster
		// Kind naming is clustername-control-plane and clustername-worker{x}
		// The idea is to find the containers and stop them
		kubeContext := cmd.Flag("kube-context").Value
		err := getKind(kubeContext.String())
		if err != nil {
			log.Errorf("Kind cluster %s doesn't exist, terminating.", kubeContext)
			if err, ok := err.(*exec.ExitError); ok {
				os.Exit(err.ExitCode())
			}
		} else {
			// init docker client
			cli, err := client.NewEnvClient()
			if err != nil {
				panic(err)
			}
			// create filter and get relevant containers
			f := filters.NewArgs()
			f.Add("name", fmt.Sprintf("%s-control-plane", kubeContext))
			f.Add("name", fmt.Sprintf("%s-worker", kubeContext))

			containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{Filters: f, All: true})
			if err != nil {
				// network not found
				panic(err)
			}
			// start stopped containers
			containerStartOpts := types.ContainerStartOptions{
				CheckpointID:  "",
				CheckpointDir: "",
			}
			for _, container := range containers {
				if container.State != "running" {
					err := cli.ContainerStart(context.Background(), container.ID, containerStartOpts)
					if err != nil {
						log.Fatalf("Unable to start container: %s", container.Names[0])
					} else {
						log.Infof("Container %s started.", container.Names[0])
					}
				}
			}
			// set kube context for started containers
			containers, err = cli.ContainerList(context.Background(), types.ContainerListOptions{Filters: f, All: true})
			for _, container := range containers {
				err := setKubeConfig(container, kubeContext.String())
				if err != nil {
					log.Fatalf("Unable to set kube context (%s) for container (%s)", kubeContext.String(), container.ID)
				}
			}
			// start spinner and wait for the cluster to be ready
			s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
			s.Color("green")
			s.Prefix = "Waiting for the kind cluster to be up..."
			s.Start()
			maxRetries := 30
			for ok := true; ok; ok = maxRetries != 0 {
				maxRetries -= 1
				err := getClusterInfo(kubeContext.String())
				if err != nil && maxRetries == 0 {
					log.Fatalf("Unable to start kind cluster (%s)", kubeContext.String())
				}
				time.Sleep(1 * time.Second)
			}
			s.Stop()
			log.Infof("Kind cluster (%s) was started.", kubeContext.String())
		}
	},
}
