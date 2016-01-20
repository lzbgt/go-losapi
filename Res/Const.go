// Const
package main

import "lzbgt/go-losapi/pub"

var (
	errCmd  = &pub.Err{1, "[%s] is not a Res command."}
	errUrl  = &pub.Err{2, "Url [%s] is unavailable"}
	errName = &pub.Err{3, "Parameter name is error.Reason :[%s]"}
	errVer  = &pub.Err{4, "Parameter ver[%d]<= lastver[%d]"}
)

const (
	ManifestFile = "Manifest.json"
	FieldApplist = "applist"
)
