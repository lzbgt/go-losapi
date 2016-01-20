// Manifest
package main

import (
	"encoding/json"
	"fmt"
	"lzbgt/go-losapi/pub"
	"os"
	"path"
)

type ResParam struct {
	pub.LoginInfo
}

type Manifest struct {
	AppName string
	ResFile []string //
	ResName []string
	Res     ResParam
	ids     []string
	stub    *pub.WebApiStub
	nm      NodeManifest
}

type NodeManifest struct {
	AppName string //或许没必要，先留着
	LastVer string `json:"last"`
	Release string `json:"release"`
	ResFile map[string][]string
}

func (m *Manifest) Init(filename string) error {
	fmt.Println("init manifest", filename)
	//读取配置文件
	if err := pub.ReadCfg(filename, m); err != nil {
		return err
	}
	LenPath := len(m.ResFile)
	if len(m.ResName) < LenPath {
		ay := make([]string, 0, LenPath)
		m.ResName = append(ay, m.ResName...)
	}
	fmt.Println("init manifest", m.ResName)
	for i, v := range m.ResFile {
		if m.ResName[i] == "" {
			m.ResName[i] = path.Base(v)
		}
	}
	fmt.Println(m.ResName)
	fmt.Printf("cfg=%v\r\n", &cfg)
	//首先要进入Manifest的路径
	fmt.Printf("path=%s\r\n", path.Dir(filename))
	os.Chdir(path.Dir(filename))
	return nil
}

func (m *Manifest) Push(stub *pub.WebApiStub, strPath string) (err error) {
	//联网
	if stub == nil {
		stub, err = pub.GetStub(m.Res.URL)
		if err != nil {
			return err
		}
		reply, err := stub.Login(m.Res.ID, m.Res.PPT)
		if err != nil {
			return err
		}

		stub.Sid = reply.Sid
	}

	m.stub = stub
	fmt.Println(stub.Sid)
	//依次发布资源,并取得资源id list
	m.ids = make([]string, len(m.ResFile), len(m.ResFile))
	for i, v := range m.ResFile {
		//发布资源：参数：name,资源，返回一个id
		name := m.ResName[i]
		if v == "" {
			m.ids[i] = ""
			continue
		}
		fmt.Println(v)
		if m.ids[i], err = PushRes(stub, m.Res.ID, name, strPath+v, 0); err != nil {
			return err
		}
	}
	return nil
}

func (nm *NodeManifest) init(stub *pub.WebApiStub, ID, appName string) error {
	o, err := stub.Hget(ID, FieldApplist, appName)
	//	fmt.Println(o, err)
	if err != nil {
		return err
	}
	fmt.Println("NodeManifest init")
	str, _ := o.(string)
	buf := []byte(str)
	return json.Unmarshal(buf, nm)
}

//func (nm *NodeManifest) Show()
//func (m *Manifest) readNodeManifest() error {
//	return m.nm.init(m.stub, m.Res.ID, m.AppName)
//	if m.stub == nil {
//		return errors.New("stub is nil")
//	}
//	o, err := m.stub.Hget(m.Res.ID, FieldApplist, m.AppName)
//	fmt.Println(o, err)
//	if err != nil {
//		return err
//	}
//	fmt.Println("readNodeManifest tbd", o)
//	str, _ := o.(string)
//	buf := []byte(str)
//	return json.Unmarshal(buf, &m.nm)
//}

func (m *Manifest) processNodeManifest() {
	if m.nm.ResFile == nil {
		m.nm.ResFile = make(map[string][]string, 10)
	}
	m.nm.AppName = m.AppName
	//lastver 增加一
	var lastVer int
	fmt.Sscanf(m.nm.LastVer, "%d", &lastVer)
	lastVer++
	fmt.Println(lastVer, m.nm.LastVer)

	m.nm.LastVer = fmt.Sprintf("%d", lastVer)
	//增加一个版本的应用
	m.nm.ResFile[m.nm.LastVer] = m.ids
	fmt.Println("processNodeManifest\r\n", &m.nm)
}

