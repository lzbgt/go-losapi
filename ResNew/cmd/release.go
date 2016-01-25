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
	"errors"
	"fmt"
	"lzbgt/go-losapi/pub"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

const (
	API_REGAPP = "http://api.leither.cn:8890/v1/res/regapp"
)

var relRes bool

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release <AppName | -r ResName> [Number | last]",
	Short: "set release number of an app or a file",
	Long:  `set release number of an pp or a file.`,
	Example: `
1. the last version of app Demo
release Demo last
2. version 10 of app Demo
release Demo 10
3. the last version of jquery
release -r jquery last`,
	Run: Release,
}

func init() {
	RootCmd.AddCommand(releaseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// releaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// releaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	releaseCmd.Flags().BoolVarP(&relRes, "resource", "r", false, "set to release a single file. otherwise release whole app")
}

func Release(cmd *cobra.Command, args []string) {
	log.Verbose(verbose)

	if len(args) < 1 {
		// fmt.Fprintln(os.Stderr, "missing AppName")
		cmd.Usage()
		return
	}

	res := args[0]
	var v int
	var err error
	if len(args) == 2 {
		v, err = strconv.Atoi(args[1])
	} else {
		v = -1
	}

	if err != nil {
		if args[1] == "last" {
			v = -1
		} else {
			fmt.Fprintln(os.Stderr, "invalid version\n", cmd.Help())
			return
		}
	}

	err = nil
	stub, err := pub.GetStubAndLoginFromCfg(cfg, &info)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	if !relRes {
		err = ReleaseManifest(stub, info.ID, res, v)

	} else {
		err = ReleaseRes(stub, info.ID, res, v)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func ReleaseRes(stub *pub.WebApiStub, id, resName string, ver int) error {
	lastVer, err := pub.O2Int(stub.Hget(id, resName, pub.FD_LastVer))
	if err != nil {
		return errors.New("failed to get version for this resouce \"" + resName + "\"\nerror:" + err.Error())
	}
	if ver == -1 {
		ver = lastVer
	}

	if _, err = stub.Hset(id, resName, pub.FD_Release, ver); err != nil {
		return err
	}

	fmt.Printf("released res \"%s\" version %d \r\n", resName, ver)

	return nil
}
