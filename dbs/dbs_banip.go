package dbs

import (
	"github.com/TraiOi/util"
)

func checkIPFromROUTE(ip string) bool {
	check := false
	routes_list := getInfo("ROUTE")
	for _, v := range routes_list {
		if v["network"] == FormatIP(ip) {
			check = true
			util.WarningLogger.Print("IP '", ip, "' is existed from ROUTE '", v["vrf"], "'")
			break
		}
	}
	return check
}

func checkIPFromACL(ip string) bool {
	check := false
	acls_list := getInfo("ACL")
	for _, v := range acls_list {
		if v["source_ip"] == FormatIP(ip) {
			check = true
			util.WarningLogger.Print("IP '", ip, "' is existed from ACL '", v["acl"], "'")
		break
		}
	}
	return check
}

func CheckIP(ip string) bool {
	var check bool
	checkACL   := checkIPFromACL(ip)
	checkROUTE := checkIPFromROUTE(ip)
	if checkACL && checkROUTE {
		check = true
	} else {
		check = false
	}
	return check
}

