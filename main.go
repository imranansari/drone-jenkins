package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli"
)

// Version set at compile-time
var Version string

func main() {
	app := cli.NewApp()
	app.Name = "jenkins plugin"
	app.Usage = "trigger jenkins jobs"
	app.Copyright = "Copyright (c) 2017 Bo-Yi Wu"
	app.Authors = []cli.Author{
		{
			Name:  "Bo-Yi Wu",
			Email: "appleboy.tw@gmail.com",
		},
	}
	app.Action = run
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "host",
			Usage:  "jenkins base url",
			EnvVar: "PLUGIN_URL,JENKINS_URL",
		},
		cli.StringFlag{
			Name:   "user,u",
			Usage:  "jenkins username",
			EnvVar: "PLUGIN_USER,JENKINS_USER",
		},
		cli.StringFlag{
			Name:   "token,t",
			Usage:  "jenkins token",
			EnvVar: "PLUGIN_TOKEN,JENKINS_TOKEN",
		},
		cli.StringSliceFlag{
			Name:   "job,j",
			Usage:  "jenkins job",
			EnvVar: "PLUGIN_JOB,JENKINS_JOB",
		},
		cli.StringFlag{
			Name:   "env-file",
			Usage:  "source env file",
			EnvVar: "ENV_FILE",
			Value:  ".env",
		},
	}

	// Override a template
	cli.AppHelpTemplate = `
________                                            ____.              __   .__
\______ \_______  ____   ____   ____               |    | ____   ____ |  | _|__| ____   ______
 |    |  \_  __ \/  _ \ /    \_/ __ \   ______     |    |/ __ \ /    \|  |/ /  |/    \ /  ___/
 |    |   \  | \(  <_> )   |  \  ___/  /_____/ /\__|    \  ___/|   |  \    <|  |   |  \\___ \
/_______  /__|   \____/|___|  /\___  >         \________|\___  >___|  /__|_ \__|___|  /____  >
        \/                  \/     \/                        \/     \/     \/       \/     \/
                                                                    version: {{.Version}}
NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
   {{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
   {{end}}{{if .Version}}
VERSION:
   {{.Version}}
   {{end}}
REPOSITORY:
    Github: https://github.com/appleboy/drone-line
`

	app.Run(os.Args)
}

func run(c *cli.Context) error {
	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}

	plugin := Plugin{
		BaseURL:  c.String("host"),
		Username: c.String("user"),
		Token:    c.String("token"),
		Job:      c.StringSlice("job"),
	}

	return plugin.Exec()
}
