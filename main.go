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
    addFunc(func() {
        js.Global().Set("ecdhGenerate", js.FuncOf(encryption.Generate))
    })

    // share ecdh keys
    addFunc(func() {
        js.Global().Set("ecdhShare", js.FuncOf(encryption.Share))
    })

    <-done
}

func addFunc(cb func()) {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                console.Error(fmt.Sprintf("Recovered from func: %v", r))
            }
        }()

        cb()
    }()
}
