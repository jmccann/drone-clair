package main

import (
	"os"
	"os/exec"

	"github.com/Sirupsen/logrus"
)

type (
	Plugin struct {
		Url       string
		Username  string
		Password  string
		ScanImage string
	}
)

func (p Plugin) Exec() error {
	os.Setenv("CLAIR_ADDR", p.Url)
	_, exist := os.LookupEnv("DOCKER_USERNAME")
	if ! exist {
		os.Setenv("DOCKER_USERNAME", p.Username)
	}
	_, exist = os.LookupEnv("DOCKER_PASSWORD")
	if ! exist {
		os.Setenv("DOCKER_PASSWORD", p.Password)
	}

	command := exec.Command(
		"klar",
		p.ScanImage,
	)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Failed to execute a command")
	}

	return nil
}
