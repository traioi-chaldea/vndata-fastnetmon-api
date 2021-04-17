package fnm

import (
	"github.com/TraiOi/checkip"
	"github.com/TraiOi/dbs"
	"github.com/TraiOi/switch_api"
)

type IPInfo struct {
	IP       string `json:"ip"`
	ACL      dbs.ACL
	ROUTE    dbs.ACL
	SW_INFO  *switch_api.SW_INFO
	IsBanned bool
}

func (this *IPInfo) GetIP(ip string) {
	this.IP = ip
	this.getIPInfo()
}

func (this *IPInfo) getIPInfo() {
	ip := new(checkip.IPInfo)
	ip.GetIP(this.IP)
	this.ACL = ip.GetACL(0)
	// Route info
	this.ROUTE = ip.GetRoute(0)
	this.IsBanned = ip.IsBanned
	// Get Switch Hostname
	this.getSWInfo()
}

func (this *IPInfo) getSWInfo() {
	this.SW_INFO = new(switch_api.SW_INFO)
	this.SW_INFO.GetSWInfo()
	this.ACL.Hostname = this.SW_INFO.ACL
	this.ROUTE.Hostname = this.SW_INFO.ROUTE
}
