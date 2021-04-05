package dbs

type TELE_INFO struct {
	Name    string `bson:"name,omitempty"`
	TOKEN   string `bson:"token,omitempty"`
	CHANNEL string `bson:"channel,omitempty"`
}

type API_INFO struct {
	URL      string `bson:"url,omitempty"`
	Username string `bson:"user,omitempty"`
	Password string `bson:"pass,omitempty"`
}

type SW_INFO struct {
	ID      int    `bson:"_id,omitempty" json:"_id"`
	Name    string `bson:"name,omitempty" json:"name"`
	ACL     string `bson:"acl,omitempty" json:"acl"`
	Route   string `bson:"route,omitempty" json:"route"`
}

type VLAN_INFO struct {
	ID      int    `bson:"_id",omitempty" json:"_id"`
	Name    string `bson:"name,omitempty" json:"name"`
	ACL     string `bson:"acl,omitempty" json:"acl"`
	RangeIP string `bson:"range_ip,omitempty" json:"range_ip"`
	Route   string `bson:"route,omitempty" json:"route"`
}

type ROUTE_INFO struct {
	ID      int    `bson:"_id,omitempty" json:"_id"`
	VRF     string `bson:"vrf,omitempty" json:"vrf"`
	NextHop string `bson:"next_hop,omitempty" json:"next_hop"`
	AD      string `bson:"ad,omitempty" json:"ad"`
	Route   string `bson:"route,omitempty" json:"route"`
}

type ROUTE struct {
	VRF      string `bson:"vrf,omitempty" json:"vrf"`
	Hostname string `bson:"hostname,omitempty" json:"hostname"`
	Network  string `bson:"network,omitempty" json:"network"`
	NextHop  string `bson:"next_hop,omitempty" json:"next_hop"`
	AD       string `bson:"ad,omitempty" json:"ad"`
}

type ACL struct {
	Hostname   string `bson:"hostname,omitempty" json:"hostname"`
	ACLName    string `bson:"acl_name,omitempty" json:"acl_name"`
	SEQ        int    `bson:"seq,omitempty" json:"seq"`
	ACLAction  string `bson:"acl_action,omitempty" json:"acl_action"`
	Protocol   string `bson:"protocol,omitempty" json:"protocol"`
	SourceIP   string `bson:"source_ip,omitempty" json:"source_ip"`
	SourcePort string `bson:"source_port,omitempty" json:"source_port"`
	DestIP     string `bson:"dest_ip,omitempty" json:"dest_ip"`
	DestPort   string `bson:"dest_port,omitempty" json:"dest_port"`
}

type ALL_INFO struct {
  SW_INFO    SW_INFO `json:"switch"`
  VLAN_INFO  VLAN_INFO `json:"vlan"`
  ROUTE_INFO ROUTE_INFO `json:"route"`
}

