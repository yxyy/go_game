package pk5001

import (
	"errors"
	"lhc.go.game.sdk/model"
)

type XingwanLogin struct {
	LoginToken string `json:"loginToken" form:"loginToken"`
	model.CommonLoginParam
}

func NewXingwanLogin() *XingwanLogin {
	var login XingwanLogin
	login.Package = 5001
	return &login
}

func (this *XingwanLogin) Vaild() error {
	if this.Package <= 0  {
		return errors.New("包号错误")
	}
	if this.Channel == ""  {
		return errors.New("渠道错误")
	}
	if this.LoginToken == ""  {
		return errors.New("loginToken错误")
	}

	//urls := viper.GetString("pk5001.urls")
	//if urls == "" {
	//	return errors.New("验证url有误")
	//}
	//param := url.Values{}
	//param.Add("userId",this.UserId)
	//param.Add("loginToken",this.LoginToken)
	//param.Add("Token","我是大傻逼!")
	//resp, err := http.Get(urls + "?" + param.Encode())
	//if err!=nil {
	//	return err
	//}
	//defer resp.Body.Close()
	//bytes, err := ioutil.ReadAll(resp.Body)
	//if err !=nil {
	//	return err
	//}
	//if string(bytes) != "ok" {
	//	return errors.New(string(bytes))
	//}

	return nil
}

//server 提升 sync.Pool
func (this *XingwanLogin) GetServer() (interface{} ,error) {
	server := model.NewServer()
	server.Param = &this.CommonLoginParam
	server.Param.Account = server.Param.Channel + "_" + server.Param.UserId
	return  server.DefulatGetServerList()
}
