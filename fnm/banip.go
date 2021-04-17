package fnm

import (
	"fmt"
	"github.com/TraiOi/dbs"
)

func (this *IPInfo) BanIPtoDatabase() {
	if dbs.BanIP(this.ACL, this.ROUTE) {
		fmt.Println("Banned IP '" + this.IP + "' success!")
	} else {
		fmt.Println("ERROR: Failed to ban IP '" + this.IP + "'!")
	}
}

func (this *IPInfo) BanIPtoSwitch() {
	this.SW_INFO.BanIP(this.ACL)
	this.SW_INFO.BanIP(this.ROUTE)
}

func (this *IPInfo) Ban() {
	if this.IsBanned {
		fmt.Println("ERROR: IP '" +  this.IP + "' is existed!")
	} else {
		this.BanIPtoDatabase()
		this.BanIPtoSwitch()
	}
}