func (m *Manifest) generateJs() []byte {
	return nil
}

func (nm *NodeManifest) uploadNodeMainifest(stub *pub.WebApiStub, ID string) error {
	//生成json格式
	buf, err := json.Marshal(nm)
	if err != nil {
		return err
	}
	_, err = stub.Hset(ID, FieldApplist, nm.AppName, string(buf))
	if err != nil {
		return err
	}
	fmt.Println("uploadManifest ok")
	return nil
}

func PushManifest(stub *pub.WebApiStub, filename string) error {
	fmt.Println("PushManifest")
	//读取配置
	var cfg Manifest
	//从配置文件中初始化变量
	err := cfg.Init(filename)
	if err != nil {
		return err
	}
	//把配置中所描述的资源发布，id都放在成员ids中
	if err = cfg.Push(stub, path.Dir(filename)+"/"); err != nil {
		return err
	}

	//读出线上的节点信息，放到成员nm中
	//cfg.readNodeManifest()
	cfg.nm.init(cfg.stub, cfg.Res.ID, cfg.AppName)
	//fmt.Println(cfg.nm)
	//把配置中的资源放到一个新的版本中
	cfg.processNodeManifest()
	//把新的配置信息上传到线上
	cfg.nm.uploadNodeMainifest(cfg.stub, cfg.Res.ID)

	return nil
}

func ReleaseManifest(stub *pub.WebApiStub, id, appName string, ver int) error {
	fmt.Println("ReleaseManifest")
	if ver < -2 {
		return fmt.Errorf("ver[%d] < 0 ", ver)
	}
	var nm NodeManifest
	err := nm.init(stub, id, appName)
	if err != nil {
		return err
	}
	var strVer string
	switch ver {
	case -1:
		strVer = nm.Release
	case -2:
		strVer = nm.LastVer
	default:
		strVer = fmt.Sprint(ver)
	}
	_, ok := nm.ResFile[strVer]
	if !ok {
		return fmt.Errorf("ver[%d] is not exist")
	}
	nm.Release = strVer
	return nm.uploadNodeMainifest(stub, id)
}

// ver = -1 release
// ver = -2 last
func clearManifest(stub *pub.WebApiStub, id, appName string, ver int) (err error) {
	fmt.Println("clearManifest")
	if ver < -2 {
		return fmt.Errorf("ver[%d] < 0 ", ver)
	}
	var nm NodeManifest
	err = nm.init(stub, id, appName)
	if err != nil {
		return err
	}
	//getVer := func(strType string, def int) (ver int) {
	//	if _, err = fmt.Sscan(nm.Release, &ver); err != nil {
	//		return def
	//	}
	//	return
	//}
	//var strVer string
	var lVer, rVer int
	fmt.Sscan(nm.Release, &rVer)
	fmt.Sscan(nm.LastVer, &lVer)

	switch ver {
	case -1:
		ver = rVer
	case -2:
		ver = lVer
	default:
		//strVer = fmt.Sprint(ver)
	}

	fmt.Println("rver=", rVer, "lver=", lVer, "ver=", ver)
	for k, _ := range nm.ResFile {
		fmt.Println("k=", k)
		var v int
		if _, err = fmt.Sscan(k, &v); err != nil {
			return err
		}
		fmt.Println("v=", v)
		if v >= ver || v == lVer || v == rVer {
			continue
		}
		fmt.Println("clear ver=", k)
		delete(nm.ResFile, k)
	}
	return nm.uploadNodeMainifest(stub, id)
}

func showManifest(stub *pub.WebApiStub, id, appName string) (err error) {
	var nm NodeManifest
	err = nm.init(stub, id, appName)
	if err != nil {
		return err
	}
	fmt.Println(nm)
	return nil
}
