package main

import (
	"flag"
	"github.com/TraiOi/checkip"
	"github.com/TraiOi/fnm"
)

func main() {
	check_ip   := flag.Bool("checkip"  , false, "Check IP address.")
	ban_ip     := flag.Bool("banip"    , false, "Ban IP address.")
	unban_ip   := flag.Bool("unbanip"  , false, "Unban IP address.")
	show_acl   := flag.Bool("showacl"  , false, "Show ACLs List.")
	show_route := flag.Bool("showroute", false, "Show ROUTEs List.")
	flag.Parse()

	traioi := new(fnm.IPInfo)
	traioi.GetIP(flag.Args()[0])

	if *check_ip {
		traicam := new(checkip.IPInfo)
		traicam.GetIP(flag.Args()[0])
		traicam.CheckIP()
	}

	if *ban_ip {
		traioi.Ban()
	}

	if *unban_ip {
		traioi.Unban()
	}

	if *show_acl {
		traioi.ShowACL()
	}

	if *show_route {
		traioi.ShowROUTE()
	}
}
