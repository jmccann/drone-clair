package main

import (
	"io/ioutil"
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
		CaCert    string
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

	var commands []*exec.Cmd

	if p.CaCert != "" {
		commands = append(commands, installCaCert(p.CaCert))
	}

	commands = append(commands, scanImage(p.ScanImage))

	for _, command := range commands {
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Fatal("Failed to execute a command")
		}
	}

	return nil
}

func installCaCert(cacert string) *exec.Cmd {
	ioutil.WriteFile("/usr/local/share/ca-certificates/ca_cert.crt", []byte(cacert), 0644)
	return exec.Command(
		"update-ca-certificates",
	)
}

func scanImage(image string) *exec.Cmd {
	return exec.Command(
		"klar",
		image,
	)
}
