package main

import (
	"os"
	"path"
	"testing"
)

func Test_Array(t *testing.T) {
	//Path := "..\\html\\a.txt"

	t.Log(path.Base("/a/b"))
	t.Log(path.Dir("/a/b"))
	t.Log(path.IsAbs("/a/b"))
	os.Chdir("..")
}
