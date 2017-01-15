package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var revision string // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "clair plugin"
	app.Usage = "clair plugin"
	app.Action = run
	app.Version = revision
	app.Flags = []cli.Flag{

		//
		// plugin args
		//

		cli.StringFlag{
			Name: "url",
			Usage: "clair server URL",
			EnvVar: "CLAIR_URL,PLUGIN_URL",
		},
		cli.StringFlag{
			Name: "username",
			Usage: "docker username",
			EnvVar: "DOCKER_USERNAME,PLUGIN_USERNAME",
		},
		cli.StringFlag{
			Name: "password",
			Usage: "docker password",
			EnvVar: "DOCKER_PASSWORD,PLUGIN_PASSWORD",
		},
		cli.StringFlag{
			Name: "scan_image",
			Usage: "docker image to scan with clair",
			EnvVar: "PLUGIN_SCAN_IMAGE",
		},
		cli.StringFlag{
			Name:   "ca_cert",
			Usage:  "ca cert to trust",
			EnvVar: "CLAIR_CA_CERT,PLUGIN_CA_CERT",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	logrus.WithFields(logrus.Fields{
		"Revision": revision,
	}).Info("Drone Clair Plugin Version")

	// Validate required values provided
	if c.String("url") == "" {
		return fmt.Errorf("Please provide the URL to your Clair server.")
	}
	if c.String("username") == "" {
		return fmt.Errorf("Please provide the username to your Docker Registry.")
	}
	if c.String("password") == "" {
		return fmt.Errorf("Please provide the password to your Docker Registry.")
	}
	if c.String("scan_image") == "" {
		return fmt.Errorf("Please provide the image you want to scan.")
	}

	plugin := Plugin{
		Url:       c.String("url"),
		Username:  c.String("username"),
		Password:  c.String("password"),
		ScanImage: c.String("scan_image"),
		CaCert:    c.String("ca_cert"),
	}

	return plugin.Exec()
}
