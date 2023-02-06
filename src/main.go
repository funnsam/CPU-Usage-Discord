package main

import (
	"fmt"
	"time"

	"github.com/hugolgst/rich-go/client"
	"github.com/shirou/gopsutil/v3/host"
)

var cpuImages = []string{"cpu_g", "cpu_y", "cpu_r"}

func main() {
	checkErr(client.Login("1071015586034884618"))
	fmt.Printf("\x1b[0;36;1minfo:\x1b[0m Logged into RPC.\n")

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	_bootTime, err := host.BootTime()
	checkErr(err)
	bootTime := time.Unix(int64(_bootTime), 0)

	info, err := host.Info()
	checkErr(err)

	largeText := fmt.Sprintf("I'm on %s btw", info.OS)

	for range ticker.C {
		details, cpuLevel := getStat()
		_ = cpuLevel

		client.SetActivity(client.Activity{
			Details: details,
			Timestamps: &client.Timestamps{
				Start: &bootTime,
			},
			LargeImage: cpuImages[cpuLevel],
			LargeText:  largeText,
			Buttons: []*client.Button{
				{
					Label: "😎 Check out this",
					Url:   "https://github.com/funnsam",
				},
			},
		})
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
