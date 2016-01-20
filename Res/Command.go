// Command
package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"lzbgt/go-losapi/pub"
	"strings"
)

func PushRes(stub *pub.WebApiStub, BID, name, filename string, ver int) (ID string, err error) {
	fmt.Println("PushRes: ", name, "file:", filename)
	if name == "" {
		return "", errName.Format("is null")
	}
	var oldver int
	var oldId string
	oldver, err = pub.O2Int(stub.Hget(BID, name, pub.FD_LastVer))
	if err != nil {
		oldver = 0
	}
	fmt.Println("old version:", oldver)
	if ver > 0 && ver <= oldver {
		return "", errVer.Format(ver, oldver)
	}
	strVer := fmt.Sprint(oldver)
	if oldId, err = pub.O2Str(stub.Hget(BID, name, strVer)); err != nil {
		fmt.Println(err.Error())
		oldId = ""
	}
	fmt.Println("old id:", oldId)
	//读取文件
	var rd []byte
	rd, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	ID, err = stub.Setdata(BID, rd)
	if err != nil {
		return
	}
	fmt.Println("new id:", ID)

	if ID == oldId {
		fmt.Println("resource not changed")
		return
	}
	ver = oldver + 1
	strver := fmt.Sprintf("%d", oldver+1)

	stub.Hset(BID, name, strver, ID)
	stub.Hset(BID, name, pub.FD_LastVer, ver)
	return
}

//参数：
//-name :资源名
//-file :文件名
//-ver:资源的版本，缺省是自动生成
func Push() error {
	if flag.NArg() < 2 {
		return errors.New("Res 资源名")
	}

	var info pub.LoginInfo
	stub, err := pub.GetStubAndLoginFromCfg(*cfg, &info)
	if err != nil {
		return fmt.Errorf("GetStubAndLoginFromCfg err:%s", err.Error())
	}

	filename := flag.Arg(1)
	//fmt.Printf("Push:filename=%s name=%s, cfg=%s, ver=%d\r\n", filename, *name, *cfg, *ver)
	if strings.HasSuffix(filename, ManifestFile) {
		return PushManifest(stub, filename)
	}

	var id string
	id, err = PushRes(stub, info.ID, *name, filename, *ver)
	fmt.Println("resource current id is ", id)
	return err
}

//参数：
//-name :资源名
//-ver:资源的版本，缺省是自动生成
func Release() error {
	//if *name == "" {
	//	if flag.NArg() < 2 {
	//		return errors.New("Res show name")
	//	}
	//	*name = flag.Arg(1)
	//}

	var info pub.LoginInfo
	stub, err := pub.GetStubAndLoginFromCfg(*cfg, &info)
	if err != nil {
		return err
	}
	if *appName != "" {
		return ReleaseManifest(stub, info.ID, *appName, *ver)
	}
	return ReleaseRes(stub, info.ID, *name, *ver)
}

func ReleaseRes(stub *pub.WebApiStub, id, resName string, ver int) (err error) {
	if ver == 0 {
		ver, err = pub.O2Int(stub.Hget(id, resName, pub.FD_LastVer))
		if err != nil {
			ver = 0
		}
	}
	if _, err = stub.Hset(id, resName, pub.FD_Release, ver); err != nil {
		return err
	}
	fmt.Printf("set res[%s] version to [%d]\r\n", resName, ver)
	return nil
}

//参数：
//-name :资源名
func Show() error {
	fmt.Printf("Push: name=%s\r\n", *name)
	var info pub.LoginInfo
	stub, err := pub.GetStubAndLoginFromCfg(*cfg, &info)
	if err != nil {
		return err
	}
	if *appName != "" {
		showManifest(stub, info.ID, *appName)
	}
	return showRes(stub, info.ID, *name)
}

