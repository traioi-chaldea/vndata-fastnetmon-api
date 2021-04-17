package switch_api

import (
	"fmt"
	"encoding/json"
	"github.com/TraiOi/util"
)

func (this *SW_INFO) GetACL(hostname string, acl_name string) {
	data     := fmt.Sprintf("hostname=%s&acl_name=%s", hostname,
							   acl_name)
	url_acl  := this.API.URL + "/acls"
	url      := fmt.Sprintf("%s/?%s", url_acl, data)
	this.get(url)
}

func (this *SW_INFO) get(url string) {
	var respLog APIResponse
	respBody := util.Curl("GET", url, nil, this.Token)

	err := json.Unmarshal([]byte(respBody),  &respLog)
	if err != nil {
		util.ErrorLogger.Print(err)
	} else {
		fmt.Println(respLog.Data)
	}
}
