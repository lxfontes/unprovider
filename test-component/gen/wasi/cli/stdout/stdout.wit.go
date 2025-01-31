// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package stdout represents the imported interface "wasi:cli/stdout@0.2.0".
package stdout

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	"github.com/lxfontes/unprovider/test-component/gen/wasi/io/streams"
)

// GetStdout represents the imported function "get-stdout".
//
//	get-stdout: func() -> output-stream
//
//go:nosplit
func GetStdout() (result streams.OutputStream) {
	result0 := wasmimport_GetStdout()
	result = cm.Reinterpret[streams.OutputStream]((uint32)(result0))
	return
}

//go:wasmimport wasi:cli/stdout@0.2.0 get-stdout
//go:noescape
func wasmimport_GetStdout() (result0 uint32)
