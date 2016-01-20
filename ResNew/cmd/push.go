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
	"io/ioutil"
	"lzbgt/go-losapi/pub"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "push local resource to LOS node",
	Long:  `push local resource file to LOS node`,
	Example: `
  1. lib/jquery.js
    push jquery lib/jquery.js
  2. the whole LOS app described in Manifest.json
    push
  3. main.html and leitherapi.js
    push LeitherApi leitherapi.min.js
    push AppTemplate main.min.bootstrap.html
`,
	Run: Push,
}

var pushTmpl string = `Usage:{{if .Runnable}}
  {{.UseLine}}{{if .HasFlags}} [ResName PathToFileName] [flags]{{end}}{{end}}{{if .HasSubCommands}}
  {{ .CommandPath}} [command]{{end}}{{if gt .Aliases 0}}

Aliases:
  {{.NameAndAliases}}
{{end}}{{if .HasExample}}

Examples:
{{ .Example }}{{end}}{{ if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimRightSpace}}{{end}}{{ if .HasInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimRightSpace}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsHelpCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasSubCommands }}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

func init() {
	RootCmd.AddCommand(pushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//pushCmd.Flags().StringVarP(&ver, "version", "v", "0", "version number")
	pushCmd.SetUsageTemplate(pushTmpl)

}

func PushRes(stub *pub.WebApiStub, BID, name, filename string, ver int) (ID string, err error) {
	log.V("  push", name, filename)
	if name == "" {
		return "", errName.Format("is null")
	}
	var oldver int
	var oldId string
	oldver, err = pub.O2Int(stub.Hget(BID, name, pub.FD_LastVer))
	if err != nil {
		oldver = 0
	}
	log.V("last version:", oldver)
	if ver > 0 && ver <= oldver {
		return "", errVer.Format(ver, oldver)
	}
	strVer := fmt.Sprint(oldver)
	if oldId, err = pub.O2Str(stub.Hget(BID, name, strVer)); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		oldId = ""
	}
	log.V("last id:", oldId)
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
	log.V("new id:", ID)

	if ID == oldId {
		log.V("resource not changed")
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
func Push(cmd *cobra.Command, args []string) {
	if verbose {
		log.Verbose(verbose)
	}

	//fmt.Println("verbose: ", verbose)

	var info pub.LoginInfo
	stub, err := pub.GetStubAndLoginFromCfg(cfg, &info)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to login", err.Error())
		return
	}

	if len(args) == 0 {
		if err := PushManifest(stub, ManifestFile); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	} else if len(args) >= 2 {
		var id string
		var v int
		if len(args) == 3 {
			v, err = strconv.Atoi(args[2])
			if err != nil {
				if args[2] == "last" {
					v = -1
				} else {
					fmt.Fprintln(os.Stderr, "invalid version number:", args[2])
					cmd.Usage()
					return
				}
			}
		}

		id, err = PushRes(stub, info.ID, args[0], args[1], v)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			return
		}
		log.Info("id:", id, "\npushed ok")
	}
}