func showRes(stub *pub.WebApiStub, ID, resName string) error {
	keys, err := stub.Hkeys(ID, resName)
	if err != nil {
		return err
	}
	for _, key := range keys {
		fmt.Println(key)
	}
	var oldver int
	oldver, err = pub.O2Int(stub.Hget(ID, resName, pub.FD_LastVer))
	if err != nil {
		oldver = 0
	}
	fmt.Printf("last version is %d\r\n", oldver)
	oldver, err = pub.O2Int(stub.Hget(ID, resName, pub.FD_Release))
	if err != nil {
		oldver = 0
	}
	fmt.Printf("release version is %d\r\n", oldver)

	return nil
}

//参数：
//-name :资源名
//-ver :版本
//清除给定资源给定版本之前的所有资源
func Clear() error {
	var info pub.LoginInfo
	stub, err := pub.GetStubAndLoginFromCfg(*cfg, &info)
	if err != nil {
		return err
	}
	if *appName != "" {
		clearManifest(stub, info.ID, *appName, *ver)
	}
	return clearRes(stub, info.ID, *name, *ver)
}

// ver = -1 release
// ver = -2 last
func clearRes(stub *pub.WebApiStub, ID, name string, ver int) (err error) {
	//if name == "" {
	//	if flag.NArg() < 2 {
	//		return errors.New("Res clear name ver")
	//	}
	//	name = flag.Arg(1)
	//	if *ver == 0 {
	//		if flag.NArg() < 3 {
	//			return errors.New("Res clear name ver")
	//		}
	//		strVer := flag.Arg(2)
	//		if strVer == pub.FD_Release {
	//			*ver, err = pub.O2Int(stub.Hget(ID, name, pub.FD_Release))
	//			if err != nil {
	//				*ver = 0
	//			}
	//		} else {
	//			if _, err = fmt.Sscan(strVer, ver); err != nil {
	//				return err
	//			}
	//		}
	//	}
	//}
	var c int64
	getVer := func(strType string, def int) (ver int) {
		ver, err = pub.O2Int(stub.Hget(ID, name, strType))
		if err != nil {
			ver = def
		}
		return
	}
	if ver == -1 {
		ver = getVer(pub.FD_Release, 0)
	} else if ver == -2 {
		ver = getVer(pub.FD_LastVer, 0)
	}
	rVer := getVer(pub.FD_Release, -1)
	lVer := getVer(pub.FD_LastVer, -1)
	//rVer, err := pub.O2Int(stub.Hget(ID, name, pub.FD_Release))
	//if err != nil {
	//	//fmt.Println(err.Error())
	//	rVer = -1
	//}
	//lVer, err := pub.O2Int(stub.Hget(ID, name, pub.FD_LastVer))
	//if err != nil {
	//	//fmt.Println(err.Error())
	//	lVer = -1
	//}
	fmt.Printf("clear name=%s ver =%d\r\n", name, ver)
	//versions, err := stub.Hkeys(ID, name)
	//if err != nil {
	//	return err
	//}
	fvs, err := stub.Hgetall(ID, name)
	if err != nil {
		return err
	}

	var mapRef = make(map[string]int, 10)
	for _, fv := range fvs {
		id, _ := fv.Value.(string) //resid
		mapRef[id] = mapRef[id] + 1
	}

	fmt.Printf("last version:%d, release version:%d\r\n", lVer, rVer)
	for _, fv := range fvs {
		v, err := pub.O2Int(fv.Field, nil) //版本v
		if err != nil {
			continue
		}
		id, _ := fv.Value.(string) //resid
		fmt.Println("cur version:", v, " get id :", fv.Value, "refs=", mapRef[id])
		if v >= ver || v == lVer || v == rVer {
			//要确保release和debug能正常使用
			continue
		}
		//根据版本取资源的id
		if mapRef[id] <= 1 {
			fmt.Printf("ref[%d] <= 1, no use, to del res\r\n", mapRef[id])
			if c, err = stub.Del(ID, id); err != nil {
				return err
			}
			fmt.Println("del res ok. c=", c, "id:", fv.Value)
		}

		if c, err = stub.Hdel(ID, name, fv.Field); err != nil {
			return err
		}
		fmt.Println("c=", c, "version:", fv.Field, " deleted")
	}
	return nil
}
