package fnm

func (this *IPInfo) ShowACL() {
	this.SW_INFO.GetACL(this.ACL.Hostname, this.ACL.ACLName)
}
