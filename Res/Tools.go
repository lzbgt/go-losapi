// Tools
package main

import "github.com/Leither/hprose-go/hprose"

func getClient(url string) {
	client := hprose.NewClient(url)
	var stub interface{}
	client.UseService(stub)
}
