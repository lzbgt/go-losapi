// Tools
package pub

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadCfg(cfgfile string, v interface{}) error {
	rd, err := ioutil.ReadFile(cfgfile)
	if err != nil {
		return err
	}
	if len(rd) == 0 {
		return nil
	}
	rdbuf := bytes.NewBuffer(rd)
	decoder := json.NewDecoder(rdbuf)
	return decoder.Decode(v)
}

func SaveCfg(cfgfile string, v interface{}) error {
	fp, err := os.Create(cfgfile)
	if err != nil {
		return err
	}
	defer fp.Close()
	encoder := json.NewEncoder(fp)
	return encoder.Encode(v)
}

func O2Int(o interface{}, err error) (int, error) {
	if err != nil {
		return 0, err
	}
	switch v := o.(type) {
	case int:
		return v, nil
	case float64:
		return int(v), nil
	case string:
		var i int
		_, err := fmt.Sscan(v, &i)
		return i, err
	default:
		return 0, errValue.Format(v)
	}
}

func O2Str(o interface{}, err error) (string, error) {
	if err != nil {
		return "", err
	}
	switch v := o.(type) {
	case string:
		return v, nil
	default:
		return "", errValue.Format(v)
	}
}

func Obj2Json(o interface{}) ([]byte, error) {
	var ay []byte
	switch v := o.(type) {
	case *[]uint8:
		ay = *v
	case []uint8:
		ay = v
	default:
		return json.Marshal(o)
	}

	b := make([]uint8, 0, len(ay)+1)
	b = append(b, '(')
	b = append(b, ay...)
	return b, nil
}

func Obj2Key(o interface{}) (string, error) {
	if b, err := Obj2Json(o); err != nil {
		return "", err
	} else {
		sha := sha256.Sum256(b)
		s := sha[:]
		key := base64.URLEncoding.EncodeToString(s)
		Len := len(key)
		if Len > 0 && key[Len-1] == '=' {
			key = key[:Len-1]
		}
		return key, nil
	}
}
