package main

import (
	"flag"
	"github.com/TraiOi/checkip"
	"github.com/TraiOi/banip"
)

func main() {
	check_ip := flag.Bool("checkip", false, "Check IP address.")
	ban_ip := flag.Bool("banip", false, "Ban IP address.")
	flag.Parse()

	if *check_ip {
		traioi := new(checkip.IPInfo)
		traioi.GetIP(flag.Args()[0])
		traioi.CheckIP()
	}

	if *ban_ip {
		traioi := new(banip.IPInfo)
		traioi.GetIP(flag.Args()[0])
		traioi.BanIPtoSwitch()
	}
}
