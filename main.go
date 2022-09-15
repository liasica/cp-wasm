// to build wasm:
// 1. GOOS=js GOARCH=wasm go build -o build/encryption.wasm .
// 2. cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" build/
// 3. import encryption.wasm and wasm_exec.js to html

//go:build js && wasm

package main

import (
    "fmt"
    "github.com/chatpuppy/encryption-wasm/console"
    "github.com/chatpuppy/encryption-wasm/encryption"
    "syscall/js"
)

var done chan struct{}

func init() {
    done = make(chan struct{})
}

func main() {
    console.Info("wasm worked😇🥳🤗")

    // generate keys
    addFunc("ecdhGenerate", encryption.Generate)

    // share ecdh keys
    addFunc("ecdhShare", encryption.Share)

    // encrypt
    addFunc("ecdhEncrypt", encryption.AesEncrypt)

    // decrypt
    addFunc("ecdhDecrypt", encryption.AesDecrypt)

    <-done
}

func addFunc(name string, fn func(this js.Value, args []js.Value) any) {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                console.Error(fmt.Sprintf("Recovered from func: %v", r))
            }
        }()
        js.Global().Set(name, js.FuncOf(fn))
    }()
}
