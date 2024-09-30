package main

import (
	"github.com/lxfontes/unprovider/test-component/gen/lxfontes/unprovider-example/caller"
	"github.com/lxfontes/unprovider/test-component/gen/lxfontes/unprovider/runner"
)

//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world client --out gen ./wit

func init() {
	caller.Exports.Call = callHandler
}

func callHandler() string {
	res := runner.Call(`{"please": "indent", "this": "json", "forme": true}`)
	if res.IsErr() {
		return "Error: " + *res.Err()
	}

	return *res.OK()
}

func main() {}
