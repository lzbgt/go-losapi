// Login
package pub

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Leither/hprose-go/hprose"
)

type LoginInfo struct {
	ID  string
	PPT string
	URL string
}

func GetStubAndLoginFromCfg(file string, info *LoginInfo) (stub *WebApiStub, err error) {
	if info == nil {
		info = &LoginInfo{}
	}
	if err = ReadCfg(file, &info); err != nil {
		return nil, err
	}

	if stub, err = GetStub(info.URL); err != nil {
		return nil, err
	}
	reply, err := stub.Login(info.ID, info.PPT)
	if err != nil {
		return nil, err
	}
	stub.Sid = reply.Sid
	return
}

func GetStub(url string) (*WebApiStub, error) {
	var stub WebApiStub
	if len(url) == 0 {
		return nil, errors.New("url length is zero")
	}
	t := strings.Replace(url, "http", "ws", 1) //现在支持
	if t == url {
		fmt.Printf("url[%s]的格式可能不合法\r\n", url)
	}
	if t[len(t)-1] != '/' {
		t = t + "/"
	}
	url = t + "ws/" //待优化

	//fmt.Printf(url)
	client := hprose.NewClient(url)
	client.UseService(&stub)
	return &stub, nil
}
