package fnm

func (this *IPInfo) ShowROUTE() {
	this.SW_INFO.GetACL(this.ROUTE.Hostname, this.ROUTE.ACLName)
}
