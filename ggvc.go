package main

import (
  "fmt"
  "log"
  "os"
  "github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "Gradle git version calculator"
	app.Usage = "Script for generating semantic version based on the version in gradle.properties and git branch"
	app.Author = "Pavel Titenkov" 
	app.Version = "0.0.1"
}

func commands() {
  app.Commands = []cli.Command{
    {
      Name:    "version",
      Aliases: []string{"v", "-v", "--v", "--version"},
      Usage:   "Calculate semantic version",
      Action: func(c *cli.Context) {
        // TODO: implementation should come here..
        fmt.Println("1.0.0")
      },
    },
  }
}

func main() {
  info()
  commands()

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}