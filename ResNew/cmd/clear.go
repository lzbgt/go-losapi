// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"lzbgt/go-losapi/pub"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var clsRes bool

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear <AppName | -r ResName> <Number | last> ",
	Short: "reduce the num of versions on LOS node",
	Long:  `reduce the num of versions on LOS node.`,
	Example: `
  1. clear app Demo versions 0 to 10
    clear Demo 10
  2. clear resource jquery version 0 to 10
    clear -r jquery 10`,
	Run: Clear,
}

func init() {
	RootCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	clearCmd.Flags().BoolVarP(&clsRes, "resource", "r", false, "set to clear a single file. otherwise clear whole app")
}

func Clear(cmd *cobra.Command, args []string) {
	log.Verbose(verbose)

	if len(args) != 2 {
		//fmt.Fprintln(os.Stderr, "invalid num of aruguments")
		cmd.Usage()
		return
	}

	var info pub.LoginInfo
	stub, err := pub.GetStubAndLoginFromCfg(cfg, &info)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to login", err.Error())
		return
	}

	v, err := strconv.Atoi(args[1])
	if err != nil {
		if args[1] == "last" {
			v = -1
		} else if args[1] == "release" {
			v = -2
		} else {
			fmt.Fprintln(os.Stderr, "invalid verison number", err.Error())
			return
		}
	}

	if v != -1 && v != -2 && (v-4) < 0 {
		log.Info("no need to clear")
		return
	}
	if clsRes {
		clearRes(stub, info.ID, args[0], v)
	} else {
		clearManifest(stub, info.ID, args[0], v)
	}
}

func clearRes(stub *pub.WebApiStub, ID, name string, ver int) (err error) {
	var c int64
	getVer := func(strType string, def int) (ver int) {
		ver, err = pub.O2Int(stub.Hget(ID, name, strType))
		if err != nil {
			ver = def
		}
		return
	}
	if ver == -1 {
		ver = getVer(pub.FD_Release, -1)
		if ver == -1 {
			log.Info("failed to get release version for res", name)
			return nil
		}
	} else if ver == -2 {
		if ver == -1 {
			log.Info("failed to get last version for res", name)
			return nil
		}
	}

	if (ver - 4) < 0 {
		log.Info("no need to clear")
		return
	}

	rVer := getVer(pub.FD_Release, -1)
	lVer := getVer(pub.FD_LastVer, -1)

	log.Info("clear resouce ", name)
	log.V("range: version 0 ~", ver-4, "except v", rVer, "(release) and", lVer, "(last)")

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

		// for safety, preserve at least 4 versions
		if v >= (ver-4) || v == lVer || v == rVer {
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
