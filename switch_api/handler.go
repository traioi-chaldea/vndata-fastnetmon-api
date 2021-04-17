package switch_api

import (
	"github.com/TraiOi/dbs"
)

type SW_INFO struct {
	Name  string
	ACL   string
	ROUTE string
	API   dbs.API_INFO
	Token string
}

type APIResponse struct {
	Data    string `json:"data"`
	Message string `json:"message"`
}

func (this *SW_INFO) GetSWInfo() {
	sws := dbs.GetSWInfo().([]map[string]interface{})
	sw  := sws[0]
	this.Name = sw["name"].(string)
	this.ACL = sw["acl"].(string)
	this.ROUTE = sw["route"].(string)
	this.getToken()
}

func (this *SW_INFO) GetAPI() {
	api_info := dbs.GetAPIInfo().(map[string]interface{})
	this.API.Username = api_info["user"].(string)
	this.API.Password = api_info["pass"].(string)
	this.API.URL      = api_info["url"].(string)
}
