package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate returns the command line application ready to execute
func Generate() *cli.App {
	app := cli.NewApp()

	app.Name = "Command Line Application"
	app.Usage = "Searches for IP's and Server Names on the Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "pinterest.com",
		},
	}

	app.Commands = []cli.Command{ //this is a struct
		{
			Name:   "ip",
			Usage:  "Searches the IPs of internet addresses",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servers",
			Usage:  "Searches for the name of internet servers",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
