package fnm

import (
	"fmt"
	"github.com/TraiOi/dbs"
)

func (this *IPInfo) UnbanIPfromDatabase() {
	if dbs.UnbanIP(this.ACL, this.ROUTE) {
		fmt.Println("Unbanned IP '" + this.IP + "' success!")
	} else {
		fmt.Println("ERROR: Failed to unban IP '" + this.IP + "'!")
	}
}

func (this *IPInfo) UnbanIPfromSwitch() {
	this.SW_INFO.UnbanIP(this.ACL)
	this.SW_INFO.UnbanIP(this.ROUTE)
}

func (this *IPInfo) Unban() {
	if !this.IsBanned {
		fmt.Println("ERROR: IP '" +  this.IP + "' is not existed!")
	} else {
		this.UnbanIPfromDatabase()
		this.UnbanIPfromSwitch()
	}
}
