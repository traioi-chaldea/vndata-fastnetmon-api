package banip

import (
	"fmt"
	"github.com/TraiOi/util"
	"github.com/TraiOi/checkip"
	"github.com/TraiOi/dbs"
)

type IPInfo struct {
	IP    string `json:"ip"`
	ACL   dbs.ACL
	ROUTE dbs.ACL
}

func (this *IPInfo) GetIP(ip string) {
	this.IP = ip
	this.getIPInfo()
}

func (this *IPInfo) checkIP() bool {
	return dbs.CheckIP(this.IP)
}

func (this *IPInfo) getIPInfo() {
	ip := new(checkip.IPInfo)
	ip.GetIP(this.IP)
	this.ACL = ip.GetACL()
	// Route info
	this.ROUTE = ip.GetRoute()
}

func (this *IPInfo) BanIPtoSwitch() {
	switchs_info := dbs.GetSWInfo().([]map[string]interface{})
	switch_info := switchs_info[0]
	this.ACL.Hostname = switch_info["acl"].(string)
	this.ROUTE.Hostname = switch_info["route"].(string)

	fmt.Println(util.FormatData(this.ACL))
	fmt.Println(util.FormatData(this.ROUTE))
}
