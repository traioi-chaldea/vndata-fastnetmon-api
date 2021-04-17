package dbs

import (
	"net"
	"os"
	"github.com/TraiOi/util"
	"github.com/yl2chen/cidranger"
	"go.mongodb.org/mongo-driver/bson"
)

type IPInfo struct {
	IP            string `json:"ip"`
	RangeIP       string `json:"range_ip"`
	Gateway       string `json:"gateway"`
	VLAN          string `json:"vlan"`
	ACL	      string `json:"acl"`
	Route         string `json:"route"`
	IsRouteBanned bool   `json:"isroutebanned"`
	IsACLBanned   bool   `json:"isaclbanned"`
}

func FindIPInfo(ip string) IPInfo {
	var check  bool
	var err    error
	var result IPInfo
	vlans := getInfo("VLAN_INFO")

	ranger := cidranger.NewPCTrieRanger()
	for _, v := range vlans {
		range_ip := v["range_ip"].(string)
		_, tmp, _ := net.ParseCIDR(range_ip)
		ranger.Insert(cidranger.NewBasicRangerEntry(*tmp))
		check, err = ranger.Contains(net.ParseIP(ip))
		if err != nil {
			util.ErrorLogger.Print(err)
		}
		if check {
			util.InfoLogger.Print("Found range IP '", range_ip, "'")
			result.VLAN = v["name"].(string)
			result.RangeIP = range_ip
			result.Gateway = v["gateway"].(string)
			result.ACL = v["acl"].(string)
			result.Route = v["route"].(string)
			result.IsRouteBanned = checkIsRouteBanned(ip)
			result.IsACLBanned = checkIsACLBanned(ip)
			break
		}
		ranger.Remove(*tmp)
	}
	if !check {
		util.ErrorLogger.Print("IP '", ip, "' can't be found")
		os.Exit(1)
	}
	return result
}

func checkIsRouteBanned(ip string) bool {
	var err error
	var tmp bson.M

	col  := DB.Collection("ROUTE")
	ip32 := FormatIP(ip)

	if err = col.FindOne(ctx, bson.M{"dest_ip": ip32}).Decode(&tmp); err != nil {
	    util.ErrorLogger.Print(err)
	}

	return !util.IsNil(tmp)
}

func checkIsACLBanned(ip string) bool {
	var err	error
	var tmp bson.M

	col  := DB.Collection("ACL")
	ip32 := FormatIP(ip)

	if err = col.FindOne(ctx, bson.M{"source_ip": ip32}).Decode(&tmp); err != nil {
	    util.ErrorLogger.Print(err)
	}

	return !util.IsNil(tmp)
}
