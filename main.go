package main

import (
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
			EnvVar: "PLUGIN_URL",
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
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	logrus.WithFields(logrus.Fields{
		"Revision": revision,
	}).Info("Drone clair Plugin Version")

	plugin := Plugin{
		Url:       c.String("url"),
		Username:  c.String("username"),
		Password:  c.String("password"),
		ScanImage: c.String("scan_image"),
	}

	return plugin.Exec()
}
