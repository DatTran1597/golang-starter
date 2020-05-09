package util

import "github.com/urfave/cli"

func StringFlag(env, name, usage, value string) cli.StringFlag {
	return cli.StringFlag{
		EnvVar: env,
		Name:   name,
		Usage:  usage,
		Value:  value,
	}
}
