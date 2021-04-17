package switch_api

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/TraiOi/util"
)

type SW_LOGIN_RESP struct {
	TOKEN  string `json:"access_token"`
	TYPE   string `json:"token_type"`
	STATUS bool `json:"status"`
}

func (this *SW_INFO) getToken() {
	var respData SW_LOGIN_RESP

	this.GetAPI()
	url  := this.API.URL + "/login/account"

	reqData := fmt.Sprintf("username=%s&password=%s",this.API.Username, this.API.Password)
	body := strings.NewReader(reqData)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		util.ErrorLogger.Print(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		util.ErrorLogger.Print(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		util.ErrorLogger.Print(err)
	}

	err = json.Unmarshal([]byte(respBody), &respData)
	if err != nil {
		util.ErrorLogger.Print(err)
	} else {
		util.InfoLogger.Print("Getting API TOKEN ...")
	}

	this.Token = fmt.Sprintf("Bearer %s", respData.TOKEN)
}
