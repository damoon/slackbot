package main

import (
	"fmt"
	"log"
	"os"

	bot "slackbot/pkg"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {
	err := run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	err := loadEnv()
	if err != nil {
		return fmt.Errorf("load config: %v", err)
	}

	err = app().Run(args)
	if err != nil {
		return fmt.Errorf("run worklog application: %v", err)
	}

	return nil
}

func loadEnv() error {
	_, err := os.Stat(".env")
	if os.IsNotExist(err) {
		return nil
	}

	err = godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("load .env file: %v", err)
	}

	return nil
}

func app() *cli.App {
	return &cli.App{
		Name:  "slackbot",
		Usage: "a bot to interact via slack",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "token",
				Usage:    "token to access slack",
				EnvVars:  []string{"SLACKBOT_TOKEN"},
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			reponderes := []bot.Reponder{
				&bot.GNULinux{},
				&bot.Smile{},
				&bot.Dadjoke{},
				&bot.CurrentTime{},
			}

			bot.NewBot(c.String("token"), reponderes).Run()

			return nil
		},
	}
}
