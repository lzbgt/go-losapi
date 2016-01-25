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

	"github.com/spf13/cobra"
)

var bShowRes bool

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show <ResName | -r AppName>",
	Short: "show info of app or resource file",
	Long:  `show info of app or resource file on LOS node.`,
	Example: `
  1. examine the Demo app
    show Demo
  2. examine the jquery resouce
    show -r jquery`,
	Run: Show,
}

func init() {
	RootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	showCmd.Flags().BoolVarP(&bShowRes, "resouce", "r", false, "show a resource file")

}

func Show(cmd *cobra.Command, args []string) {
	log.Verbose(verbose)

	if len(args) < 1 {
		//fmt.Fprintln(os.Stderr, "missing arguments")
		cmd.Usage()
		return
	}

	stub, err := pub.GetStubAndLoginFromCfg(cfg, &info)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to Login\n", "error:", err.Error())
		return
	}
	if bShowRes {
		err = showRes(stub, info.ID, args[0])
	} else {
		err = showManifest(stub, info.ID, args[0])
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err.Error())
	}

}

func showRes(stub *pub.WebApiStub, ID, resName string) error {
	var oldver int
	var err error
	oldver, err = pub.O2Int(stub.Hget(ID, resName, pub.FD_LastVer))
	log.Info("ResName:", resName)
	if err != nil {
		return errors.New("failed to get last version of " + resName + "\nerror: " + err.Error())
	}
	log.Info("Last Version:", oldver)
	oldver, err = pub.O2Int(stub.Hget(ID, resName, pub.FD_Release))
	if err != nil {
		return errors.New("failed to get release version of " + resName + "\nerror: " + err.Error())
	}

	log.Info("Release Version:", oldver)

	//	if log.verbose {
	//		keys, err := stub.Hkeys(ID, resName)
	//		if err != nil {
	//			return errors.New("failed to get info of ", resName, "\nerror: ", err.Error())
	//		}

	//		for _, key := range keys {
	//			fmt.Println(key)
	//		}
	//	}

	return nil
}
