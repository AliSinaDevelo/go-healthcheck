package main

import (
	"os"
	"fmt"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App {
		Name: "Healthchecker",
		Usage: "A tiny tool that checks if the website is up and running",
		Flags: []cli.Flag{
			&cli.StringFlag {
				Name: "domain",
				Aliases: []string{"d"},
				Usage: "Domain name to see if running",
				Required: true
			} ,
			&cli.StringFlag {
				Name: "port",
				Aliases: []string{"p"},
				Usage: "Port number",
				Required: false
			} ,
		},
		Action: func(c *cli.Context) error {
			port := c.String("port") 
			if c.String("port") == "" {
				port = "80"
			}
			status := Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
 		},
		
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}