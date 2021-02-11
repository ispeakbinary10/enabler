package util

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

func InSlice(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func NewLogger(level string, formatter logrus.Formatter) *logrus.Logger {
	// initialize new logger
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{FullTimestamp: true}

	if formatter != nil {
		log.Formatter = formatter
	}
	// set level
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		log.Error(errors.Wrapf(err, "Unable to parse log level (level=%s), setting default", level))
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.Debugf("Setting log level (level=%s)", lvl)
		log.SetLevel(lvl)
	}
	return log
}

func GetKind(cluster string, log *logrus.Logger) error {
	command := exec.Command("kind", "get", "clusters")
	cmdOut, err := command.Output()
	if err != nil {
		// unable to get kind clusters
		log.Errorf("Unable to get kind cluster: %s", cluster)
		if err, ok := err.(*exec.ExitError); ok {
			log.Error(err.ExitCode())
			return err
		}
	}
	clusters := strings.Split(string(cmdOut), "\n")
	_, found := InSlice(clusters, cluster)
	if found {
		return nil
	}
	return errors.New(fmt.Sprintf("Cluster %s not found.", cluster))
}

func GetClusterInfo(cluster string) error {
	command := exec.Command("kubectl", "cluster-info", "--context", fmt.Sprintf("kind-%s", cluster))
	_, err := command.Output()
	if err != nil {
		return err
	}
	return nil
}

func SetKubeConfig(container types.Container, cluster string, log *logrus.Logger) error {
	log.Debugf("Setting port: %d on container (%s)", container.Ports[0].PublicPort, container.ID)
	command := exec.Command("kubectl", "config", "set-cluster", fmt.Sprintf("kind-%s", cluster),
		"--server", fmt.Sprintf("https://127.0.0.1:%d", container.Ports[0].PublicPort))
	_, err := command.Output()
	if err != nil {
		// unable to set kube config
		log.Errorf("Unable to set kube config: %s", cluster)
		return err
	}
	return nil
}