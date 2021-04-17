package checkip

import (
	"fmt"
	"strings"
	"github.com/TraiOi/dbs"
)

type IPInfo struct {
	IP	 string `json:"ip"`
	Prefix   string `json:"prefix"`
	Netmask  string `json:"netmask"`
	Gateway  string `json:"gateway"`
	VLAN	 string `json:"vlan"`
	ACL      dbs.ACL
	ROUTE    dbs.ACL
	Route	 string `json:"route"`
	IsBanned bool
}

func (this *IPInfo) GetIP(ip string) {
	this.IP = ip
	this.getInfo()
}

func (this *IPInfo) getPrefix(range_ip string) {
	this.Prefix = strings.Split(range_ip, "/")[1]
}

func (this *IPInfo) getNetmask() {
	switch this.Prefix {
	case "16": this.Netmask = "255.255.0.0"
	case "17": this.Netmask = "255.255.128.0"
	case "18": this.Netmask = "255.255.192.0"
	case "19": this.Netmask = "255.255.224.0"
	case "20": this.Netmask = "255.255.240.0"
	case "21": this.Netmask = "255.255.248.0"
	case "22": this.Netmask = "255.255.252.0"
	case "23": this.Netmask = "255.255.254.0"
	case "24": this.Netmask = "255.255.255.0"
	case "25": this.Netmask = "255.255.255.128"
	case "26": this.Netmask = "255.255.255.192"
	case "27": this.Netmask = "255.255.255.224"
	case "28": this.Netmask = "255.255.255.240"
	case "29": this.Netmask = "255.255.255.248"
	case "30": this.Netmask = "255.255.255.252"
	}
}

func (this *IPInfo) GetACL(seq int) dbs.ACL {
	this.ACL.SEQ = seq
	this.ACL.ACLAction = "permit"
	this.ACL.Protocol = "ip"
	this.ACL.SourceIP = dbs.FormatIP(this.IP)
	this.ACL.SourcePort = "any"
	this.ACL.DestIP = "any"
	this.ACL.DestPort = "any"
	return this.ACL
}

func (this *IPInfo) GetRoute(seq int) dbs.ACL {
	this.ROUTE.SEQ = seq
	this.ROUTE.ACLAction = "permit"
	this.ROUTE.Protocol = "ip"
	this.ROUTE.SourceIP = "any"
	this.ROUTE.SourcePort = "any"
	this.ROUTE.DestIP = dbs.FormatIP(this.IP)
	this.ROUTE.DestPort = "any"
	return this.ROUTE
}

func (this *IPInfo) getInfo() {
	ip_info := dbs.FindIPInfo(this.IP)
	this.VLAN = ip_info.VLAN
	this.ACL.ACLName = ip_info.ACL
	this.ROUTE.ACLName = "ACL_RS_IDS"
	this.Route = ip_info.Route
	this.getPrefix(ip_info.RangeIP)
	this.getNetmask()
	this.Gateway = ip_info.Gateway
	this.IsBanned = ip_info.IsRouteBanned && ip_info.IsACLBanned
}

func (this *IPInfo) CheckIP() {
	fmt.Println("IP Information:")
	fmt.Println("- IP:", this.IP)
	fmt.Println("- Prefix:", this.Prefix)
	fmt.Println("- Netmask:", this.Netmask)
	fmt.Println("- Gateway:", this.Gateway)
	fmt.Println("- VLAN:", this.VLAN)
	fmt.Println("- Route:", this.Route)
}
