package switch_api

import (
	"bytes"
	"encoding/json"
	"github.com/TraiOi/util"
	"github.com/TraiOi/dbs"
)

func (this *SW_INFO) UnbanIP(o_data dbs.ACL) {
	url  := this.API.URL + "/acls/"
	data := util.FormatData(o_data)
	this.del(url, data)
}

func (this *SW_INFO) del(url string, data *bytes.Buffer) {
	var respLog APIResponse
	respBody := util.Curl("DELETE", url, data, this.Token)

	err := json.Unmarshal([]byte(respBody), &respLog)
	if err != nil {
		util.ErrorLogger.Print(err)
	} else {
		util.InfoLogger.Print(respLog.Message)
	}
}
